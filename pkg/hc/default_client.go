package hc

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)

type goHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type defaultClient struct {
	options options
	client  goHttpClient
}

func New(opts ...*options) *defaultClient {
	var o options
	if len(opts) > 0 {
		o = *opts[0]
	} else {
		v := Opts()
		o = *v
	}

	return &defaultClient{
		options: o,
		client: &http.Client{
			Timeout: time.Duration(o.timeout) * time.Second,
		},
	}
}

func (c *defaultClient) Get(ctx context.Context, endpoint string, q *Q, r ...*request) (*response, error) {
	return c.do(ctx, http.MethodGet, endpoint, q, nil, r...)
}

func (c *defaultClient) Post(ctx context.Context, endpoint string, body io.Reader, r ...*request) (*response, error) {
	return c.do(ctx, http.MethodPost, endpoint, nil, body, r...)
}

func (c *defaultClient) Patch(ctx context.Context, endpoint string, body io.Reader, r ...*request) (*response, error) {
	return c.do(ctx, http.MethodPatch, endpoint, nil, body, r...)
}

func (c *defaultClient) Put(ctx context.Context, endpoint string, body io.Reader, r ...*request) (*response, error) {
	return c.do(ctx, http.MethodPut, endpoint, nil, body, r...)
}

func (c *defaultClient) Delete(ctx context.Context, endpoint string, r ...*request) (*response, error) {
	return c.do(ctx, http.MethodDelete, endpoint, nil, nil, r...)
}

func (c *defaultClient) do(ctx context.Context, method, endpoint string, q *Q, body io.Reader, r ...*request) (*response, error) {
	fullUrl := c.options.baseUrl + endpoint

	req, err := http.NewRequestWithContext(ctx, method, fullUrl, body)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req, r...)
	c.setQueryString(req, q, r...)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return &response{res}, nil
}

func (c *defaultClient) setHeaders(req *http.Request, r ...*request) {
	if len(c.options.defaultHeaders) > 0 {
		for k, v := range c.options.defaultHeaders {
			req.Header.Set(k, v)
		}
	}

	if r != nil && len(r) > 0 && len(r[0].headers) > 0 {
		for k, v := range r[0].headers {
			req.Header.Set(k, v)
		}
	}
}

func (c *defaultClient) setQueryString(req *http.Request, q *Q, r ...*request) {
	res := url.Values{}

	if len(c.options.defaultQuery) > 0 {
		for k, v := range c.options.defaultQuery {
			res.Set(k, v)
		}
	}

	if len(r) > 0 && r != nil {
		for k, v := range r[0].query {
			res.Set(k, v)
		}
	}

	if q != nil {
		for k, v := range *q {
			res.Set(k, v)
		}
	}

	if len(res) > 0 {
		req.URL.RawQuery = res.Encode()
	}
}
