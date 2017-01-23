package tests

import (
	"database/sql"
	"fmt"

	"github.com/src-d/go-kallax"
	"github.com/src-d/go-kallax/types"
)

var _ types.SQLType

// NewEventsFixture returns a new instance of EventsFixture.
func NewEventsFixture() (record *EventsFixture) {
	record = newEventsFixture()
	if record != nil {
		record.SetID(kallax.NewID())
	}
	return
}

func (r *EventsFixture) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return &r.Model.ID, nil
	case "checks":
		return types.JSON(&r.Checks), nil
	case "must_fail_before":
		return types.JSON(&r.MustFailBefore), nil
	case "must_fail_after":
		return types.JSON(&r.MustFailAfter), nil

	default:
		return nil, fmt.Errorf("invalid column in EventsFixture: %s", col)
	}
}

func (r *EventsFixture) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.Model.ID, nil
	case "checks":
		return types.JSON(r.Checks), nil
	case "must_fail_before":
		return types.JSON(r.MustFailBefore), nil
	case "must_fail_after":
		return types.JSON(r.MustFailAfter), nil

	default:
		return nil, fmt.Errorf("invalid column in EventsFixture: %s", col)
	}
}

func (r *EventsFixture) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {

	}
	return nil, fmt.Errorf("model EventsFixture has no relationship %s", field)
}

func (r *EventsFixture) SetRelationship(field string, record kallax.Record) error {
	switch field {

	}
	return fmt.Errorf("model EventsFixture has no relationship %s", field)
}

// EventsFixtureStore is the entity to access the records of the type EventsFixture
// in the database.
type EventsFixtureStore struct {
	*kallax.Store
}

// NewEventsFixtureStore creates a new instance of EventsFixtureStore
// using a SQL database.
func NewEventsFixtureStore(db *sql.DB) *EventsFixtureStore {
	return &EventsFixtureStore{kallax.NewStore(db, Schema.EventsFixture.BaseSchema)}
}

// Insert inserts a EventsFixture in the database. A non-persisted object is
// required for this operation.
func (s *EventsFixtureStore) Insert(record *EventsFixture) error {

	return s.Store.Insert(record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *EventsFixtureStore) Update(record *EventsFixture, cols ...kallax.SchemaField) (int64, error) {

	return s.Store.Update(record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *EventsFixtureStore) Save(record *EventsFixture) (updated bool, err error) {

	return s.Store.Save(record)
}

// Delete removes the given record from the database.
func (s *EventsFixtureStore) Delete(record *EventsFixture) error {
	return s.Store.Delete(record)
}

// Find returns the set of results for the given query.
func (s *EventsFixtureStore) Find(q *EventsFixtureQuery) (*EventsFixtureResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewEventsFixtureResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *EventsFixtureStore) MustFind(q *EventsFixtureQuery) *EventsFixtureResultSet {
	return NewEventsFixtureResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *EventsFixtureStore) Count(q *EventsFixtureQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *EventsFixtureStore) MustCount(q *EventsFixtureQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `sql.ErrNoRows` is returned if there are no results.
func (s *EventsFixtureStore) FindOne(q *EventsFixtureQuery) (*EventsFixture, error) {
	q.Limit(1)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *EventsFixtureStore) MustFindOne(q *EventsFixtureQuery) *EventsFixture {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the EventsFixture with the data in the database and
// makes it writable.
func (s *EventsFixtureStore) Reload(record *EventsFixture) error {
	return s.Store.Reload(record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *EventsFixtureStore) Transaction(callback func(*EventsFixtureStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&EventsFixtureStore{store})
	})
}

// EventsFixtureQuery is the object used to create queries for the EventsFixture
// entity.
type EventsFixtureQuery struct {
	*kallax.BaseQuery
}

// NewEventsFixtureQuery returns a new instance of EventsFixtureQuery.
func NewEventsFixtureQuery() *EventsFixtureQuery {
	return &EventsFixtureQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.EventsFixture.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *EventsFixtureQuery) Select(columns ...kallax.SchemaField) *EventsFixtureQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *EventsFixtureQuery) SelectNot(columns ...kallax.SchemaField) *EventsFixtureQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *EventsFixtureQuery) Copy() *EventsFixtureQuery {
	return &EventsFixtureQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *EventsFixtureQuery) Order(cols ...kallax.ColumnOrder) *EventsFixtureQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *EventsFixtureQuery) BatchSize(size uint64) *EventsFixtureQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *EventsFixtureQuery) Limit(n uint64) *EventsFixtureQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *EventsFixtureQuery) Offset(n uint64) *EventsFixtureQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *EventsFixtureQuery) Where(cond kallax.Condition) *EventsFixtureQuery {
	q.BaseQuery.Where(cond)
	return q
}

// EventsFixtureResultSet is the set of results returned by a query to the
// database.
type EventsFixtureResultSet struct {
	*kallax.ResultSet
	last    *EventsFixture
	lastErr error
}

func NewEventsFixtureResultSet(rs *kallax.ResultSet) *EventsFixtureResultSet {
	return &EventsFixtureResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *EventsFixtureResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.Close()
		return false
	}

	rs.last = new(EventsFixture)
	rs.lastErr = rs.Scan(rs.last)
	if rs.lastErr != nil {
		rs.last = nil
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *EventsFixtureResultSet) Get() (*EventsFixture, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *EventsFixtureResultSet) ForEach(fn func(*EventsFixture) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *EventsFixtureResultSet) All() ([]*EventsFixture, error) {
	var result []*EventsFixture
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *EventsFixtureResultSet) One() (*EventsFixture, error) {
	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// NewEventsSaveFixture returns a new instance of EventsSaveFixture.
func NewEventsSaveFixture() (record *EventsSaveFixture) {
	record = newEventsSaveFixture()
	if record != nil {
		record.SetID(kallax.NewID())
	}
	return
}

func (r *EventsSaveFixture) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return &r.Model.ID, nil
	case "checks":
		return types.JSON(&r.Checks), nil
	case "must_fail_before":
		return types.JSON(&r.MustFailBefore), nil
	case "must_fail_after":
		return types.JSON(&r.MustFailAfter), nil

	default:
		return nil, fmt.Errorf("invalid column in EventsSaveFixture: %s", col)
	}
}

func (r *EventsSaveFixture) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.Model.ID, nil
	case "checks":
		return types.JSON(r.Checks), nil
	case "must_fail_before":
		return types.JSON(r.MustFailBefore), nil
	case "must_fail_after":
		return types.JSON(r.MustFailAfter), nil

	default:
		return nil, fmt.Errorf("invalid column in EventsSaveFixture: %s", col)
	}
}

func (r *EventsSaveFixture) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {

	}
	return nil, fmt.Errorf("model EventsSaveFixture has no relationship %s", field)
}

func (r *EventsSaveFixture) SetRelationship(field string, record kallax.Record) error {
	switch field {

	}
	return fmt.Errorf("model EventsSaveFixture has no relationship %s", field)
}

// EventsSaveFixtureStore is the entity to access the records of the type EventsSaveFixture
// in the database.
type EventsSaveFixtureStore struct {
	*kallax.Store
}

// NewEventsSaveFixtureStore creates a new instance of EventsSaveFixtureStore
// using a SQL database.
func NewEventsSaveFixtureStore(db *sql.DB) *EventsSaveFixtureStore {
	return &EventsSaveFixtureStore{kallax.NewStore(db, Schema.EventsSaveFixture.BaseSchema)}
}

