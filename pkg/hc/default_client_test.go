package hc

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/jacoz/go-http-client/pkg/hc/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestDefaultClient_Get(t *testing.T) {
	ctx := context.Background()

	var tests = []struct {
		name      string
		options   *options
		endpoint  string
		query     *Q
		request   *request
		want      *response
		wantError error
	}{
		{
			"should return an error",
			Opts().BaseUrl("https://example.com/api/v1"),
			"/error",
			&Q{},
			Req(),
			nil,
			errors.New("foo"),
		},
		{
			"should return a successful response",
			Opts().BaseUrl("https://example.com/api/v1"),
			"/foo/bar",
			&Q{
				"foo": "bar",
			},
			Req().WithHeader("req", "true"),
			&response{
				response: &http.Response{},
			},
			nil,
		},
		{
			"should include defaults configuration",
			Opts().
				BaseUrl("https://example.com/api/v1").
				WithDefaultHeader("foo", "bar").
				WithDefaultQuery(Q{"default": "query"}),
			"/foo/bar",
			&Q{
				"foo": "bar",
			},
			Req(),
			&response{
				response: &http.Response{},
			},
			nil,
		},
	}

	goHttpClientMock := mocks.NewGoHttpClient(t)

	goHttpClientMock.On("Do", mock.MatchedBy(func(req *http.Request) bool {
		return req.URL.String() == "https://example.com/api/v1/error"
	})).Return(nil, errors.New("foo"))

	goHttpClientMock.On("Do", mock.MatchedBy(func(req *http.Request) bool {
		return req.Context() == ctx &&
			req.Method == http.MethodGet &&
			req.URL.String() == "https://example.com/api/v1/foo/bar?foo=bar" &&
			req.Body == nil &&
			req.Header.Get("req") == "true"
	})).Return(&http.Response{}, nil)

	goHttpClientMock.On("Do", mock.MatchedBy(func(req *http.Request) bool {
		return req.Context() == ctx &&
			req.Method == http.MethodGet &&
			req.URL.String() == "https://example.com/api/v1/foo/bar?default=query&foo=bar" &&
			req.Body == nil &&
			req.Header.Get("foo") == "bar"
	})).Return(&http.Response{}, nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.options)
			c.client = goHttpClientMock
			got, err := c.Get(ctx, tt.endpoint, tt.query, tt.request)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantError, err)
		})
	}
}

func TestDefaultClient_Post(t *testing.T) {
	ctx := context.Background()

	var tests = []struct {
		name      string
		options   *options
		endpoint  string
		body      io.Reader
		want      *response
		wantError error
	}{
		{
			"should return an error",
			Opts().BaseUrl("https://example.com/api/v1"),
			"/error",
			nil,
			nil,
			errors.New("foo"),
		},
		{
			"should return a successful response",
			Opts().BaseUrl("https://example.com/api/v1"),
			"/foo/bar",
			strings.NewReader(`{"foo": "bar"}`),
			&response{
				response: &http.Response{},
			},
			nil,
		},
		{
			"should include defaults configuration",
			Opts().
				BaseUrl("https://example.com/api/v1").
				WithDefaultHeader("foo", "bar").
				WithDefaultQuery(Q{"default": "query"}),
			"/foo/bar2",
			strings.NewReader(`{"foo": "bar 2"}`),
			&response{
				response: &http.Response{},
			},
			nil,
		},
	}

	goHttpClientMock := mocks.NewGoHttpClient(t)

	goHttpClientMock.On("Do", mock.MatchedBy(func(req *http.Request) bool {
		return req.URL.String() == "https://example.com/api/v1/error"
	})).Return(nil, errors.New("foo"))

	goHttpClientMock.On("Do", mock.MatchedBy(func(req *http.Request) bool {
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		body := buf.String()

		fmt.Println(body)

		return req.Context() == ctx &&
			req.Method == http.MethodPost &&
			req.URL.String() == "https://example.com/api/v1/foo/bar" &&
			body == `{"foo": "bar"}`
	})).Return(&http.Response{}, nil)

	goHttpClientMock.On("Do", mock.MatchedBy(func(req *http.Request) bool {
		return req.Context() == ctx &&
			req.Method == http.MethodPost &&
			req.URL.String() == "https://example.com/api/v1/foo/bar2?default=query" &&
			req.Header.Get("foo") == "bar"
	})).Return(&http.Response{}, nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.options)
			c.client = goHttpClientMock
			got, err := c.Post(ctx, tt.endpoint, tt.body)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantError, err)
		})
	}
}

