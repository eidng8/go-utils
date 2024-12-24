package utils

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/json-iterator/go"
)

var Jsoniter = jsoniter.ConfigCompatibleWithStandardLibrary

// Integer is a type constraint that includes all integer types
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Floats interface {
	~float32 | ~float64
}

type ComplexNumbers interface {
	~complex64 | ~complex128
}

type Numbers interface {
	Integer | Floats | ComplexNumbers
}

func CloneDeepJsonable[P *V, V any](src P) (P, error) {
	if nil == src {
		return nil, nil
	}
	js, err := Jsoniter.Marshal(src)
	if nil != err {
		return nil, err
	}
	var dst V
	err = Jsoniter.Unmarshal(js, &dst)
	if nil != err {
		return nil, err
	}
	return &dst, nil
}

func ContainsAny[A ~[]T, T comparable](a A, b A) bool {
	for _, v := range a {
		if slices.Contains(b, v) {
			return true
		}
	}
	return false
}

func FilterFunc[A ~[]T, T any](a A, fn func(T, int, A) bool) A {
	var res []T
	for i, v := range a {
		if fn(v, i, a) {
			res = append(res, v)
		}
	}
	return res
}

func Intersect[A ~[]T, T comparable](a A, b A) A {
	var res []T
	for _, v := range a {
		if slices.Contains(b, v) {
			res = append(res, v)
		}
	}
	return res
}

func IntersectFunc[A ~[]T, T any, O comparable](
	a A, b A, fn func(T, int, A) O,
) A {
	var res []T
	for i, v := range a {
		va := fn(v, i, a)
		for j, x := range b {
			if va == fn(x, j, b) {
				res = append(res, v)
				break
			}
		}
	}
	return res
}

// JoinInteger converts a slice of integers to a string, joining them with the
// separator.
func JoinInteger[A ~[]T, T Integer](a A, sep string) string {
	return JoinNumbersWithFormat(a, sep, "%d")
}

func JoinNumbersWithFormat[A ~[]T, T Numbers](a A, sep, format string) string {
	e := make([]string, len(a))
	for i, v := range a {
		e[i] = fmt.Sprintf(format, v)
	}
	return strings.Join(e, sep)
}

// MapToType converts an element of interface{} to an element of specific type,
// returning an error if the conversion is not possible. Mainly for use as
// predicate function in higher-order functions such as SliceMapFunc.
func MapToType[T any](val any) (T, error) {
	var zero T
	v, ok := val.(T)
	if !ok {
		return zero, errors.New("inconvertible value")
	}
	return v, nil
}

// Pluck extracts a list of values from a slice of structs.
// The predicate function fn is applied to each element of the slice, and the
// result is stored in a new slice.
func Pluck[I ~[]V, V any, T any](
	a I, fn func(V, int, I) T,
) []T {
	r := make([]T, len(a))
	for i, v := range a {
		r[i] = fn(v, i, a)
	}
	return r
}

// SliceMapFunc applies the predicate function to each element of the slice,
// returning a new slice with the results. If the predicate function returns an
// error, the function stops and returns the error.
func SliceMapFunc[AI ~[]I, AO ~[]O, I, O any](
	a AI, fn func(I, int, AI) (O, error),
) (AO, error) {
	var err error
	r := make([]O, len(a))
	for i, v := range a {
		r[i], err = fn(v, i, a)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// SliceFindFunc returns the first element in the slice that satisfies the
// predicate function, or the zero value of the element type if none is found.
func SliceFindFunc[A ~[]I, I any](a A, fn func(I, int, A) bool) I {
	var zero I
	for i, v := range a {
		if fn(v, i, a) {
			return v
		}
	}
	return zero
}

// ApplyFunc applies the given function to each element of the slice.
// This function is intended to operate directly on elements in the slice.
func ApplyFunc[A ~[]T, T any](a A, fn func(T, int, A)) A {
	for i, v := range a {
		fn(v, i, a)
	}
	return a
}