// Insert inserts a EventsSaveFixture in the database. A non-persisted object is
// required for this operation.
func (s *EventsSaveFixtureStore) Insert(record *EventsSaveFixture) error {

	return s.Store.Insert(record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *EventsSaveFixtureStore) Update(record *EventsSaveFixture, cols ...kallax.SchemaField) (int64, error) {

	return s.Store.Update(record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *EventsSaveFixtureStore) Save(record *EventsSaveFixture) (updated bool, err error) {

	return s.Store.Save(record)
}

// Delete removes the given record from the database.
func (s *EventsSaveFixtureStore) Delete(record *EventsSaveFixture) error {
	return s.Store.Delete(record)
}

// Find returns the set of results for the given query.
func (s *EventsSaveFixtureStore) Find(q *EventsSaveFixtureQuery) (*EventsSaveFixtureResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewEventsSaveFixtureResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *EventsSaveFixtureStore) MustFind(q *EventsSaveFixtureQuery) *EventsSaveFixtureResultSet {
	return NewEventsSaveFixtureResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *EventsSaveFixtureStore) Count(q *EventsSaveFixtureQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *EventsSaveFixtureStore) MustCount(q *EventsSaveFixtureQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `sql.ErrNoRows` is returned if there are no results.
func (s *EventsSaveFixtureStore) FindOne(q *EventsSaveFixtureQuery) (*EventsSaveFixture, error) {
	q.Limit(1)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *EventsSaveFixtureStore) MustFindOne(q *EventsSaveFixtureQuery) *EventsSaveFixture {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the EventsSaveFixture with the data in the database and
// makes it writable.
func (s *EventsSaveFixtureStore) Reload(record *EventsSaveFixture) error {
	return s.Store.Reload(record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *EventsSaveFixtureStore) Transaction(callback func(*EventsSaveFixtureStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&EventsSaveFixtureStore{store})
	})
}

// EventsSaveFixtureQuery is the object used to create queries for the EventsSaveFixture
// entity.
type EventsSaveFixtureQuery struct {
	*kallax.BaseQuery
}

// NewEventsSaveFixtureQuery returns a new instance of EventsSaveFixtureQuery.
func NewEventsSaveFixtureQuery() *EventsSaveFixtureQuery {
	return &EventsSaveFixtureQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.EventsSaveFixture.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *EventsSaveFixtureQuery) Select(columns ...kallax.SchemaField) *EventsSaveFixtureQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *EventsSaveFixtureQuery) SelectNot(columns ...kallax.SchemaField) *EventsSaveFixtureQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *EventsSaveFixtureQuery) Copy() *EventsSaveFixtureQuery {
	return &EventsSaveFixtureQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *EventsSaveFixtureQuery) Order(cols ...kallax.ColumnOrder) *EventsSaveFixtureQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *EventsSaveFixtureQuery) BatchSize(size uint64) *EventsSaveFixtureQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *EventsSaveFixtureQuery) Limit(n uint64) *EventsSaveFixtureQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *EventsSaveFixtureQuery) Offset(n uint64) *EventsSaveFixtureQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *EventsSaveFixtureQuery) Where(cond kallax.Condition) *EventsSaveFixtureQuery {
	q.BaseQuery.Where(cond)
	return q
}

// EventsSaveFixtureResultSet is the set of results returned by a query to the
// database.
type EventsSaveFixtureResultSet struct {
	*kallax.ResultSet
	last    *EventsSaveFixture
	lastErr error
}

func NewEventsSaveFixtureResultSet(rs *kallax.ResultSet) *EventsSaveFixtureResultSet {
	return &EventsSaveFixtureResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *EventsSaveFixtureResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.Close()
		return false
	}

	rs.last = new(EventsSaveFixture)
	rs.lastErr = rs.Scan(rs.last)
	if rs.lastErr != nil {
		rs.last = nil
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *EventsSaveFixtureResultSet) Get() (*EventsSaveFixture, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *EventsSaveFixtureResultSet) ForEach(fn func(*EventsSaveFixture) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *EventsSaveFixtureResultSet) All() ([]*EventsSaveFixture, error) {
	var result []*EventsSaveFixture
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *EventsSaveFixtureResultSet) One() (*EventsSaveFixture, error) {
	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// NewMultiKeySortFixture returns a new instance of MultiKeySortFixture.
func NewMultiKeySortFixture() (record *MultiKeySortFixture) {
	record = &MultiKeySortFixture{}
	if record != nil {
		record.SetID(kallax.NewID())
	}
	return
}

func (r *MultiKeySortFixture) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return &r.Model.ID, nil
	case "name":
		return &r.Name, nil
	case "start":
		return &r.Start, nil
	case "_end":
		return &r.End, nil

	default:
		return nil, fmt.Errorf("invalid column in MultiKeySortFixture: %s", col)
	}
}

func (r *MultiKeySortFixture) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.Model.ID, nil
	case "name":
		return r.Name, nil
	case "start":
		return r.Start, nil
	case "_end":
		return r.End, nil

	default:
		return nil, fmt.Errorf("invalid column in MultiKeySortFixture: %s", col)
	}
}

func (r *MultiKeySortFixture) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {

	}
	return nil, fmt.Errorf("model MultiKeySortFixture has no relationship %s", field)
}

func (r *MultiKeySortFixture) SetRelationship(field string, record kallax.Record) error {
	switch field {

	}
	return fmt.Errorf("model MultiKeySortFixture has no relationship %s", field)
}

// MultiKeySortFixtureStore is the entity to access the records of the type MultiKeySortFixture
// in the database.
type MultiKeySortFixtureStore struct {
	*kallax.Store
}

// NewMultiKeySortFixtureStore creates a new instance of MultiKeySortFixtureStore
// using a SQL database.
func NewMultiKeySortFixtureStore(db *sql.DB) *MultiKeySortFixtureStore {
	return &MultiKeySortFixtureStore{kallax.NewStore(db, Schema.MultiKeySortFixture.BaseSchema)}
}

// Insert inserts a MultiKeySortFixture in the database. A non-persisted object is
// required for this operation.
func (s *MultiKeySortFixtureStore) Insert(record *MultiKeySortFixture) error {

	return s.Store.Insert(record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *MultiKeySortFixtureStore) Update(record *MultiKeySortFixture, cols ...kallax.SchemaField) (int64, error) {

	return s.Store.Update(record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *MultiKeySortFixtureStore) Save(record *MultiKeySortFixture) (updated bool, err error) {

	return s.Store.Save(record)
}

// Delete removes the given record from the database.
func (s *MultiKeySortFixtureStore) Delete(record *MultiKeySortFixture) error {
	return s.Store.Delete(record)
}

// Find returns the set of results for the given query.
func (s *MultiKeySortFixtureStore) Find(q *MultiKeySortFixtureQuery) (*MultiKeySortFixtureResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewMultiKeySortFixtureResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *MultiKeySortFixtureStore) MustFind(q *MultiKeySortFixtureQuery) *MultiKeySortFixtureResultSet {
	return NewMultiKeySortFixtureResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *MultiKeySortFixtureStore) Count(q *MultiKeySortFixtureQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *MultiKeySortFixtureStore) MustCount(q *MultiKeySortFixtureQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `sql.ErrNoRows` is returned if there are no results.
func (s *MultiKeySortFixtureStore) FindOne(q *MultiKeySortFixtureQuery) (*MultiKeySortFixture, error) {
	q.Limit(1)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *MultiKeySortFixtureStore) MustFindOne(q *MultiKeySortFixtureQuery) *MultiKeySortFixture {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the MultiKeySortFixture with the data in the database and
// makes it writable.
func (s *MultiKeySortFixtureStore) Reload(record *MultiKeySortFixture) error {
	return s.Store.Reload(record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *MultiKeySortFixtureStore) Transaction(callback func(*MultiKeySortFixtureStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&MultiKeySortFixtureStore{store})
	})
}

// MultiKeySortFixtureQuery is the object used to create queries for the MultiKeySortFixture
// entity.
type MultiKeySortFixtureQuery struct {
	*kallax.BaseQuery
}

// NewMultiKeySortFixtureQuery returns a new instance of MultiKeySortFixtureQuery.
func NewMultiKeySortFixtureQuery() *MultiKeySortFixtureQuery {
	return &MultiKeySortFixtureQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.MultiKeySortFixture.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *MultiKeySortFixtureQuery) Select(columns ...kallax.SchemaField) *MultiKeySortFixtureQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *MultiKeySortFixtureQuery) SelectNot(columns ...kallax.SchemaField) *MultiKeySortFixtureQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *MultiKeySortFixtureQuery) Copy() *MultiKeySortFixtureQuery {
	return &MultiKeySortFixtureQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *MultiKeySortFixtureQuery) Order(cols ...kallax.ColumnOrder) *MultiKeySortFixtureQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *MultiKeySortFixtureQuery) BatchSize(size uint64) *MultiKeySortFixtureQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *MultiKeySortFixtureQuery) Limit(n uint64) *MultiKeySortFixtureQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *MultiKeySortFixtureQuery) Offset(n uint64) *MultiKeySortFixtureQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *MultiKeySortFixtureQuery) Where(cond kallax.Condition) *MultiKeySortFixtureQuery {
	q.BaseQuery.Where(cond)
	return q
}

