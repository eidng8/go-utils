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

func Pluck[I interface{ ~[]V }, V any, T any](
	a I, fn func(V) T,
) []T {
	r := make([]T, len(a))
	for i, v := range a {
		r[i] = fn(v)
	}
	return r
}

func JoinInteger[A interface{ ~[]T }, T Integer](a A, sep string) string {
	e := make([]string, len(a))
	for i, v := range a {
		e[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(e, sep)
}

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

func MapToType[T any](val any) (T, error) {
	var zero T
	v, ok := val.(T)
	if !ok {
		return zero, errors.New("inconvertible value")
	}
	return v, nil
}
