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

// ApplyFunc applies the given function to each element of the slice.
// This function is intended to operate directly on elements in the slice.
func ApplyFunc[A ~[]T, T any](a A, fn func(T)) A {
	for _, v := range a {
		fn(v)
	}
	return a
}

// ApplyFuncA applies the given function to each element of the slice.
// This function is intended to operate directly on elements in the slice.
// The predicate is called with current index and the original slice.
func ApplyFuncA[A ~[]T, T any](a A, fn func(T, int, A)) A {
	for i, v := range a {
		fn(v, i, a)
	}
	return a
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

// FilterFunc returns a new slice containing only the elements of the original
// slice that satisfy the predicate function.
func FilterFunc[A ~[]T, T any](a A, fn func(T) bool) A {
	var res []T
	for _, v := range a {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}

// FilterFuncA returns a new slice containing only the elements of the original
// slice that satisfy the predicate function. The predicate is called with the
// current element, the index of the current element, and the original slice.
func FilterFuncA[A ~[]T, T any](a A, fn func(T, int, A) bool) A {
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

// IntersectFunc returns a new slice containing the intersection of two slices
// based on the result of the predicate function.
func IntersectFunc[A ~[]T, T any, O comparable](a A, b A, fn func(T) O) A {
	var res []T
	for _, v := range a {
		va := fn(v)
		for _, x := range b {
			if va == fn(x) {
				res = append(res, v)
				break
			}
		}
	}
	return res
}

// IntersectFuncA returns a new slice containing the intersection of two slices
// based on the result of the predicate function. The predicate function is
// called with the current element, the index of the current element, and the
// original slice.
func IntersectFuncA[A ~[]T, T any, O comparable](
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

// MapToAny converts an element of specific type to an element of type any,
// returning an error if the conversion is not possible. Mainly for use as
// predicate function in higher-order functions such as SliceMapFuncAE.
func MapToAny[T any](val T) (any, error) {
	return val, nil
}

// MapToType converts an element of interface{} to an element of specific type,
// returning an error if the conversion is not possible. Mainly for use as
// predicate function in higher-order functions such as SliceMapFunc.
func MapToType[T any](val any) T {
	return val.(T)
}

// MapToTypeE converts an element of interface{} to an element of specific type,
// returning an error if the conversion is not possible. Mainly for use as
// predicate function in higher-order functions such as SliceMapFuncAE.
func MapToTypeE[T any](val any) (T, error) {
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
func Pluck[A ~[]V, V any, T any](a A, fn func(V) T) []T {
	r := make([]T, len(a))
	for i, v := range a {
		r[i] = fn(v)
	}
	return r
}

// PluckA extracts a list of values from a slice of structs.
// The predicate function fn is applied to each element of the slice, and the
// result is stored in a new slice.
// The predicate is called with current index and the original slice.
func PluckA[A ~[]V, V any, T any](a A, fn func(V, int, A) T) []T {
	r := make([]T, len(a))
	for i, v := range a {
		r[i] = fn(v, i, a)
	}
	return r
}

// SliceMapFunc applies the predicate function to each element of the slice,
// returning a new slice with the results.
func SliceMapFunc[AO ~[]O, AI ~[]I, O, I any](a AI, fn func(I) O) AO {
	r := make([]O, len(a))
	for i, v := range a {
		r[i] = fn(v)
	}
	return r
}

// SliceMapFuncE applies the predicate function to each element of the slice,
// returning a new slice with the results. If the predicate function returns an
// error, the function stops and returns the error.
func SliceMapFuncE[AO ~[]O, AI ~[]I, O, I any](
	a AI, fn func(I) (O, error),
) (AO, error) {
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

// SliceMapFuncA applies the predicate function to each element of the slice,
// returning a new slice with the results. The predicate is called with
// current index and the original slice.
func SliceMapFuncA[AO ~[]O, AI ~[]I, O, I any](a AI, fn func(I, int, AI) O) AO {
	r := make([]O, len(a))
	for i, v := range a {
		r[i] = fn(v, i, a)
	}
	return r
}

// SliceMapFuncAE applies the predicate function to each element of the slice,
// returning a new slice with the results. If the predicate function returns an
// error, the function stops and returns the error. The predicate is called with
// current index and the original slice.
func SliceMapFuncAE[AO ~[]O, AI ~[]I, O, I any](
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
func SliceFindFunc[A ~[]I, I any](a A, fn func(I) bool) I {
	var zero I
	for _, v := range a {
		if fn(v) {
			return v
		}
	}
	return zero
}

// SliceFindFuncA returns the first element in the slice that satisfies the
// predicate function, or the zero value of the element type if none is found.
// The predicate is called with current index and the original slice.
func SliceFindFuncA[A ~[]I, I any](a A, fn func(I, int, A) bool) I {
	var zero I
	for i, v := range a {
		if fn(v, i, a) {
			return v
		}
	}
	return zero
}

func Union[A ~[]T, T comparable](a A, b A) A {
	c := make(A, len(a))
	copy(c, a)
	for _, v := range b {
		if slices.Contains(c, v) {
			continue
		}
		c = append(c, v)
	}
	return c
}

func UnionFunc[A ~[]I, I any](a A, b A, fn func(I) bool) A {
	c := make(A, len(a))
	copy(c, a)
	for _, v := range b {
		if fn(v) {
			c = append(c, v)
		}
	}
	return c
}

func UnionFuncA[A ~[]I, I any](a A, b A, fn func(I, int, A) bool) A {
	c := make(A, len(a))
	copy(c, a)
	for i, v := range b {
		if fn(v, i, c) {
			c = append(c, v)
		}
	}
	return c
}