func TestDefaultClient_Patch(t *testing.T) {
	ctx := context.Background()

	var tests = []struct {
		name      string
		options   *options
		endpoint  string
		body      io.Reader
		want      *response
		wantError error
	}{
		{
			"should return an error",
			Opts().BaseUrl("https://example.com/api/v1"),
			"/error",
			nil,
			nil,
			errors.New("foo"),
		},
		{
			"should return a successful response",
			Opts().BaseUrl("https://example.com/api/v1"),
			"/foo/bar",
			strings.NewReader(`{"foo": "bar"}`),
			&response{
				response: &http.Response{},
			},
			nil,
		},
	}

	goHttpClientMock := mocks.NewGoHttpClient(t)

	goHttpClientMock.On("Do", mock.MatchedBy(func(req *http.Request) bool {
		return req.URL.String() == "https://example.com/api/v1/error"
	})).Return(nil, errors.New("foo"))

	goHttpClientMock.On("Do", mock.MatchedBy(func(req *http.Request) bool {
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		body := buf.String()

		return req.Context() == ctx &&
			req.Method == http.MethodPatch &&
			req.URL.String() == "https://example.com/api/v1/foo/bar" &&
			body == `{"foo": "bar"}`
	})).Return(&http.Response{}, nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.options)
			c.client = goHttpClientMock
			got, err := c.Patch(ctx, tt.endpoint, tt.body)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantError, err)
		})
	}
}

func TestDefaultClient_Put(t *testing.T) {
	ctx := context.Background()

	var tests = []struct {
		name      string
		options   *options
		endpoint  string
		body      io.Reader
		want      *response
		wantError error
	}{
		{
			"should return an error",
			Opts().BaseUrl("https://example.com/api/v1"),
			"/error",
			nil,
			nil,
			errors.New("foo"),
		},
		{
			"should return a successful response",
			Opts().BaseUrl("https://example.com/api/v1"),
			"/foo/bar",
			strings.NewReader(`{"foo": "bar"}`),
			&response{
				response: &http.Response{},
			},
			nil,
		},
	}

	goHttpClientMock := mocks.NewGoHttpClient(t)

	goHttpClientMock.On("Do", mock.MatchedBy(func(req *http.Request) bool {
		return req.URL.String() == "https://example.com/api/v1/error"
	})).Return(nil, errors.New("foo"))

	goHttpClientMock.On("Do", mock.MatchedBy(func(req *http.Request) bool {
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		body := buf.String()

		return req.Context() == ctx &&
			req.Method == http.MethodPut &&
			req.URL.String() == "https://example.com/api/v1/foo/bar" &&
			body == `{"foo": "bar"}`
	})).Return(&http.Response{}, nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.options)
			c.client = goHttpClientMock
			got, err := c.Put(ctx, tt.endpoint, tt.body)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantError, err)
		})
	}
}

func TestDefaultClient_Delete(t *testing.T) {
	ctx := context.Background()

	var tests = []struct {
		name      string
		options   *options
		endpoint  string
		want      *response
		wantError error
	}{
		{
			"should return an error",
			Opts().BaseUrl("https://example.com/api/v1"),
			"/error",
			nil,
			errors.New("foo"),
		},
		{
			"should return a successful response",
			Opts().BaseUrl("https://example.com/api/v1"),
			"/foo/bar",
			&response{
				response: &http.Response{},
			},
			nil,
		},
	}

	goHttpClientMock := mocks.NewGoHttpClient(t)

	goHttpClientMock.On("Do", mock.MatchedBy(func(req *http.Request) bool {
		return req.URL.String() == "https://example.com/api/v1/error"
	})).Return(nil, errors.New("foo"))

	goHttpClientMock.On("Do", mock.MatchedBy(func(req *http.Request) bool {
		return req.Context() == ctx &&
			req.Method == http.MethodDelete &&
			req.URL.String() == "https://example.com/api/v1/foo/bar" &&
			req.Body == nil
	})).Return(&http.Response{}, nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.options)
			c.client = goHttpClientMock
			got, err := c.Delete(ctx, tt.endpoint)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantError, err)
		})
	}
}
