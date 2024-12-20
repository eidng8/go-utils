package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupEnvTest(tb testing.TB) {
	err := os.Setenv("TEST_ENV", "TEST")
	assert.Nil(tb, err)
	err = os.Setenv("TEST_ENV_EMPTY", "")
	assert.Nil(tb, err)
}

func TestGetEnvWithDefault(t *testing.T) {
	setupEnvTest(t)
	got := GetEnvWithDefault("TEST_ENV", "defaultValue")
	assert.Equal(t, "TEST", got)
	got = GetEnvWithDefault("TEST_ENV_EMPTY", "defaultValue")
	assert.Empty(t, got)
	got = GetEnvWithDefault("NO_DEF", "defaultValue")
	assert.Equal(t, "defaultValue", got)
}

func TestGetEnvWithDefaultNE(t *testing.T) {
	setupEnvTest(t)
	got := GetEnvWithDefaultNE("TEST_ENV", "defaultValue")
	assert.Equal(t, "TEST", got)
	got = GetEnvWithDefaultNE("TEST_ENV_EMPTY", "defaultValue")
	assert.Equal(t, "defaultValue", got)
	got = GetEnvWithDefaultNE("NO_DEF", "defaultValue")
	assert.Equal(t, "defaultValue", got)
}

func TestPanicIfError(t *testing.T) {
	assert.Panics(t, func() { PanicIfError(assert.AnError) })
	assert.NotPanics(t, func() { PanicIfError(nil) })
}
