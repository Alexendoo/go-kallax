package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"time"

	"github.com/lib/pq"
)

// Nullable gives the ability to scan nil values to the given type
// only if they implement sql.Scanner.
func Nullable(typ interface{}) interface{} {
	// TODO(erizocosmico): implement the rest of go basic types
	switch typ := typ.(type) {
	case *string:
		return &nullString{typ}
	case *bool:
		return &nullBool{typ}
	case *int64:
		return &nullInt64{typ}
	case *float64:
		return &nullFloat64{typ}
	case *time.Time:
		return &nullTime{typ}
	case *time.Duration:
		return &nullDuration{typ}
	case sql.Scanner:
		return &nullable{typ}
	}

	return &nullableErr{typ}
}

type nullableErr struct {
	v interface{}
}

func (n *nullableErr) Scan(_ interface{}) error {
	return fmt.Errorf("type %T is not nullable", n.v)
}

type nullable struct {
	typ sql.Scanner
}

func (n *nullable) Scan(v interface{}) error {
	if v == nil {
		return nil
	}
	return n.typ.Scan(v)
}

type nullString struct {
	v *string
}

func (n *nullString) Scan(v interface{}) error {
	ns := new(sql.NullString)
	if err := ns.Scan(v); err != nil {
		return err
	}
	*n.v = ns.String
	return nil
}

type nullBool struct {
	v *bool
}

func (n *nullBool) Scan(v interface{}) error {
	ns := new(sql.NullBool)
	if err := ns.Scan(v); err != nil {
		return err
	}
	*n.v = ns.Bool
	return nil
}

type nullInt64 struct {
	v *int64
}

func (n *nullInt64) Scan(v interface{}) error {
	ns := new(sql.NullInt64)
	if err := ns.Scan(v); err != nil {
		return err
	}
	*n.v = ns.Int64
	return nil
}

type nullFloat64 struct {
	v *float64
}

func (n *nullFloat64) Scan(v interface{}) error {
	ns := new(sql.NullFloat64)
	if err := ns.Scan(v); err != nil {
		return err
	}
	*n.v = ns.Float64
	return nil
}

type nullTime struct {
	v *time.Time
}

func (n *nullTime) Scan(v interface{}) error {
	ns := new(pq.NullTime)
	if err := ns.Scan(v); err != nil {
		return err
	}
	*n.v = ns.Time
	return nil
}

type nullDuration struct {
	v *time.Duration
}

func (n *nullDuration) Scan(v interface{}) error {
	ns := new(sql.NullInt64)
	if err := ns.Scan(v); err != nil {
		return err
	}
	*n.v = time.Duration(ns.Int64)
	return nil
}

// URL is a wrapper of url.URL that implements SQLType interface.
type URL url.URL

func (u *URL) Scan(v interface{}) error {
	switch t := v.(type) {
	case []byte:
		return u.Scan(string(t))
	case string:
		url, err := url.Parse(t)
		if err != nil {
			return fmt.Errorf("kallax: error scanning url: %s", err)
		}

		*u = URL(*url)
		return nil
	}
	return fmt.Errorf("kallax: cannot scan type %s into URL type", reflect.TypeOf(v))
}

func (u URL) Value() (interface{}, error) {
	url := url.URL(u)
	return (&url).String(), nil
}

// ScanJSON scans json v into dst.
// WARNING: This is here temporarily, might be removed in the future, use
// `JSON` instead.
func ScanJSON(v interface{}, dst interface{}) error {
	switch v := v.(type) {
	case []byte:
		return json.Unmarshal(v, dst)
	case string:
		return ScanJSON([]byte(v), dst)
	}

	return fmt.Errorf("kallax: cannot scan type %s into JSON type", reflect.TypeOf(v))
}

// JSONValue converts something into json.
// WARNING: This is here temporarily, might be removed in the future, use
// `JSON` instead.
func JSONValue(v interface{}) (driver.Value, error) {
	return json.Marshal(v)
}

// SQLType is the common interface a type has to fulfill to be considered a
// SQL type.
type SQLType interface {
	sql.Scanner
	driver.Valuer
}

type array struct {
	val  reflect.Value
	size int
}

func Array(v interface{}, size int) SQLType {
	return &array{reflect.ValueOf(v), size}
}

func (a *array) Scan(v interface{}) error {
	sliceTyp := reflect.SliceOf(a.val.Type().Elem().Elem())
	newSlice := reflect.MakeSlice(sliceTyp, 0, 0)
	slicePtr := reflect.New(sliceTyp)
	slicePtr.Elem().Set(newSlice)
	if err := pq.Array(slicePtr.Interface()).Scan(v); err != nil {
		return err
	}

	if slicePtr.Elem().Len() != a.size {
		return fmt.Errorf(
			"kallax: cannot scan array of size %d into array of size %d",
			newSlice.Len(),
			a.size,
		)
	}

	for i := 0; i < a.size; i++ {
		a.val.Elem().Index(i).Set(slicePtr.Elem().Index(i))
	}

	return nil
}

func (a *array) Value() (driver.Value, error) {
	sliceTyp := reflect.SliceOf(a.val.Type().Elem().Elem())
	newSlice := reflect.MakeSlice(sliceTyp, a.size, a.size)
	for i := 0; i < a.size; i++ {
		newSlice.Index(i).Set(a.val.Elem().Index(i))
	}

	slicePtr := reflect.New(sliceTyp)
	slicePtr.Elem().Set(newSlice)
	return pq.Array(slicePtr.Interface()).Value()
}

type sqlJSON struct {
	val interface{}
}

// JSON makes sure the given value is converted to and scanned from SQL as
// a JSON. Note that this uses the standard json.Unmarshal and json.Marshal
// and it relies on reflection. To speed up the encoding/decoding you can
// implement interfaces json.Marshaller and json.Unmarshaller for your type
// with, for example, ffjson.
func JSON(v interface{}) SQLType {
	return &sqlJSON{v}
}

func (j *sqlJSON) Scan(v interface{}) error {
	return ScanJSON(v, j.val)
}

func (j *sqlJSON) Value() (driver.Value, error) {
	return JSONValue(j.val)
}
