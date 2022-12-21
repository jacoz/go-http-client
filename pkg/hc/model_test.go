package hc

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestJson(t *testing.T) {
	var tests = []struct {
		name  string
		input D
		want  string
	}{
		{
			"empty",
			D{},
			`{}`,
		},
		{
			"with data",
			D{"foo": "bar"},
			`{"foo":"bar"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(strings.Builder)
			io.Copy(buf, Json(tt.input))

			assert.Equal(t, tt.want, buf.String())
		})
	}
}
