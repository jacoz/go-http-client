package hc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReq(t *testing.T) {
	var tests = []struct {
		name  string
		input *request
		want  *request
	}{
		{
			"defaults",
			Req(),
			&request{
				headers: headers{},
				query:   nil,
			},
		},
		{
			"with data",
			Req().
				Query(Q{
					"task": "test",
					"page": "1",
				}).
				WithHeader("x-foo", "foo").
				WithHeader("x-bar", "bar").
				WithJsonContentType().
				WithBearerToken("foo"),
			&request{
				headers: headers{
					"x-foo":         "foo",
					"x-bar":         "bar",
					"Content-Type":  "application/json",
					"Authorization": "Bearer foo",
				},
				query: Q{
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
