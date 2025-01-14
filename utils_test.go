package utils

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PanicIfError(t *testing.T) {
	require.Panics(t, func() { PanicIfError(assert.AnError) })
	require.NotPanics(t, func() { PanicIfError(nil) })
}

func Test_Ptr(t *testing.T) {
	num := 1
	expected := &num
	require.Equal(t, expected, Ptr(1))
}

func Test_ReturnOrPanic(t *testing.T) {
	require.Panics(
		t, func() { ReturnOrPanic[struct{}](struct{}{}, assert.AnError) },
	)
	require.NotPanics(t, func() { ReturnOrPanic(struct{}{}, nil) })
}

func Test_NewUuid_returns_V7(t *testing.T) {
	id, err := NewUuid()
	require.Nil(t, err)
	require.Equal(t, uuid.Version(7), id.Version())
}

func Test_NewUuid_returns_V6(t *testing.T) {
	uuidV7 = func() (uuid.UUID, error) { return uuid.Nil, assert.AnError }
	id, err := NewUuid()
	require.Nil(t, err)
	require.Equal(t, uuid.Version(6), id.Version())
}

func Test_NewUuid_returns_V4(t *testing.T) {
	uuidV7 = func() (uuid.UUID, error) { return uuid.Nil, assert.AnError }
	uuidV6 = func() (uuid.UUID, error) { return uuid.Nil, assert.AnError }
	id, err := NewUuid()
	require.Nil(t, err)
	require.Equal(t, uuid.Version(4), id.Version())
}

func Test_NewUuid_returns_nil(t *testing.T) {
	uuidV7 = func() (uuid.UUID, error) { return uuid.Nil, assert.AnError }
	uuidV6 = func() (uuid.UUID, error) { return uuid.Nil, assert.AnError }
	uuidV4 = func() (uuid.UUID, error) { return uuid.Nil, assert.AnError }
	id, err := NewUuid()
	require.NotNil(t, err)
	require.Equal(t, uuid.Nil, id)
}