// MultiKeySortFixtureResultSet is the set of results returned by a query to the
// database.
type MultiKeySortFixtureResultSet struct {
	*kallax.ResultSet
	last    *MultiKeySortFixture
	lastErr error
}

func NewMultiKeySortFixtureResultSet(rs *kallax.ResultSet) *MultiKeySortFixtureResultSet {
	return &MultiKeySortFixtureResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *MultiKeySortFixtureResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.Close()
		return false
	}

	rs.last = new(MultiKeySortFixture)
	rs.lastErr = rs.Scan(rs.last)
	if rs.lastErr != nil {
		rs.last = nil
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *MultiKeySortFixtureResultSet) Get() (*MultiKeySortFixture, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *MultiKeySortFixtureResultSet) ForEach(fn func(*MultiKeySortFixture) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *MultiKeySortFixtureResultSet) All() ([]*MultiKeySortFixture, error) {
	var result []*MultiKeySortFixture
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *MultiKeySortFixtureResultSet) One() (*MultiKeySortFixture, error) {
	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// NewQueryFixture returns a new instance of QueryFixture.
func NewQueryFixture(f string) (record *QueryFixture) {
	record = newQueryFixture(f)
	if record != nil {
		record.SetID(kallax.NewID())
	}
	return
}

func (r *QueryFixture) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return &r.Model.ID, nil
	case "foo":
		return &r.Foo, nil

	default:
		return nil, fmt.Errorf("invalid column in QueryFixture: %s", col)
	}
}

func (r *QueryFixture) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.Model.ID, nil
	case "foo":
		return r.Foo, nil

	default:
		return nil, fmt.Errorf("invalid column in QueryFixture: %s", col)
	}
}

func (r *QueryFixture) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {

	}
	return nil, fmt.Errorf("model QueryFixture has no relationship %s", field)
}

func (r *QueryFixture) SetRelationship(field string, record kallax.Record) error {
	switch field {

	}
	return fmt.Errorf("model QueryFixture has no relationship %s", field)
}

// QueryFixtureStore is the entity to access the records of the type QueryFixture
// in the database.
type QueryFixtureStore struct {
	*kallax.Store
}

// NewQueryFixtureStore creates a new instance of QueryFixtureStore
// using a SQL database.
func NewQueryFixtureStore(db *sql.DB) *QueryFixtureStore {
	return &QueryFixtureStore{kallax.NewStore(db, Schema.QueryFixture.BaseSchema)}
}

// Insert inserts a QueryFixture in the database. A non-persisted object is
// required for this operation.
func (s *QueryFixtureStore) Insert(record *QueryFixture) error {

	return s.Store.Insert(record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *QueryFixtureStore) Update(record *QueryFixture, cols ...kallax.SchemaField) (int64, error) {

	return s.Store.Update(record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *QueryFixtureStore) Save(record *QueryFixture) (updated bool, err error) {

	return s.Store.Save(record)
}

// Delete removes the given record from the database.
func (s *QueryFixtureStore) Delete(record *QueryFixture) error {
	return s.Store.Delete(record)
}

// Find returns the set of results for the given query.
func (s *QueryFixtureStore) Find(q *QueryFixtureQuery) (*QueryFixtureResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewQueryFixtureResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *QueryFixtureStore) MustFind(q *QueryFixtureQuery) *QueryFixtureResultSet {
	return NewQueryFixtureResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *QueryFixtureStore) Count(q *QueryFixtureQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *QueryFixtureStore) MustCount(q *QueryFixtureQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `sql.ErrNoRows` is returned if there are no results.
func (s *QueryFixtureStore) FindOne(q *QueryFixtureQuery) (*QueryFixture, error) {
	q.Limit(1)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *QueryFixtureStore) MustFindOne(q *QueryFixtureQuery) *QueryFixture {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the QueryFixture with the data in the database and
// makes it writable.
func (s *QueryFixtureStore) Reload(record *QueryFixture) error {
	return s.Store.Reload(record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *QueryFixtureStore) Transaction(callback func(*QueryFixtureStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&QueryFixtureStore{store})
	})
}

// QueryFixtureQuery is the object used to create queries for the QueryFixture
// entity.
type QueryFixtureQuery struct {
	*kallax.BaseQuery
}

// NewQueryFixtureQuery returns a new instance of QueryFixtureQuery.
func NewQueryFixtureQuery() *QueryFixtureQuery {
	return &QueryFixtureQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.QueryFixture.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *QueryFixtureQuery) Select(columns ...kallax.SchemaField) *QueryFixtureQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *QueryFixtureQuery) SelectNot(columns ...kallax.SchemaField) *QueryFixtureQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *QueryFixtureQuery) Copy() *QueryFixtureQuery {
	return &QueryFixtureQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *QueryFixtureQuery) Order(cols ...kallax.ColumnOrder) *QueryFixtureQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *QueryFixtureQuery) BatchSize(size uint64) *QueryFixtureQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *QueryFixtureQuery) Limit(n uint64) *QueryFixtureQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *QueryFixtureQuery) Offset(n uint64) *QueryFixtureQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *QueryFixtureQuery) Where(cond kallax.Condition) *QueryFixtureQuery {
	q.BaseQuery.Where(cond)
	return q
}

// QueryFixtureResultSet is the set of results returned by a query to the
// database.
type QueryFixtureResultSet struct {
	*kallax.ResultSet
	last    *QueryFixture
	lastErr error
}

func NewQueryFixtureResultSet(rs *kallax.ResultSet) *QueryFixtureResultSet {
	return &QueryFixtureResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *QueryFixtureResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.Close()
		return false
	}

	rs.last = new(QueryFixture)
	rs.lastErr = rs.Scan(rs.last)
	if rs.lastErr != nil {
		rs.last = nil
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *QueryFixtureResultSet) Get() (*QueryFixture, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *QueryFixtureResultSet) ForEach(fn func(*QueryFixture) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *QueryFixtureResultSet) All() ([]*QueryFixture, error) {
	var result []*QueryFixture
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *QueryFixtureResultSet) One() (*QueryFixture, error) {
	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// NewResultSetFixture returns a new instance of ResultSetFixture.
func NewResultSetFixture(f string) (record *ResultSetFixture) {
	record = newResultSetFixture(f)
	if record != nil {
		record.SetID(kallax.NewID())
	}
	return
}

func (r *ResultSetFixture) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return &r.Model.ID, nil
	case "foo":
		return &r.Foo, nil

	default:
		return nil, fmt.Errorf("invalid column in ResultSetFixture: %s", col)
	}
}

func (r *ResultSetFixture) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.Model.ID, nil
	case "foo":
		return r.Foo, nil

	default:
		return nil, fmt.Errorf("invalid column in ResultSetFixture: %s", col)
	}
}

func (r *ResultSetFixture) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {

	}
	return nil, fmt.Errorf("model ResultSetFixture has no relationship %s", field)
}

func (r *ResultSetFixture) SetRelationship(field string, record kallax.Record) error {
	switch field {

	}
	return fmt.Errorf("model ResultSetFixture has no relationship %s", field)
}

// ResultSetFixtureStore is the entity to access the records of the type ResultSetFixture
// in the database.
type ResultSetFixtureStore struct {
	*kallax.Store
}

// NewResultSetFixtureStore creates a new instance of ResultSetFixtureStore
// using a SQL database.
func NewResultSetFixtureStore(db *sql.DB) *ResultSetFixtureStore {
	return &ResultSetFixtureStore{kallax.NewStore(db, Schema.ResultSetFixture.BaseSchema)}
}

