package hc

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type response struct {
	response *http.Response
}

// UnmarshalJson decodes a json response into a struct
func (r *response) UnmarshalJson(v any) error {
	body := r.response.Body
	defer body.Close()

	return json.NewDecoder(body).Decode(&v)
}

// UnmarshalXml decodes a xml response into a struct
func (r *response) UnmarshalXml(v any) error {
	body := r.Get().Body
	defer body.Close()

	return xml.Unmarshal(r.Debug(), &v)
}

// StatusCode returns the response status code
func (r *response) StatusCode() int {
	return r.response.StatusCode
}

// Ok is a shortcut to check if the response has the status 200
func (r *response) Ok() bool {
	return r.response.StatusCode == http.StatusOK
}

// Created is a shortcut to check if the response has the status 201
func (r *response) Created() bool {
	return r.response.StatusCode == http.StatusCreated
}

// NoContent is a shortcut to check if the response has the status 204
func (r *response) NoContent() bool {
	return r.response.StatusCode == http.StatusNoContent
}

// BadRequest is a shortcut to check if the response has the status 400
func (r *response) BadRequest() bool {
	return r.response.StatusCode == http.StatusBadRequest
}

// NotFound is a shortcut to check if the response has the status 404
func (r *response) NotFound() bool {
	return r.response.StatusCode == http.StatusNotFound
}

// Debug returns the response object
func (r *response) Debug() []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.response.Body)
	return buf.Bytes()
}

// Get returns the response object
func (r *response) Get() *http.Response {
	return r.response
}
