package hc

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse_UnmarshalJson(t *testing.T) {
	var tests = []struct {
		name      string
		input     *http.Response
		want      interface{}
		wantError error
	}{
		{
			"empty",
			&http.Response{
				Body: io.NopCloser(strings.NewReader(``)),
			},
			nil,
			errors.New("EOF"),
		},
		{
			"empty object",
			&http.Response{
				Body: io.NopCloser(strings.NewReader(`{}`)),
			},
			map[string]interface{}{},
			nil,
		},
		{
			"object",
			&http.Response{
				Body: io.NopCloser(strings.NewReader(`{"foo":"bar"}`)),
			},
			map[string]interface{}{
				"foo": "bar",
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := response{
				response: tt.input,
			}

			var got interface{}
			err := res.UnmarshalJson(&got)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantError, err)
		})
	}
}

func TestResponse_StatusCode(t *testing.T) {
	var tests = []struct {
		name  string
		input *http.Response
		want  int
	}{
		{
			"status 1",
			&http.Response{
				StatusCode: 200,
			},
			200,
		},
		{
			"status 2",
			&http.Response{
				StatusCode: 500,
			},
			500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := response{
				response: tt.input,
			}
			got := res.StatusCode()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestResponse_Ok(t *testing.T) {
	var tests = []struct {
		name  string
		input *http.Response
		want  bool
	}{
		{
			"status 1",
			&http.Response{
				StatusCode: 200,
			},
			true,
		},
		{
			"status 2",
			&http.Response{
				StatusCode: 500,
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := response{
				response: tt.input,
			}
			got := res.Ok()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestResponse_Created(t *testing.T) {
	var tests = []struct {
		name  string
		input *http.Response
		want  bool
	}{
		{
			"status 1",
			&http.Response{
				StatusCode: 201,
			},
			true,
		},
		{
			"status 2",
			&http.Response{
				StatusCode: 500,
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := response{
				response: tt.input,
			}
			got := res.Created()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestResponse_NoContent(t *testing.T) {
	var tests = []struct {
		name  string
		input *http.Response
		want  bool
	}{
		{
			"status 1",
			&http.Response{
				StatusCode: 204,
			},
			true,
		},
		{
			"status 2",
			&http.Response{
				StatusCode: 500,
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := response{
				response: tt.input,
			}
			got := res.NoContent()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestResponse_BadRequest(t *testing.T) {
	var tests = []struct {
		name  string
		input *http.Response
		want  bool
	}{
		{
			"status 1",
			&http.Response{
				StatusCode: 400,
			},
			true,
		},
		{
			"status 2",
			&http.Response{
				StatusCode: 500,
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := response{
				response: tt.input,
			}
			got := res.BadRequest()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestResponse_NotFound(t *testing.T) {
	var tests = []struct {
		name  string
		input *http.Response
		want  bool
	}{
		{
			"status 1",
			&http.Response{
				StatusCode: 404,
			},
			true,
		},
		{
			"status 2",
			&http.Response{
				StatusCode: 500,
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := response{
				response: tt.input,
			}
			got := res.NotFound()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestResponse_Debug(t *testing.T) {
	var tests = []struct {
		name  string
		input *http.Response
		want  string
	}{
		{
			"object",
			&http.Response{
				Body: io.NopCloser(strings.NewReader(`{"foo":"bar"}`)),
			},
			`{"foo":"bar"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := response{
				response: tt.input,
			}

			assert.Equal(t, tt.want, res.Debug())
		})
	}
}

func TestResponse_Get(t *testing.T) {
	var tests = []struct {
		name  string
		input *http.Response
		want  *http.Response
	}{
		{
			"status 1",
			&http.Response{
				StatusCode: 200,
			},
			&http.Response{
				StatusCode: 200,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := response{
				response: tt.input,
			}
			got := res.Get()
			assert.Equal(t, tt.want, got)
		})
	}
}
