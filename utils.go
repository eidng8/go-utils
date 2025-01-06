package utils

// PanicIfError panics if the error is not nil.
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ReturnOrPanic[T any](v T, err error) T {
	PanicIfError(err)
	return v
}
