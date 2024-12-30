package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PanicIfError(t *testing.T) {
	assert.Panics(t, func() { PanicIfError(assert.AnError) })
	assert.NotPanics(t, func() { PanicIfError(nil) })
}
