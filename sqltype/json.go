// Package sqltype provides database-aware types for use with database/sql.
package sqltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

var _ sql.Scanner = &JSON[any]{}

// JSON is a generic type that holds a value of type T and serializes it as
// JSON when reading from or writing to a database. It implements sql.Scanner
// and driver.Valuer, so it can be used as a column type with database/sql.
//
// JSON caches the last marshaled bytes so repeated Value() calls without
// changing data avoid re-marshaling. Use Get and Set to read or update the
// wrapped value.
//
// Example:
//
//	type MyRow struct {
//	    ID   int         `db:"id"`
//	    Meta sqltype.JSON[map[string]any] `db:"meta"`
//	}
//	// ... scan row into MyRow, then meta := row.Meta.Get()
type JSON[T any] struct {
	Data T

	rawJson []byte
}

// Get returns the wrapped value. If the receiver is nil, it returns the zero
// value of T.
func (j *JSON[T]) Get() T {
	if j == nil {
		var zero T
		return zero
	}

	return j.Data
}

// Set stores the given value and clears the cached JSON so the next Value()
// call will marshal the new data.
func (j *JSON[T]) Set(data T) {
	if j == nil {
		return
	}

	j.Data = data
	j.rawJson = nil
}

// Scan implements sql.Scanner. It accepts []byte or string from the driver,
// unmarshals JSON into the wrapped type T, and caches the raw bytes for
// Value().
func (j *JSON[T]) Scan(src any) error {
	if src == nil {
		return nil
	}

	*j = JSON[T]{}

	switch v := src.(type) {
	case []byte:
		j.rawJson = v
		return json.Unmarshal(v, &j.Data)
	case string:
		j.rawJson = []byte(v)
		return json.Unmarshal([]byte(v), &j.Data)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}

// Value implements driver.Valuer. It returns the cached raw JSON when
// available (e.g. after Scan), otherwise marshals the wrapped data, caches
// it, and returns the bytes. Returns (nil, nil) when the receiver is nil.
func (j *JSON[T]) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	if len(j.rawJson) > 0 {
		return j.rawJson, nil
	}

	rawJson, err := json.Marshal(j.Data)
	if err != nil {
		return nil, err
	}

	j.rawJson = rawJson

	return rawJson, nil
}