// Insert inserts a ResultSetFixture in the database. A non-persisted object is
// required for this operation.
func (s *ResultSetFixtureStore) Insert(record *ResultSetFixture) error {

	return s.Store.Insert(record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *ResultSetFixtureStore) Update(record *ResultSetFixture, cols ...kallax.SchemaField) (int64, error) {

	return s.Store.Update(record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *ResultSetFixtureStore) Save(record *ResultSetFixture) (updated bool, err error) {

	return s.Store.Save(record)
}

// Delete removes the given record from the database.
func (s *ResultSetFixtureStore) Delete(record *ResultSetFixture) error {
	return s.Store.Delete(record)
}

// Find returns the set of results for the given query.
func (s *ResultSetFixtureStore) Find(q *ResultSetFixtureQuery) (*ResultSetFixtureResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewResultSetFixtureResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *ResultSetFixtureStore) MustFind(q *ResultSetFixtureQuery) *ResultSetFixtureResultSet {
	return NewResultSetFixtureResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *ResultSetFixtureStore) Count(q *ResultSetFixtureQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *ResultSetFixtureStore) MustCount(q *ResultSetFixtureQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `sql.ErrNoRows` is returned if there are no results.
func (s *ResultSetFixtureStore) FindOne(q *ResultSetFixtureQuery) (*ResultSetFixture, error) {
	q.Limit(1)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *ResultSetFixtureStore) MustFindOne(q *ResultSetFixtureQuery) *ResultSetFixture {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the ResultSetFixture with the data in the database and
// makes it writable.
func (s *ResultSetFixtureStore) Reload(record *ResultSetFixture) error {
	return s.Store.Reload(record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *ResultSetFixtureStore) Transaction(callback func(*ResultSetFixtureStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&ResultSetFixtureStore{store})
	})
}

// ResultSetFixtureQuery is the object used to create queries for the ResultSetFixture
// entity.
type ResultSetFixtureQuery struct {
	*kallax.BaseQuery
}

// NewResultSetFixtureQuery returns a new instance of ResultSetFixtureQuery.
func NewResultSetFixtureQuery() *ResultSetFixtureQuery {
	return &ResultSetFixtureQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.ResultSetFixture.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *ResultSetFixtureQuery) Select(columns ...kallax.SchemaField) *ResultSetFixtureQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *ResultSetFixtureQuery) SelectNot(columns ...kallax.SchemaField) *ResultSetFixtureQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *ResultSetFixtureQuery) Copy() *ResultSetFixtureQuery {
	return &ResultSetFixtureQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *ResultSetFixtureQuery) Order(cols ...kallax.ColumnOrder) *ResultSetFixtureQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *ResultSetFixtureQuery) BatchSize(size uint64) *ResultSetFixtureQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *ResultSetFixtureQuery) Limit(n uint64) *ResultSetFixtureQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *ResultSetFixtureQuery) Offset(n uint64) *ResultSetFixtureQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *ResultSetFixtureQuery) Where(cond kallax.Condition) *ResultSetFixtureQuery {
	q.BaseQuery.Where(cond)
	return q
}

// ResultSetFixtureResultSet is the set of results returned by a query to the
// database.
type ResultSetFixtureResultSet struct {
	*kallax.ResultSet
	last    *ResultSetFixture
	lastErr error
}

func NewResultSetFixtureResultSet(rs *kallax.ResultSet) *ResultSetFixtureResultSet {
	return &ResultSetFixtureResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *ResultSetFixtureResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.Close()
		return false
	}

	rs.last = new(ResultSetFixture)
	rs.lastErr = rs.Scan(rs.last)
	if rs.lastErr != nil {
		rs.last = nil
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *ResultSetFixtureResultSet) Get() (*ResultSetFixture, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *ResultSetFixtureResultSet) ForEach(fn func(*ResultSetFixture) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *ResultSetFixtureResultSet) All() ([]*ResultSetFixture, error) {
	var result []*ResultSetFixture
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *ResultSetFixtureResultSet) One() (*ResultSetFixture, error) {
	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// NewResultSetInitFixture returns a new instance of ResultSetInitFixture.
func NewResultSetInitFixture() (record *ResultSetInitFixture) {
	record = &ResultSetInitFixture{}
	if record != nil {
		record.SetID(kallax.NewID())
	}
	return
}

func (r *ResultSetInitFixture) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return &r.Model.ID, nil
	case "foo":
		return &r.Foo, nil

	default:
		return nil, fmt.Errorf("invalid column in ResultSetInitFixture: %s", col)
	}
}

func (r *ResultSetInitFixture) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.Model.ID, nil
	case "foo":
		return r.Foo, nil

	default:
		return nil, fmt.Errorf("invalid column in ResultSetInitFixture: %s", col)
	}
}

func (r *ResultSetInitFixture) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {

	}
	return nil, fmt.Errorf("model ResultSetInitFixture has no relationship %s", field)
}

func (r *ResultSetInitFixture) SetRelationship(field string, record kallax.Record) error {
	switch field {

	}
	return fmt.Errorf("model ResultSetInitFixture has no relationship %s", field)
}

// ResultSetInitFixtureStore is the entity to access the records of the type ResultSetInitFixture
// in the database.
type ResultSetInitFixtureStore struct {
	*kallax.Store
}

// NewResultSetInitFixtureStore creates a new instance of ResultSetInitFixtureStore
// using a SQL database.
func NewResultSetInitFixtureStore(db *sql.DB) *ResultSetInitFixtureStore {
	return &ResultSetInitFixtureStore{kallax.NewStore(db, Schema.ResultSetInitFixture.BaseSchema)}
}

// Insert inserts a ResultSetInitFixture in the database. A non-persisted object is
// required for this operation.
func (s *ResultSetInitFixtureStore) Insert(record *ResultSetInitFixture) error {

	return s.Store.Insert(record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *ResultSetInitFixtureStore) Update(record *ResultSetInitFixture, cols ...kallax.SchemaField) (int64, error) {

	return s.Store.Update(record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *ResultSetInitFixtureStore) Save(record *ResultSetInitFixture) (updated bool, err error) {

	return s.Store.Save(record)
}

// Delete removes the given record from the database.
func (s *ResultSetInitFixtureStore) Delete(record *ResultSetInitFixture) error {
	return s.Store.Delete(record)
}

// Find returns the set of results for the given query.
func (s *ResultSetInitFixtureStore) Find(q *ResultSetInitFixtureQuery) (*ResultSetInitFixtureResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewResultSetInitFixtureResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *ResultSetInitFixtureStore) MustFind(q *ResultSetInitFixtureQuery) *ResultSetInitFixtureResultSet {
	return NewResultSetInitFixtureResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *ResultSetInitFixtureStore) Count(q *ResultSetInitFixtureQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *ResultSetInitFixtureStore) MustCount(q *ResultSetInitFixtureQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `sql.ErrNoRows` is returned if there are no results.
func (s *ResultSetInitFixtureStore) FindOne(q *ResultSetInitFixtureQuery) (*ResultSetInitFixture, error) {
	q.Limit(1)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *ResultSetInitFixtureStore) MustFindOne(q *ResultSetInitFixtureQuery) *ResultSetInitFixture {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the ResultSetInitFixture with the data in the database and
// makes it writable.
func (s *ResultSetInitFixtureStore) Reload(record *ResultSetInitFixture) error {
	return s.Store.Reload(record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *ResultSetInitFixtureStore) Transaction(callback func(*ResultSetInitFixtureStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&ResultSetInitFixtureStore{store})
	})
}

// ResultSetInitFixtureQuery is the object used to create queries for the ResultSetInitFixture
// entity.
type ResultSetInitFixtureQuery struct {
	*kallax.BaseQuery
}

// NewResultSetInitFixtureQuery returns a new instance of ResultSetInitFixtureQuery.
func NewResultSetInitFixtureQuery() *ResultSetInitFixtureQuery {
	return &ResultSetInitFixtureQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.ResultSetInitFixture.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *ResultSetInitFixtureQuery) Select(columns ...kallax.SchemaField) *ResultSetInitFixtureQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *ResultSetInitFixtureQuery) SelectNot(columns ...kallax.SchemaField) *ResultSetInitFixtureQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *ResultSetInitFixtureQuery) Copy() *ResultSetInitFixtureQuery {
	return &ResultSetInitFixtureQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *ResultSetInitFixtureQuery) Order(cols ...kallax.ColumnOrder) *ResultSetInitFixtureQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *ResultSetInitFixtureQuery) BatchSize(size uint64) *ResultSetInitFixtureQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *ResultSetInitFixtureQuery) Limit(n uint64) *ResultSetInitFixtureQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *ResultSetInitFixtureQuery) Offset(n uint64) *ResultSetInitFixtureQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *ResultSetInitFixtureQuery) Where(cond kallax.Condition) *ResultSetInitFixtureQuery {
	q.BaseQuery.Where(cond)
	return q
}

