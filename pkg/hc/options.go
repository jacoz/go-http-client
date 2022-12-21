package hc

type options struct {
	baseUrl        string
	timeout        int
	defaultHeaders headers
	defaultQuery   Q
}

// Opts sets global configuration options
func Opts() *options {
	return &options{
		baseUrl:        "",
		timeout:        10,
		defaultHeaders: headers{},
		defaultQuery:   Q{},
	}
}

// BaseUrl sets the base url that all the requests from this client will have
func (o *options) BaseUrl(v string) *options {
	o.baseUrl = v
	return o
}

// Timeout of the requests
func (o *options) Timeout(v int) *options {
	o.timeout = v
	return o
}

// WithDefaultHeader allows to define a bunch of headers that will be included in every request
func (o *options) WithDefaultHeader(k, v string) *options {
	o.defaultHeaders[k] = v
	return o
}

// WithDefaultQuery allows to define a bunch of query string parameters that will be included in every request
func (o *options) WithDefaultQuery(v Q) *options {
	o.defaultQuery = v
	return o
}
