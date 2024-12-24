package utils

import (
	"net/http"
	"net/url"
)

// RequestBaseUrl returns the base URL of the request.
func RequestBaseUrl(request *http.Request) *url.URL {
	u := *request.URL
	u.RawQuery = ""
	return &u
}

// RequestUrlWithQueryParam returns a new URL instance with the given query
// parameter set.
func RequestUrlWithQueryParam(
	request *http.Request, name string, value string,
) *url.URL {
	nu := *request.URL
	return UrlWithQueryParam(nu, name, value)
}

// RequestUrlWithQueryParams returns a new URL instance with the given query
// parameters set.
func RequestUrlWithQueryParams(
	request *http.Request, params map[string]string,
) *url.URL {
	nu := *request.URL
	return UrlWithQueryParams(nu, params)
}

// RequestUrlWithoutQueryParams returns a new URL instance without specified
// query parameters.
func RequestUrlWithoutQueryParams(
	req *http.Request, params ...string,
) *url.URL {
	nu := *req.URL
	return UrlWithoutQueryParams(nu, params...)
}

// RequestUriWithoutSchemeHost returns a new URL instance without scheme and
// host (domain).
func RequestUriWithoutSchemeHost(req *http.Request) *url.URL {
	nu := *req.URL
	nu.Scheme = ""
	nu.Host = ""
	return &nu
}

// UrlWithQueryParam returns a URL with the given query parameters set.
func UrlWithQueryParam(u url.URL, name string, value string) *url.URL {
	query := u.Query()
	query.Set(name, value)
	u.RawQuery = query.Encode()
	return &u
}

// UrlWithQueryParams returns a URL with the given query parameters set.
func UrlWithQueryParams(u url.URL, params map[string]string) *url.URL {
	query := u.Query()
	for name, value := range params {
		query.Set(name, value)
	}
	u.RawQuery = query.Encode()
	return &u
}

// UrlWithoutQueryParams returns a URL without specified query parameters.
func UrlWithoutQueryParams(u url.URL, params ...string) *url.URL {
	query := u.Query()
	for _, param := range params {
		query.Del(param)
	}
	u.RawQuery = query.Encode()
	return &u
}