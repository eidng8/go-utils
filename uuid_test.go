package utils

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var id UUID = &Uuid{}

func Test_NewUuid_returns_V7(t *testing.T) {
	require.Nil(t, id.New())
	require.Equal(t, uuid.Version(7), id.Version())
}

func Test_NewUuid_returns_V6(t *testing.T) {
	defer func() { uuidV7 = uuid.NewV7 }()
	uuidV7 = func() (uuid.UUID, error) { return uuid.Nil, assert.AnError }
	require.Nil(t, id.New())
	require.Equal(t, uuid.Version(6), id.Version())
}

func Test_NewUuid_returns_V4(t *testing.T) {
	defer func() {
		uuidV7 = uuid.NewV7
		uuidV6 = uuid.NewV6
	}()
	uuidV7 = func() (uuid.UUID, error) { return uuid.Nil, assert.AnError }
	uuidV6 = func() (uuid.UUID, error) { return uuid.Nil, assert.AnError }
	require.Nil(t, id.New())
	require.Equal(t, uuid.Version(4), id.Version())
}

func Test_NewUuid_returns_nil(t *testing.T) {
	defer func() {
		uuidV7 = uuid.NewV7
		uuidV6 = uuid.NewV6
		uuidV4 = uuid.NewRandom
	}()
	uuidV7 = func() (uuid.UUID, error) { return uuid.Nil, assert.AnError }
	uuidV6 = func() (uuid.UUID, error) { return uuid.Nil, assert.AnError }
	uuidV4 = func() (uuid.UUID, error) { return uuid.Nil, assert.AnError }
	require.NotNil(t, id.New())
	require.Equal(t, uuid.Nil, id.Get())
}

func Test_UUID_MarshalBinary(t *testing.T) {
	require.Nil(t, id.New())
	b, err := id.MarshalBinary()
	require.Nil(t, err)
	require.Equal(t, 16, len(b))
}

func Test_UUID_MarshalText(t *testing.T) {
	require.Nil(t, id.New())
	b, err := id.MarshalText()
	require.Nil(t, err)
	require.Equal(t, 36, len(b))
}

func Test_UUID_UnmarshalBinary(t *testing.T) {
	require.Nil(t, id.New())
	b, err := id.MarshalBinary()
	require.Nil(t, err)
	require.Nil(t, id.UnmarshalBinary(b))
}

func Test_UUID_UnmarshalText(t *testing.T) {
	require.Nil(t, id.New())
	b, err := id.MarshalText()
	require.Nil(t, err)
	require.Nil(t, id.UnmarshalText(b))
}