// ResultSetInitFixtureResultSet is the set of results returned by a query to the
// database.
type ResultSetInitFixtureResultSet struct {
	*kallax.ResultSet
	last    *ResultSetInitFixture
	lastErr error
}

func NewResultSetInitFixtureResultSet(rs *kallax.ResultSet) *ResultSetInitFixtureResultSet {
	return &ResultSetInitFixtureResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *ResultSetInitFixtureResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.Close()
		return false
	}

	rs.last = new(ResultSetInitFixture)
	rs.lastErr = rs.Scan(rs.last)
	if rs.lastErr != nil {
		rs.last = nil
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *ResultSetInitFixtureResultSet) Get() (*ResultSetInitFixture, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *ResultSetInitFixtureResultSet) ForEach(fn func(*ResultSetInitFixture) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *ResultSetInitFixtureResultSet) All() ([]*ResultSetInitFixture, error) {
	var result []*ResultSetInitFixture
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *ResultSetInitFixtureResultSet) One() (*ResultSetInitFixture, error) {
	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// NewSchemaFixture returns a new instance of SchemaFixture.
func NewSchemaFixture() (record *SchemaFixture) {
	record = &SchemaFixture{}
	if record != nil {
		record.SetID(kallax.NewID())
	}
	return
}

func (r *SchemaFixture) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return &r.Model.ID, nil
	case "string":
		return &r.String, nil
	case "int":
		return &r.Int, nil
	case "inline":
		return &r.Inline.Inline, nil
	case "map_of_string":
		return types.JSON(&r.MapOfString), nil
	case "map_of_interface":
		return types.JSON(&r.MapOfInterface), nil
	case "map_of_some_type":
		return types.JSON(&r.MapOfSomeType), nil

	default:
		return nil, fmt.Errorf("invalid column in SchemaFixture: %s", col)
	}
}

func (r *SchemaFixture) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.Model.ID, nil
	case "string":
		return r.String, nil
	case "int":
		return r.Int, nil
	case "inline":
		return r.Inline.Inline, nil
	case "map_of_string":
		return types.JSON(r.MapOfString), nil
	case "map_of_interface":
		return types.JSON(r.MapOfInterface), nil
	case "map_of_some_type":
		return types.JSON(r.MapOfSomeType), nil

	default:
		return nil, fmt.Errorf("invalid column in SchemaFixture: %s", col)
	}
}

func (r *SchemaFixture) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "Nested":
		return &SchemaFixture{}, nil

	}
	return nil, fmt.Errorf("model SchemaFixture has no relationship %s", field)
}

func (r *SchemaFixture) SetRelationship(field string, record kallax.Record) error {
	switch field {
	case "Nested":
		val, ok := record.(*SchemaFixture)
		if !ok {
			return fmt.Errorf("record of type %t can't be assigned to relationship Nested", record)
		}
		r.Nested = val
		return nil

	}
	return fmt.Errorf("model SchemaFixture has no relationship %s", field)
}

// SchemaFixtureStore is the entity to access the records of the type SchemaFixture
// in the database.
type SchemaFixtureStore struct {
	*kallax.Store
}

// NewSchemaFixtureStore creates a new instance of SchemaFixtureStore
// using a SQL database.
func NewSchemaFixtureStore(db *sql.DB) *SchemaFixtureStore {
	return &SchemaFixtureStore{kallax.NewStore(db, Schema.SchemaFixture.BaseSchema)}
}

// Insert inserts a SchemaFixture in the database. A non-persisted object is
// required for this operation.
func (s *SchemaFixtureStore) Insert(record *SchemaFixture) error {

	return s.Store.Insert(record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *SchemaFixtureStore) Update(record *SchemaFixture, cols ...kallax.SchemaField) (int64, error) {

	return s.Store.Update(record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *SchemaFixtureStore) Save(record *SchemaFixture) (updated bool, err error) {

	return s.Store.Save(record)
}

// Delete removes the given record from the database.
func (s *SchemaFixtureStore) Delete(record *SchemaFixture) error {
	return s.Store.Delete(record)
}

// Find returns the set of results for the given query.
func (s *SchemaFixtureStore) Find(q *SchemaFixtureQuery) (*SchemaFixtureResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewSchemaFixtureResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *SchemaFixtureStore) MustFind(q *SchemaFixtureQuery) *SchemaFixtureResultSet {
	return NewSchemaFixtureResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *SchemaFixtureStore) Count(q *SchemaFixtureQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *SchemaFixtureStore) MustCount(q *SchemaFixtureQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `sql.ErrNoRows` is returned if there are no results.
func (s *SchemaFixtureStore) FindOne(q *SchemaFixtureQuery) (*SchemaFixture, error) {
	q.Limit(1)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *SchemaFixtureStore) MustFindOne(q *SchemaFixtureQuery) *SchemaFixture {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the SchemaFixture with the data in the database and
// makes it writable.
func (s *SchemaFixtureStore) Reload(record *SchemaFixture) error {
	return s.Store.Reload(record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *SchemaFixtureStore) Transaction(callback func(*SchemaFixtureStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&SchemaFixtureStore{store})
	})
}

// SchemaFixtureQuery is the object used to create queries for the SchemaFixture
// entity.
type SchemaFixtureQuery struct {
	*kallax.BaseQuery
}

// NewSchemaFixtureQuery returns a new instance of SchemaFixtureQuery.
func NewSchemaFixtureQuery() *SchemaFixtureQuery {
	return &SchemaFixtureQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.SchemaFixture.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *SchemaFixtureQuery) Select(columns ...kallax.SchemaField) *SchemaFixtureQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *SchemaFixtureQuery) SelectNot(columns ...kallax.SchemaField) *SchemaFixtureQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *SchemaFixtureQuery) Copy() *SchemaFixtureQuery {
	return &SchemaFixtureQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *SchemaFixtureQuery) Order(cols ...kallax.ColumnOrder) *SchemaFixtureQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *SchemaFixtureQuery) BatchSize(size uint64) *SchemaFixtureQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *SchemaFixtureQuery) Limit(n uint64) *SchemaFixtureQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *SchemaFixtureQuery) Offset(n uint64) *SchemaFixtureQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *SchemaFixtureQuery) Where(cond kallax.Condition) *SchemaFixtureQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *SchemaFixtureQuery) WithNested() *SchemaFixtureQuery {
	q.AddRelation(Schema.SchemaFixture.BaseSchema, "Nested")
	return q
}

// SchemaFixtureResultSet is the set of results returned by a query to the
// database.
type SchemaFixtureResultSet struct {
	*kallax.ResultSet
	last    *SchemaFixture
	lastErr error
}

