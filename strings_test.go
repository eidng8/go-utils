package utils

import (
	"errors"
	"io"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RandomAlphaNum_contains_only_alphanumeric(t *testing.T) {
	actual, err := RandomAlphaNum(10)
	require.Nil(t, err)
	require.Regexp(t, "^[a-zA-Z0-9]+$", actual)
}

func Test_RandomAlphaNum_returns_random_error(t *testing.T) {
	tmp := randomInts
	randomInts = func(_ io.Reader, _ *big.Int) (*big.Int, error) {
		return nil, errors.New("random error")
	}
	defer func() { randomInts = tmp }()
	_, err := RandomAlphaNum(10)
	require.NotNil(t, err)
	require.Equal(t, "random error", err.Error())
}

func Test_RandomPrintable_contains_only_printable(t *testing.T) {
	actual, err := RandomPrintable(10)
	require.Nil(t, err)
	require.Regexp(t, `^[\pL\pM\pN\pP\pS]+$`, actual)
}

func Test_RandomPrintable_returns_random_error(t *testing.T) {
	tmp := randomInts
	randomInts = func(_ io.Reader, _ *big.Int) (*big.Int, error) {
		return nil, errors.New("random error")
	}
	defer func() { randomInts = tmp }()
	_, err := RandomPrintable(10)
	require.NotNil(t, err)
	require.Equal(t, "random error", err.Error())
}

func Test_StringIndexOfAny_returns_index_of_first_substring(t *testing.T) {
	require.Equal(t, 1, StringIndexOfAny("test", []string{"e", "s"}))
}

func Test_StringIndexOfAny_returns_neg_one_if_not_found(t *testing.T) {
	require.Equal(t, -1, StringIndexOfAny("abc", []string{"e", "s"}))
}

func Test_StringIndexOfAny_returns_index_of_second_substring(t *testing.T) {
	require.Equal(t, 2, StringIndexOfAny("task", []string{"e", "s"}))
}

func Test_StringContainsAny_finds_first_substring(t *testing.T) {
	require.True(t, StringContainsAny("test", []string{"e", "s"}))
}

func Test_StringContainsAny_finds_second_substring(t *testing.T) {
	require.True(t, StringContainsAny("task", []string{"e", "s"}))
}
