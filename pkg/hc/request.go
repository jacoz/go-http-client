package hc

import "fmt"

type headers map[string]string

type request struct {
	headers headers
	query   Q
}

// Req allows to define extra configuration for a request
func Req() *request {
	return &request{
		headers: headers{},
		query:   nil,
	}
}

// Query sets a query string for the request
func (r *request) Query(v Q) *request {
	r.query = v
	return r
}

// WithHeader sets an extra header for the request
func (r *request) WithHeader(k, v string) *request {
	r.headers[k] = v
	return r
}

// WithContentType is a shortcut for setting the Content-Type header
func (r *request) WithContentType(v string) *request {
	return r.WithHeader("Content-Type", v)
}

// WithJsonContentType is a shortcut for setting the Content-Type header for json requests
func (r *request) WithJsonContentType() *request {
	return r.WithContentType("application/json")
}

// WithBearerToken is a shortcut for setting the Authorization header, it will prepend to the token the "Bearer" keyword
func (r *request) WithBearerToken(v string) *request {
	return r.WithHeader("Authorization", fmt.Sprintf("Bearer %s", v))
}