func NewSchemaFixtureResultSet(rs *kallax.ResultSet) *SchemaFixtureResultSet {
	return &SchemaFixtureResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *SchemaFixtureResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.Close()
		return false
	}

	rs.last = new(SchemaFixture)
	rs.lastErr = rs.Scan(rs.last)
	if rs.lastErr != nil {
		rs.last = nil
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *SchemaFixtureResultSet) Get() (*SchemaFixture, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *SchemaFixtureResultSet) ForEach(fn func(*SchemaFixture) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *SchemaFixtureResultSet) All() ([]*SchemaFixture, error) {
	var result []*SchemaFixture
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *SchemaFixtureResultSet) One() (*SchemaFixture, error) {
	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// NewStoreFixture returns a new instance of StoreFixture.
func NewStoreFixture() (record *StoreFixture) {
	record = &StoreFixture{}
	if record != nil {
		record.SetID(kallax.NewID())
	}
	return
}

func (r *StoreFixture) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return &r.Model.ID, nil
	case "foo":
		return &r.Foo, nil

	default:
		return nil, fmt.Errorf("invalid column in StoreFixture: %s", col)
	}
}

func (r *StoreFixture) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.Model.ID, nil
	case "foo":
		return r.Foo, nil

	default:
		return nil, fmt.Errorf("invalid column in StoreFixture: %s", col)
	}
}

func (r *StoreFixture) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {

	}
	return nil, fmt.Errorf("model StoreFixture has no relationship %s", field)
}

func (r *StoreFixture) SetRelationship(field string, record kallax.Record) error {
	switch field {

	}
	return fmt.Errorf("model StoreFixture has no relationship %s", field)
}

// StoreFixtureStore is the entity to access the records of the type StoreFixture
// in the database.
type StoreFixtureStore struct {
	*kallax.Store
}

// NewStoreFixtureStore creates a new instance of StoreFixtureStore
// using a SQL database.
func NewStoreFixtureStore(db *sql.DB) *StoreFixtureStore {
	return &StoreFixtureStore{kallax.NewStore(db, Schema.StoreFixture.BaseSchema)}
}

// Insert inserts a StoreFixture in the database. A non-persisted object is
// required for this operation.
func (s *StoreFixtureStore) Insert(record *StoreFixture) error {

	return s.Store.Insert(record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *StoreFixtureStore) Update(record *StoreFixture, cols ...kallax.SchemaField) (int64, error) {

	return s.Store.Update(record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *StoreFixtureStore) Save(record *StoreFixture) (updated bool, err error) {

	return s.Store.Save(record)
}

// Delete removes the given record from the database.
func (s *StoreFixtureStore) Delete(record *StoreFixture) error {
	return s.Store.Delete(record)
}

// Find returns the set of results for the given query.
func (s *StoreFixtureStore) Find(q *StoreFixtureQuery) (*StoreFixtureResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewStoreFixtureResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *StoreFixtureStore) MustFind(q *StoreFixtureQuery) *StoreFixtureResultSet {
	return NewStoreFixtureResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *StoreFixtureStore) Count(q *StoreFixtureQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *StoreFixtureStore) MustCount(q *StoreFixtureQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `sql.ErrNoRows` is returned if there are no results.
func (s *StoreFixtureStore) FindOne(q *StoreFixtureQuery) (*StoreFixture, error) {
	q.Limit(1)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *StoreFixtureStore) MustFindOne(q *StoreFixtureQuery) *StoreFixture {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the StoreFixture with the data in the database and
// makes it writable.
func (s *StoreFixtureStore) Reload(record *StoreFixture) error {
	return s.Store.Reload(record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *StoreFixtureStore) Transaction(callback func(*StoreFixtureStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&StoreFixtureStore{store})
	})
}

// StoreFixtureQuery is the object used to create queries for the StoreFixture
// entity.
type StoreFixtureQuery struct {
	*kallax.BaseQuery
}

// NewStoreFixtureQuery returns a new instance of StoreFixtureQuery.
func NewStoreFixtureQuery() *StoreFixtureQuery {
	return &StoreFixtureQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.StoreFixture.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *StoreFixtureQuery) Select(columns ...kallax.SchemaField) *StoreFixtureQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *StoreFixtureQuery) SelectNot(columns ...kallax.SchemaField) *StoreFixtureQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *StoreFixtureQuery) Copy() *StoreFixtureQuery {
	return &StoreFixtureQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *StoreFixtureQuery) Order(cols ...kallax.ColumnOrder) *StoreFixtureQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *StoreFixtureQuery) BatchSize(size uint64) *StoreFixtureQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *StoreFixtureQuery) Limit(n uint64) *StoreFixtureQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *StoreFixtureQuery) Offset(n uint64) *StoreFixtureQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *StoreFixtureQuery) Where(cond kallax.Condition) *StoreFixtureQuery {
	q.BaseQuery.Where(cond)
	return q
}

// StoreFixtureResultSet is the set of results returned by a query to the
// database.
type StoreFixtureResultSet struct {
	*kallax.ResultSet
	last    *StoreFixture
	lastErr error
}

func NewStoreFixtureResultSet(rs *kallax.ResultSet) *StoreFixtureResultSet {
	return &StoreFixtureResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *StoreFixtureResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.Close()
		return false
	}

	rs.last = new(StoreFixture)
	rs.lastErr = rs.Scan(rs.last)
	if rs.lastErr != nil {
		rs.last = nil
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *StoreFixtureResultSet) Get() (*StoreFixture, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *StoreFixtureResultSet) ForEach(fn func(*StoreFixture) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *StoreFixtureResultSet) All() ([]*StoreFixture, error) {
	var result []*StoreFixture
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *StoreFixtureResultSet) One() (*StoreFixture, error) {
	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// NewStoreWithConstructFixture returns a new instance of StoreWithConstructFixture.
func NewStoreWithConstructFixture(f string) (record *StoreWithConstructFixture) {
	record = newStoreWithConstructFixture(f)
	if record != nil {
		record.SetID(kallax.NewID())
	}
	return
}

func (r *StoreWithConstructFixture) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return &r.Model.ID, nil
	case "foo":
		return &r.Foo, nil

	default:
		return nil, fmt.Errorf("invalid column in StoreWithConstructFixture: %s", col)
	}
}

func (r *StoreWithConstructFixture) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.Model.ID, nil
	case "foo":
		return r.Foo, nil

	default:
		return nil, fmt.Errorf("invalid column in StoreWithConstructFixture: %s", col)
	}
}

func (r *StoreWithConstructFixture) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {

	}
	return nil, fmt.Errorf("model StoreWithConstructFixture has no relationship %s", field)
}

func (r *StoreWithConstructFixture) SetRelationship(field string, record kallax.Record) error {
	switch field {

	}
	return fmt.Errorf("model StoreWithConstructFixture has no relationship %s", field)
}

// StoreWithConstructFixtureStore is the entity to access the records of the type StoreWithConstructFixture
// in the database.
type StoreWithConstructFixtureStore struct {
	*kallax.Store
}

// NewStoreWithConstructFixtureStore creates a new instance of StoreWithConstructFixtureStore
// using a SQL database.
func NewStoreWithConstructFixtureStore(db *sql.DB) *StoreWithConstructFixtureStore {
	return &StoreWithConstructFixtureStore{kallax.NewStore(db, Schema.StoreWithConstructFixture.BaseSchema)}
}

