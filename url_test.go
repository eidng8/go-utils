package utils

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RequestBaseUrl(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
	require.Nil(t, err)
	require.Equal(t, "http://example.com", RequestBaseUrl(request).String())
	request, err = http.NewRequest(
		http.MethodGet, "https://example.com/a/b?abc=123", nil,
	)
	require.Nil(t, err)
	require.Equal(
		t, "https://example.com/a/b", RequestBaseUrl(request).String(),
	)
}

func Test_RequestUrlWithQueryParam(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	require.Nil(t, err)
	u1 := request.URL
	u2 := RequestUrlWithQueryParam(request, "abc", "123")
	require.NotEqual(t, u1, u2)
	require.Equal(t, "https://example.com?abc=123", u2.String())
}

func Test_RequestUrlWithQueryParams(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	require.Nil(t, err)
	params := map[string]string{"abc": "123", "def": "456"}
	u := RequestUrlWithQueryParams(request, params)
	require.Equal(t, "https://example.com?abc=123&def=456", u.String())
}

func Test_RequestUrlWithoutQueryParams(t *testing.T) {
	request, err := http.NewRequest(
		http.MethodGet, "https://example.com?abc=123&def=456", nil,
	)
	require.Nil(t, err)
	u := RequestUrlWithoutQueryParams(request, "abc")
	require.Equal(t, "https://example.com?def=456", u.String())
}

func Test_RequestUriWithoutSchemeHost(t *testing.T) {
	request, err := http.NewRequest(
		http.MethodGet, "https://example.com/path/ep?abc=123&def=456", nil,
	)
	require.Nil(t, err)
	u := RequestUriWithoutSchemeHost(request)
	require.Equal(t, "/path/ep?abc=123&def=456", u.String())
}
