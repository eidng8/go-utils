package utils

import (
	"crypto/tls"
	"net/http"
	"strings"
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

func Test_RequestFullUrl(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected string
	}{
		{
			"domain only",
			"http://example.com",
			"http://example.com/",
		},
		{
			"domain with trailing slash",
			"http://example.com/",
			"http://example.com/",
		},
		{
			"path only",
			"http://example.com/a/b",
			"http://example.com/a/b",
		},
		{
			"fragment only",
			"http://example.com#frag",
			"http://example.com/#frag",
		},
		{
			"fragment with slash",
			"http://example.com/#frag",
			"http://example.com/#frag",
		},
		{
			"querystring only",
			"https://example.com?abc=123",
			"https://example.com/?abc=123",
		},
		{
			"querystring with slash",
			"https://example.com/?abc=123",
			"https://example.com/?abc=123",
		},
		{
			"all parts",
			"https://example.com/a/b?abc=123#frag",
			"https://example.com/a/b?abc=123#frag",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				request, err := http.NewRequest(http.MethodGet, tt.url, nil)
				require.Nil(t, err)
				if strings.HasPrefix(tt.url, "https") {
					request.TLS = &tls.ConnectionState{}
				}
				require.Equal(t, tt.expected, RequestFullUrl(request))
			},
		)
	}
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