// Insert inserts a StoreWithConstructFixture in the database. A non-persisted object is
// required for this operation.
func (s *StoreWithConstructFixtureStore) Insert(record *StoreWithConstructFixture) error {

	return s.Store.Insert(record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *StoreWithConstructFixtureStore) Update(record *StoreWithConstructFixture, cols ...kallax.SchemaField) (int64, error) {

	return s.Store.Update(record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *StoreWithConstructFixtureStore) Save(record *StoreWithConstructFixture) (updated bool, err error) {

	return s.Store.Save(record)
}

// Delete removes the given record from the database.
func (s *StoreWithConstructFixtureStore) Delete(record *StoreWithConstructFixture) error {
	return s.Store.Delete(record)
}

// Find returns the set of results for the given query.
func (s *StoreWithConstructFixtureStore) Find(q *StoreWithConstructFixtureQuery) (*StoreWithConstructFixtureResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewStoreWithConstructFixtureResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *StoreWithConstructFixtureStore) MustFind(q *StoreWithConstructFixtureQuery) *StoreWithConstructFixtureResultSet {
	return NewStoreWithConstructFixtureResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *StoreWithConstructFixtureStore) Count(q *StoreWithConstructFixtureQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *StoreWithConstructFixtureStore) MustCount(q *StoreWithConstructFixtureQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `sql.ErrNoRows` is returned if there are no results.
func (s *StoreWithConstructFixtureStore) FindOne(q *StoreWithConstructFixtureQuery) (*StoreWithConstructFixture, error) {
	q.Limit(1)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *StoreWithConstructFixtureStore) MustFindOne(q *StoreWithConstructFixtureQuery) *StoreWithConstructFixture {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the StoreWithConstructFixture with the data in the database and
// makes it writable.
func (s *StoreWithConstructFixtureStore) Reload(record *StoreWithConstructFixture) error {
	return s.Store.Reload(record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *StoreWithConstructFixtureStore) Transaction(callback func(*StoreWithConstructFixtureStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&StoreWithConstructFixtureStore{store})
	})
}

// StoreWithConstructFixtureQuery is the object used to create queries for the StoreWithConstructFixture
// entity.
type StoreWithConstructFixtureQuery struct {
	*kallax.BaseQuery
}

// NewStoreWithConstructFixtureQuery returns a new instance of StoreWithConstructFixtureQuery.
func NewStoreWithConstructFixtureQuery() *StoreWithConstructFixtureQuery {
	return &StoreWithConstructFixtureQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.StoreWithConstructFixture.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *StoreWithConstructFixtureQuery) Select(columns ...kallax.SchemaField) *StoreWithConstructFixtureQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *StoreWithConstructFixtureQuery) SelectNot(columns ...kallax.SchemaField) *StoreWithConstructFixtureQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *StoreWithConstructFixtureQuery) Copy() *StoreWithConstructFixtureQuery {
	return &StoreWithConstructFixtureQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *StoreWithConstructFixtureQuery) Order(cols ...kallax.ColumnOrder) *StoreWithConstructFixtureQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *StoreWithConstructFixtureQuery) BatchSize(size uint64) *StoreWithConstructFixtureQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *StoreWithConstructFixtureQuery) Limit(n uint64) *StoreWithConstructFixtureQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *StoreWithConstructFixtureQuery) Offset(n uint64) *StoreWithConstructFixtureQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *StoreWithConstructFixtureQuery) Where(cond kallax.Condition) *StoreWithConstructFixtureQuery {
	q.BaseQuery.Where(cond)
	return q
}

// StoreWithConstructFixtureResultSet is the set of results returned by a query to the
// database.
type StoreWithConstructFixtureResultSet struct {
	*kallax.ResultSet
	last    *StoreWithConstructFixture
	lastErr error
}

func NewStoreWithConstructFixtureResultSet(rs *kallax.ResultSet) *StoreWithConstructFixtureResultSet {
	return &StoreWithConstructFixtureResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *StoreWithConstructFixtureResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.Close()
		return false
	}

	rs.last = new(StoreWithConstructFixture)
	rs.lastErr = rs.Scan(rs.last)
	if rs.lastErr != nil {
		rs.last = nil
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *StoreWithConstructFixtureResultSet) Get() (*StoreWithConstructFixture, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *StoreWithConstructFixtureResultSet) ForEach(fn func(*StoreWithConstructFixture) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *StoreWithConstructFixtureResultSet) All() ([]*StoreWithConstructFixture, error) {
	var result []*StoreWithConstructFixture
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *StoreWithConstructFixtureResultSet) One() (*StoreWithConstructFixture, error) {
	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// NewStoreWithNewFixture returns a new instance of StoreWithNewFixture.
func NewStoreWithNewFixture() (record *StoreWithNewFixture) {
	record = &StoreWithNewFixture{}
	if record != nil {
		record.SetID(kallax.NewID())
	}
	return
}

func (r *StoreWithNewFixture) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return &r.Model.ID, nil
	case "foo":
		return &r.Foo, nil
	case "bar":
		return &r.Bar, nil

	default:
		return nil, fmt.Errorf("invalid column in StoreWithNewFixture: %s", col)
	}
}

func (r *StoreWithNewFixture) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.Model.ID, nil
	case "foo":
		return r.Foo, nil
	case "bar":
		return r.Bar, nil

	default:
		return nil, fmt.Errorf("invalid column in StoreWithNewFixture: %s", col)
	}
}

func (r *StoreWithNewFixture) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {

	}
	return nil, fmt.Errorf("model StoreWithNewFixture has no relationship %s", field)
}

func (r *StoreWithNewFixture) SetRelationship(field string, record kallax.Record) error {
	switch field {

	}
	return fmt.Errorf("model StoreWithNewFixture has no relationship %s", field)
}

// StoreWithNewFixtureStore is the entity to access the records of the type StoreWithNewFixture
// in the database.
type StoreWithNewFixtureStore struct {
	*kallax.Store
}

// NewStoreWithNewFixtureStore creates a new instance of StoreWithNewFixtureStore
// using a SQL database.
func NewStoreWithNewFixtureStore(db *sql.DB) *StoreWithNewFixtureStore {
	return &StoreWithNewFixtureStore{kallax.NewStore(db, Schema.StoreWithNewFixture.BaseSchema)}
}

// Insert inserts a StoreWithNewFixture in the database. A non-persisted object is
// required for this operation.
func (s *StoreWithNewFixtureStore) Insert(record *StoreWithNewFixture) error {

	return s.Store.Insert(record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *StoreWithNewFixtureStore) Update(record *StoreWithNewFixture, cols ...kallax.SchemaField) (int64, error) {

	return s.Store.Update(record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *StoreWithNewFixtureStore) Save(record *StoreWithNewFixture) (updated bool, err error) {

	return s.Store.Save(record)
}

// Delete removes the given record from the database.
func (s *StoreWithNewFixtureStore) Delete(record *StoreWithNewFixture) error {
	return s.Store.Delete(record)
}

// Find returns the set of results for the given query.
func (s *StoreWithNewFixtureStore) Find(q *StoreWithNewFixtureQuery) (*StoreWithNewFixtureResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewStoreWithNewFixtureResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *StoreWithNewFixtureStore) MustFind(q *StoreWithNewFixtureQuery) *StoreWithNewFixtureResultSet {
	return NewStoreWithNewFixtureResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *StoreWithNewFixtureStore) Count(q *StoreWithNewFixtureQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *StoreWithNewFixtureStore) MustCount(q *StoreWithNewFixtureQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `sql.ErrNoRows` is returned if there are no results.
func (s *StoreWithNewFixtureStore) FindOne(q *StoreWithNewFixtureQuery) (*StoreWithNewFixture, error) {
	q.Limit(1)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *StoreWithNewFixtureStore) MustFindOne(q *StoreWithNewFixtureQuery) *StoreWithNewFixture {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the StoreWithNewFixture with the data in the database and
// makes it writable.
func (s *StoreWithNewFixtureStore) Reload(record *StoreWithNewFixture) error {
	return s.Store.Reload(record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *StoreWithNewFixtureStore) Transaction(callback func(*StoreWithNewFixtureStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&StoreWithNewFixtureStore{store})
	})
}

// StoreWithNewFixtureQuery is the object used to create queries for the StoreWithNewFixture
// entity.
type StoreWithNewFixtureQuery struct {
	*kallax.BaseQuery
}

// NewStoreWithNewFixtureQuery returns a new instance of StoreWithNewFixtureQuery.
func NewStoreWithNewFixtureQuery() *StoreWithNewFixtureQuery {
	return &StoreWithNewFixtureQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.StoreWithNewFixture.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *StoreWithNewFixtureQuery) Select(columns ...kallax.SchemaField) *StoreWithNewFixtureQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *StoreWithNewFixtureQuery) SelectNot(columns ...kallax.SchemaField) *StoreWithNewFixtureQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *StoreWithNewFixtureQuery) Copy() *StoreWithNewFixtureQuery {
	return &StoreWithNewFixtureQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *StoreWithNewFixtureQuery) Order(cols ...kallax.ColumnOrder) *StoreWithNewFixtureQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *StoreWithNewFixtureQuery) BatchSize(size uint64) *StoreWithNewFixtureQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *StoreWithNewFixtureQuery) Limit(n uint64) *StoreWithNewFixtureQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *StoreWithNewFixtureQuery) Offset(n uint64) *StoreWithNewFixtureQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *StoreWithNewFixtureQuery) Where(cond kallax.Condition) *StoreWithNewFixtureQuery {
	q.BaseQuery.Where(cond)
	return q
}

