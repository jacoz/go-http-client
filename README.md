# go-http-client

A simple wrapper around the Go `net/http` client that provides a fluent interface for making HTTP requests.

## Installation

```bash
go get github.com/jacoz/go-http-client
```

## Usage

### Quick Start

```go
package main

import (
	"context"
	"fmt"
	"github.com/jacoz/go-http-client/pkg/hc"
)

func main() {
	client := hc.New()
	ctx := context.Background()

	res, err := client.Get(ctx, "https://api.example.com/users", nil)
	if err != nil {
		panic(err)
	}

	if res.Ok() {
		fmt.Println("Success!")
	}
}
```

### Configuration (Options)

You can configure the client using the `hc.Opts()` builder pattern:

```go
client := hc.New(
	hc.Opts().
		BaseUrl("https://api.example.com").
		Timeout(30). // Timeout in seconds
		WithDefaultHeader("User-Agent", "my-app/1.0").
		WithDefaultQuery(hc.Q{"api_key": "secret"}),
)
```

### Making Requests

The client supports `Get`, `Post`, `Put`, `Patch`, and `Delete` methods. Each method accepts a context, an endpoint (relative to BaseURL if set), and optional configuration.

#### GET

```go
// Simple GET request
res, err := client.Get(ctx, "/users", nil)

// GET with query parameters
res, err := client.Get(ctx, "/users", &hc.Q{"page": "1", "limit": "10"})

// GET with custom headers
res, err := client.Get(ctx, "/users", nil, hc.Req().WithHeader("X-Custom", "value"))
```

#### POST (JSON)

Use the `hc.Json()` helper and `hc.D` map for easy JSON payloads.

```go
data := hc.D{
	"name": "John Doe",
	"email": "john@example.com",
}

res, err := client.Post(ctx, "/users", hc.Json(data), hc.Req().WithJsonContentType())
```

### Response Handling

The `response` object provides helpful methods to check status codes and unmarshal JSON/XML.

```go
// Check status codes
if res.Ok() { /* 200 OK */ }
if res.Created() { /* 201 Created */ }
if res.NoContent() { /* 204 No Content */ }
if res.BadRequest() { /* 400 Bad Request */ }
if res.NotFound() { /* 404 Not Found */ }

// Get raw status code
code := res.StatusCode()

// Unmarshal JSON
var user User
if err := res.UnmarshalJson(&user); err != nil {
    // handle error
}

// Access raw http.Response
rawRes := res.Get()
```

## API Reference

### Client Interface

- `Get(ctx, endpoint, q, ...r)`
- `Post(ctx, endpoint, body, ...r)`
- `Patch(ctx, endpoint, body, ...r)`
- `Put(ctx, endpoint, body, ...r)`
- `Delete(ctx, endpoint, ...r)`

### Helpers

- `hc.opts()`: Build client options.
- `hc.Req()`: Build request-specific options (headers, query).
- `hc.Q`: Type alias for `map[string]string` (Query parameters).
- `hc.D`: Type alias for `map[string]interface{}` (Data/JSON).
- `hc.Json(D)`: Converts `hc.D` map to `io.Reader` for request body.

### Request Builder

- `Query(Q)`: Set query parameters.
- `WithHeader(k, v)`: Add a header.
- `WithContentType(v)`: Set Content-Type header.
- `WithJsonContentType()`: Set Content-Type to `application/json`.
- `WithBearerToken(token)`: Set Authorization header with Bearer token.
