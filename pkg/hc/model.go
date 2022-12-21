package hc

import (
	"bytes"
	"encoding/json"
	"io"
)

type (
	// Q is used for the query string mapping
	Q map[string]string

	// D is used to map request data (eg. json)
	D map[string]interface{}
)

func (d *D) jsonReader() io.Reader {
	v, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return bytes.NewReader(v)
}

// Json helper allows to easily map a request that internally is then converted to a reader with the json object
func Json(v D) io.Reader {
	return v.jsonReader()
}