// StoreWithNewFixtureResultSet is the set of results returned by a query to the
// database.
type StoreWithNewFixtureResultSet struct {
	*kallax.ResultSet
	last    *StoreWithNewFixture
	lastErr error
}

func NewStoreWithNewFixtureResultSet(rs *kallax.ResultSet) *StoreWithNewFixtureResultSet {
	return &StoreWithNewFixtureResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *StoreWithNewFixtureResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.Close()
		return false
	}

	rs.last = new(StoreWithNewFixture)
	rs.lastErr = rs.Scan(rs.last)
	if rs.lastErr != nil {
		rs.last = nil
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *StoreWithNewFixtureResultSet) Get() (*StoreWithNewFixture, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *StoreWithNewFixtureResultSet) ForEach(fn func(*StoreWithNewFixture) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *StoreWithNewFixtureResultSet) All() ([]*StoreWithNewFixture, error) {
	var result []*StoreWithNewFixture
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *StoreWithNewFixtureResultSet) One() (*StoreWithNewFixture, error) {
	if !rs.Next() {
		return nil, sql.ErrNoRows
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

type schema struct {
	EventsFixture             *schemaEventsFixture
	EventsSaveFixture         *schemaEventsSaveFixture
	MultiKeySortFixture       *schemaMultiKeySortFixture
	QueryFixture              *schemaQueryFixture
	ResultSetFixture          *schemaResultSetFixture
	ResultSetInitFixture      *schemaResultSetInitFixture
	SchemaFixture             *schemaSchemaFixture
	StoreFixture              *schemaStoreFixture
	StoreWithConstructFixture *schemaStoreWithConstructFixture
	StoreWithNewFixture       *schemaStoreWithNewFixture
}

type schemaEventsFixture struct {
	*kallax.BaseSchema
	ID             kallax.SchemaField
	Checks         kallax.SchemaField
	MustFailBefore kallax.SchemaField
	MustFailAfter  kallax.SchemaField
}

type schemaEventsSaveFixture struct {
	*kallax.BaseSchema
	ID             kallax.SchemaField
	Checks         kallax.SchemaField
	MustFailBefore kallax.SchemaField
	MustFailAfter  kallax.SchemaField
}

type schemaMultiKeySortFixture struct {
	*kallax.BaseSchema
	ID    kallax.SchemaField
	Name  kallax.SchemaField
	Start kallax.SchemaField
	End   kallax.SchemaField
}

type schemaQueryFixture struct {
	*kallax.BaseSchema
	ID  kallax.SchemaField
	Foo kallax.SchemaField
}

type schemaResultSetFixture struct {
	*kallax.BaseSchema
	ID  kallax.SchemaField
	Foo kallax.SchemaField
}

type schemaResultSetInitFixture struct {
	*kallax.BaseSchema
	ID  kallax.SchemaField
	Foo kallax.SchemaField
}

type schemaSchemaFixture struct {
	*kallax.BaseSchema
	ID     kallax.SchemaField
	String kallax.SchemaField
	Int    kallax.SchemaField
}

type schemaStoreFixture struct {
	*kallax.BaseSchema
	ID  kallax.SchemaField
	Foo kallax.SchemaField
}

type schemaStoreWithConstructFixture struct {
	*kallax.BaseSchema
	ID  kallax.SchemaField
	Foo kallax.SchemaField
}

type schemaStoreWithNewFixture struct {
	*kallax.BaseSchema
	ID  kallax.SchemaField
	Foo kallax.SchemaField
	Bar kallax.SchemaField
}

var Schema = &schema{
	EventsFixture: &schemaEventsFixture{
		BaseSchema: kallax.NewBaseSchema(
			"event",
			"__eventsfixture",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("checks"),
			kallax.NewSchemaField("must_fail_before"),
			kallax.NewSchemaField("must_fail_after"),
		),
		ID:             kallax.NewSchemaField("id"),
		Checks:         kallax.NewSchemaField("checks"),
		MustFailBefore: kallax.NewSchemaField("must_fail_before"),
		MustFailAfter:  kallax.NewSchemaField("must_fail_after"),
	},
	EventsSaveFixture: &schemaEventsSaveFixture{
		BaseSchema: kallax.NewBaseSchema(
			"event",
			"__eventssavefixture",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("checks"),
			kallax.NewSchemaField("must_fail_before"),
			kallax.NewSchemaField("must_fail_after"),
		),
		ID:             kallax.NewSchemaField("id"),
		Checks:         kallax.NewSchemaField("checks"),
		MustFailBefore: kallax.NewSchemaField("must_fail_before"),
		MustFailAfter:  kallax.NewSchemaField("must_fail_after"),
	},
	MultiKeySortFixture: &schemaMultiKeySortFixture{
		BaseSchema: kallax.NewBaseSchema(
			"query",
			"__multikeysortfixture",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("name"),
			kallax.NewSchemaField("start"),
			kallax.NewSchemaField("_end"),
		),
		ID:    kallax.NewSchemaField("id"),
		Name:  kallax.NewSchemaField("name"),
		Start: kallax.NewSchemaField("start"),
		End:   kallax.NewSchemaField("_end"),
	},
	QueryFixture: &schemaQueryFixture{
		BaseSchema: kallax.NewBaseSchema(
			"query",
			"__queryfixture",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("foo"),
		),
		ID:  kallax.NewSchemaField("id"),
		Foo: kallax.NewSchemaField("foo"),
	},
	ResultSetFixture: &schemaResultSetFixture{
		BaseSchema: kallax.NewBaseSchema(
			"resultset",
			"__resultsetfixture",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("foo"),
		),
		ID:  kallax.NewSchemaField("id"),
		Foo: kallax.NewSchemaField("foo"),
	},
	ResultSetInitFixture: &schemaResultSetInitFixture{
		BaseSchema: kallax.NewBaseSchema(
			"resultset",
			"__resultsetinitfixture",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("foo"),
		),
		ID:  kallax.NewSchemaField("id"),
		Foo: kallax.NewSchemaField("foo"),
	},
	SchemaFixture: &schemaSchemaFixture{
		BaseSchema: kallax.NewBaseSchema(
			"schema",
			"__schemafixture",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"Nested": kallax.NewSchemaField("schema_fixture_id"),
			},
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("string"),
			kallax.NewSchemaField("int"),
			kallax.NewSchemaField("inline"),
			kallax.NewSchemaField("map_of_string"),
			kallax.NewSchemaField("map_of_interface"),
			kallax.NewSchemaField("map_of_some_type"),
		),
		ID:     kallax.NewSchemaField("id"),
		String: kallax.NewSchemaField("string"),
		Int:    kallax.NewSchemaField("int"),
	},
	StoreFixture: &schemaStoreFixture{
		BaseSchema: kallax.NewBaseSchema(
			"store",
			"__storefixture",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("foo"),
		),
		ID:  kallax.NewSchemaField("id"),
		Foo: kallax.NewSchemaField("foo"),
	},
	StoreWithConstructFixture: &schemaStoreWithConstructFixture{
		BaseSchema: kallax.NewBaseSchema(
			"store_construct",
			"__storewithconstructfixture",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("foo"),
		),
		ID:  kallax.NewSchemaField("id"),
		Foo: kallax.NewSchemaField("foo"),
	},
	StoreWithNewFixture: &schemaStoreWithNewFixture{
		BaseSchema: kallax.NewBaseSchema(
			"store_new",
			"__storewithnewfixture",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("foo"),
			kallax.NewSchemaField("bar"),
		),
		ID:  kallax.NewSchemaField("id"),
		Foo: kallax.NewSchemaField("foo"),
		Bar: kallax.NewSchemaField("bar"),
	},
}
