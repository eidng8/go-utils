package testing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MockUuid_New_returns_error(t *testing.T) {
	mock := NewUuidMock(MockConfig{NewReturnsError: true})
	require.NotNil(t, mock.New())
}

func Test_MockUuid_New_does_not_return_error(t *testing.T) {
	mock := NewUuidMock(MockConfig{})
	require.NoError(t, mock.New())
}

func Test_MockUuid_Get_panics(t *testing.T) {
	mock := NewUuidMock(MockConfig{GetPanics: true})
	require.Panics(t, func() { mock.Get() })
}

func Test_MockUuid_Get_does_not_panic(t *testing.T) {
	mock := NewUuidMock(MockConfig{})
	require.Nil(t, mock.New())
	require.NotPanics(t, func() { mock.Get() })
}

func Test_MockUuid_MarshalBinary_returns_error(t *testing.T) {
	mock := NewUuidMock(MockConfig{MarshalBinaryReturnsError: true})
	require.Nil(t, mock.New())
	_, err := mock.MarshalBinary()
	require.NotNil(t, err)
}

func Test_MockUuid_MarshalBinary_does_not_return_error(t *testing.T) {
	mock := NewUuidMock(MockConfig{})
	require.Nil(t, mock.New())
	_, err := mock.MarshalBinary()
	require.NoError(t, err)
}

func Test_MockUuid_MarshalText_returns_error(t *testing.T) {
	mock := NewUuidMock(MockConfig{MarshalTextReturnsError: true})
	require.Nil(t, mock.New())
	_, err := mock.MarshalText()
	require.NotNil(t, err)
}

func Test_MockUuid_MarshalText_does_not_return_error(t *testing.T) {
	mock := NewUuidMock(MockConfig{})
	require.Nil(t, mock.New())
	_, err := mock.MarshalText()
	require.NoError(t, err)
}

func Test_MockUuid_UnmarshalBinary_returns_error(t *testing.T) {
	mock := NewUuidMock(MockConfig{UnmarshalBinaryReturnsError: true})
	require.Nil(t, mock.New())
	b, err := mock.MarshalBinary()
	require.NoError(t, err)
	require.NotNil(t, mock.UnmarshalBinary(b))
}

func Test_MockUuid_UnmarshalBinary_does_not_return_error(t *testing.T) {
	mock := NewUuidMock(MockConfig{})
	require.Nil(t, mock.New())
	b, err := mock.MarshalBinary()
	require.NoError(t, err)
	require.NoError(t, mock.UnmarshalBinary(b))
}

func Test_MockUuid_UnmarshalText_returns_error(t *testing.T) {
	mock := NewUuidMock(MockConfig{UnmarshalTextReturnsError: true})
	require.Nil(t, mock.New())
	b, err := mock.MarshalText()
	require.NoError(t, err)
	require.NotNil(t, mock.UnmarshalText(b))
}

func Test_MockUuid_UnmarshalText_does_not_return_error(t *testing.T) {
	mock := NewUuidMock(MockConfig{})
	require.Nil(t, mock.New())
	b, err := mock.MarshalText()
	require.NoError(t, err)
	require.NoError(t, mock.UnmarshalText(b))
}

func Test_MockUuid_Version_panics(t *testing.T) {
	mock := NewUuidMock(MockConfig{VersionPanics: true})
	require.Panics(t, func() { mock.Version() })
}

func Test_MockUuid_Version_does_not_panic(t *testing.T) {
	mock := NewUuidMock(MockConfig{})
	require.NoError(t, mock.New())
	require.NotPanics(t, func() { mock.Version() })
}
