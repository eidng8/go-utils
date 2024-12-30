package utils

// PanicIfError panics if the error is not nil.
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
