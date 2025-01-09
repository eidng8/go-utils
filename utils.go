package utils

// PanicIfError panics if the error is not nil.
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

// Ptr returns a pointer to the value.
func Ptr[T any](v T) *T {
	return &v
}

// ReturnOrPanic returns the value if the error is nil, otherwise panics.
func ReturnOrPanic[T any](v T, err error) T {
	PanicIfError(err)
	return v
}
