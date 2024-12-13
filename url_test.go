package utils

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RequestBaseUrl(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
	assert.Nil(t, err)
	assert.Equal(t, "http://example.com", RequestBaseUrl(request).String())
	request, err = http.NewRequest(
		http.MethodGet, "https://example.com/a/b?abc=123", nil,
	)
	assert.Nil(t, err)
	assert.Equal(t, "https://example.com/a/b", RequestBaseUrl(request).String())
}

func Test_RequestUrlWithQueryParam(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	assert.Nil(t, err)
	u1 := request.URL
	u2 := RequestUrlWithQueryParam(request, "abc", "123")
	assert.NotEqual(t, u1, u2)
	assert.Equal(t, "https://example.com?abc=123", u2.String())
}

func Test_RequestUrlWithQueryParams(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	assert.Nil(t, err)
	params := map[string]string{"abc": "123", "def": "456"}
	u := RequestUrlWithQueryParams(request, params)
	assert.Equal(t, "https://example.com?abc=123&def=456", u.String())
}

func Test_RequestUrlWithoutQueryParams(t *testing.T) {
	request, err := http.NewRequest(
		http.MethodGet, "https://example.com?abc=123&def=456", nil,
	)
	assert.Nil(t, err)
	u := RequestUrlWithoutQueryParams(request, "abc")
	assert.Equal(t, "https://example.com?def=456", u.String())
}

func Test_RequestUriWithoutSchemeHost(t *testing.T) {
	request, err := http.NewRequest(
		http.MethodGet, "https://example.com/path/ep?abc=123&def=456", nil,
	)
	assert.Nil(t, err)
	u := RequestUriWithoutSchemeHost(request)
	assert.Equal(t, "/path/ep?abc=123&def=456", u.String())
}
