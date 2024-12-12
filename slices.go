package utils

import (
	"errors"
	"fmt"
	"strings"
)

// Integer is a type constraint that includes all integer types
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Pluck extracts a list of values from a slice of structs.
// The predicate function fn is applied to each element of the slice, and the
// result is stored in a new slice.
func Pluck[I interface{ ~[]V }, V any, T any](
	a I, fn func(V) T,
) []T {
	r := make([]T, len(a))
	for i, v := range a {
		r[i] = fn(v)
	}
	return r
}

// JoinInteger converts a slice of integers to a string, joining them with the
// separator.
func JoinInteger[A interface{ ~[]T }, T Integer](a A, sep string) string {
	e := make([]string, len(a))
	for i, v := range a {
		e[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(e, sep)
}

// SliceMapFunc applies the predicate function to each element of the slice,
// returning a new slice with the results. If the predicate function returns an
// error, the function stops and returns the error.
func SliceMapFunc[I, O any](a []I, fn func(I) (O, error)) ([]O, error) {
	var err error
	r := make([]O, len(a))
	for i, v := range a {
		r[i], err = fn(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// SliceFindFunc returns the first element in the slice that satisfies the
// predicate function, or the zero value of the element type if none is found.
func SliceFindFunc[I any](a []I, fn func(I) bool) I {
	var zero I
	for _, v := range a {
		if fn(v) {
			return v
		}
	}
	return zero
}

// MapToType converts an element of interface{} to an element of specific type,
// returning an error if the conversion is not possible.
func MapToType[T any](val any) (T, error) {
	var zero T
	v, ok := val.(T)
	if !ok {
		return zero, errors.New("inconvertible value")
	}
	return v, nil
}
