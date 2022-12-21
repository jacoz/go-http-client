package hc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpts(t *testing.T) {
	var tests = []struct {
		name  string
		input *options
		want  *options
	}{
		{
			"defaults",
			Opts(),
			&options{
				baseUrl:        "",
				timeout:        10,
				defaultHeaders: headers{},
				defaultQuery:   Q{},
			},
		},
		{
			"with data",
			Opts().
				BaseUrl("https://example.com/api/v1").
				Timeout(20).
				WithDefaultHeader("x-foo", "foo").
				WithDefaultHeader("x-bar", "bar").
				WithDefaultQuery(Q{
					"task": "test",
					"page": "1",
				}),
			&options{
				baseUrl: "https://example.com/api/v1",
				timeout: 20,
				defaultHeaders: headers{
					"x-foo": "foo",
					"x-bar": "bar",
				},
				defaultQuery: Q{
					"task": "test",
					"page": "1",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.input)
		})
	}
}
