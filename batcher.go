package kallax

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
)

type batchQueryRunner struct {
	schema        Schema
	cols          []string
	q             Query
	oneToOneRels  []Relationship
	oneToManyRels []Relationship
	db            squirrel.DBProxy
	builder       squirrel.SelectBuilder
	total         int
	eof           bool
	// records is the cache of the records in the last batch.
	records []Record
}

var ErrNoMoreRows = errors.New("kallax: there are no more rows in the result set")

func newBatchQueryRunner(schema Schema, db squirrel.DBProxy, q Query) *batchQueryRunner {
	cols, builder := q.compile()
	var (
		oneToOneRels  []Relationship
		oneToManyRels []Relationship
	)

	for _, rel := range q.getRelationships() {
		switch rel.Type {
		case OneToOne:
			oneToOneRels = append(oneToOneRels, rel)
		case OneToMany:
			oneToManyRels = append(oneToManyRels, rel)
		}
	}

	return &batchQueryRunner{
		schema:        schema,
		cols:          cols,
		q:             q,
		oneToOneRels:  oneToOneRels,
		oneToManyRels: oneToManyRels,
		db:            db,
		builder:       builder,
	}
}

func (r *batchQueryRunner) next() (Record, error) {
	if r.eof {
		return nil, ErrNoMoreRows
	}

	if len(r.records) == 0 {
		records, err := r.loadNextBatch()
		if err != nil {
			return nil, err
		}

		if len(records) == 0 {
			r.eof = true
			return nil, ErrNoMoreRows
		}

		r.total += len(records)
		r.records = records[1:]
		return records[0], nil
	}

	record := r.records[0]
	r.records = r.records[1:]
	return record, nil
}

func (r *batchQueryRunner) loadNextBatch() ([]Record, error) {
	limit := r.q.GetLimit() - uint64(r.total)
	if r.q.GetBatchSize() < limit || limit <= 0 {
		limit = r.q.GetBatchSize()
	}

	rows, err := r.builder.
		Offset(r.q.GetOffset() + uint64(r.total)).
		Limit(limit).
		RunWith(r.db).
		Query()

	if err != nil {
		return nil, err
	}

	return r.processBatch(rows)
}

func (r *batchQueryRunner) processBatch(rows *sql.Rows) ([]Record, error) {
	batchRs := NewResultSet(
		rows,
		r.q.isReadOnly(),
		r.oneToOneRels,
		r.cols...,
	)

	var records []Record
	for batchRs.Next() {
		var rec = r.schema.New()
		if err := batchRs.Scan(rec); err != nil {
			return nil, err
		}
		records = append(records, rec)
	}

	if err := batchRs.Close(); err != nil {
		return nil, err
	}

	var ids = make([]interface{}, len(records))
	for i, r := range records {
		ids[i] = r.GetID()
	}

	for _, rel := range r.oneToManyRels {
		indexedResults, err := r.getRecordRelationships(ids, rel)
		if err != nil {
			return nil, err
		}

		for _, r := range records {
			err := r.SetRelationship(rel.Field, indexedResults[r.GetID()])
			if err != nil {
				return nil, err
			}

			// If the relationship is partial, we can not ensure the results
			// in the field reflect the truth of the database.
			// In this case, the parent is marked as non-writable.
			if rel.Filter != nil {
				r.setWritable(false)
			}
		}
	}

	return records, nil
}

type indexedRecords map[interface{}][]Record

func (r *batchQueryRunner) getRecordRelationships(ids []interface{}, rel Relationship) (indexedRecords, error) {
	fk, ok := r.schema.ForeignKey(rel.Field)
	if !ok {
		return nil, fmt.Errorf("kallax: cannot find foreign key on field %s for table %s", rel.Field, r.schema.Table())
	}

	filter := In(fk, ids...)
	if rel.Filter != nil {
		And(rel.Filter, filter)
	} else {
		rel.Filter = filter
	}

	q := NewBaseQuery(rel.Schema)
	q.Where(rel.Filter)
	cols, builder := q.compile()
	rows, err := builder.RunWith(r.db).Query()
	if err != nil {
		return nil, err
	}

	relRs := NewResultSet(rows, false, nil, cols...)
	var indexedResults = make(indexedRecords)
	for relRs.Next() {
		rec, err := relRs.Get(rel.Schema)
		if err != nil {
			return nil, err
		}

		val, err := rec.Value(fk.String())
		if err != nil {
			return nil, err
		}

		indexedResults[val] = append(indexedResults[val], rec)
	}

	if err := relRs.Close(); err != nil {
		return nil, err
	}

	return indexedResults, nil
}
