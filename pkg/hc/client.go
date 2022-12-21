package hc

import (
	"context"
	"io"
)

type Client interface {
	// Get performs a GET request
	Get(ctx context.Context, endpoint string, q *Q, r ...*request) (*response, error)

	// Post performs a POST request
	Post(ctx context.Context, endpoint string, body io.Reader, r ...*request) (*response, error)

	// Patch performs a PATCH request
	Patch(ctx context.Context, endpoint string, body io.Reader, r ...*request) (*response, error)

	// Put performs a PUT request
	Put(ctx context.Context, endpoint string, body io.Reader, r ...*request) (*response, error)

	// Delete performs a DELETE request
	Delete(ctx context.Context, endpoint string, r ...*request) (*response, error)
}
