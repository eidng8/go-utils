package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PanicIfError(t *testing.T) {
	assert.Panics(t, func() { PanicIfError(assert.AnError) })
	assert.NotPanics(t, func() { PanicIfError(nil) })
}

func Test_ReturnOrPanic(t *testing.T) {
	require.Panics(
		t, func() { ReturnOrPanic[struct{}](struct{}{}, assert.AnError) },
	)
	require.NotPanics(t, func() { ReturnOrPanic(struct{}{}, nil) })
}
