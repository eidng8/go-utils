package utils

import (
	"github.com/google/uuid"
)

var (
	uuidV7 = uuid.NewV7
	uuidV6 = uuid.NewV6
	uuidV4 = uuid.NewRandom
)

// PanicIfError panics if the error is not nil.
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

// Ptr returns a pointer to the value. Use with caution! This function directly
// returns a pointer to the parameter.
func Ptr[T any](v T) *T {
	return &v
}

// ReturnOrPanic returns the value if the error is nil, otherwise panics.
func ReturnOrPanic[T any](v T, err error) T {
	PanicIfError(err)
	return v
}

// NewUuid returns a new UUID. It tries to generate a UUID using V7, V6, and
// finally falls back to V4. If all methods fail, it returns an error.
func NewUuid() (uuid.UUID, error) {
	id, err := uuidV7()
	if nil == err {
		return id, nil
	}
	id, err = uuidV6()
	if nil == err {
		return id, nil
	}
	id, err = uuidV4()
	if nil == err {
		return id, nil
	}
	return uuid.Nil, err
}
