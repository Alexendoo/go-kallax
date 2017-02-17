package tests

import (
	"time"

	"github.com/src-d/go-kallax"
)

type StoreFixture struct {
	kallax.Model `table:"store"`
	ID           kallax.ULID `pk:""`
	Foo          string
}

func newStoreFixture() *StoreFixture {
	return &StoreFixture{ID: kallax.NewULID()}
}

type StoreWithConstructFixture struct {
	kallax.Model `table:"store_construct"`
	ID           kallax.ULID `pk:""`
	Foo          string
}

func newStoreWithConstructFixture(f string) *StoreWithConstructFixture {
	if f == "" {
		return nil
	}
	return &StoreWithConstructFixture{ID: kallax.NewULID(), Foo: f}
}

type StoreWithNewFixture struct {
	kallax.Model `table:"store_new"`
	ID           kallax.ULID `pk:""`
	Foo          string
	Bar          string
}

func newStoreWithNewFixture() *StoreWithNewFixture {
	return &StoreWithNewFixture{ID: kallax.NewULID()}
}

type MultiKeySortFixture struct {
	kallax.Model `table:"query"`
	ID           kallax.ULID `pk:""`
	Name         string
	Start        time.Time
	End          time.Time
}

func newMultiKeySortFixture() *MultiKeySortFixture {
	return &MultiKeySortFixture{ID: kallax.NewULID()}
}

type Nullable struct {
	kallax.Model `table:"nullable"`
	ID           int64 `pk:"autoincr"`
	T            *time.Time
}
