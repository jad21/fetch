package fetch

import (
	"io"
	"net"
	"net/http"
	"time"
)

// DefaultTimeout defined timeout default for any request
const DefaultTimeout = time.Duration(30 * time.Second)

// Options default for any request in client
type Options struct {
	Header    http.Header
	Timeout   time.Duration
	Host      string
	Transport *http.Transport
}

type FuncOptions func(op *Options)

func WithHeader(Header http.Header) FuncOptions {
	return func(op *Options) {
		op.Header = Header
	}
}

func WithTimeout(timeout time.Duration) FuncOptions {
	return func(op *Options) {
		op.Timeout = timeout
	}
}

// DefaultOptions returns options with timeout defined
func DefaultOptions() *Options {
	opt := &Options{
		Timeout: DefaultTimeout,
		Header:  http.Header{},
	}
	opt.Transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: opt.Timeout,
		}).DialContext,
		TLSHandshakeTimeout: opt.Timeout,
	}
	return opt
}

// New get new fetcher and you need to specify the netTransport.
func New(fn ...FuncOptions) *Fetch {
	opt := DefaultOptions()
	for _, cb := range fn {
		cb(opt)
	}

	return &Fetch{
		Client: &http.Client{
			Timeout:   opt.Timeout,
			Transport: opt.Transport,
		},
		Option: opt,
	}
}

// Fetch use http default but defined with a timeout.
type Fetch struct {
	*http.Client
	Option *Options
}

// IsJSON add Content-Type as JSON in header.
func (f *Fetch) IsJSON() *Fetch {
	if f.Option.Header == nil {
		f.Option.Header = http.Header{}
	}

	f.Option.Header.Set("Content-Type", "application/json")
	return f
}

// makeResponse format response from generic request
func (f Fetch) makeResponse(resp *http.Response, err error) (*Response, error) {
	if resp == nil {
		resp = &http.Response{
			StatusCode: http.StatusGatewayTimeout,
			Status:     http.StatusText(http.StatusGatewayTimeout),
		}
	}
	return &Response{Response: resp}, err
}

// Do execute any kind of request
func (f *Fetch) Do(req *http.Request) (*Response, error) {
	if f.Option.Header != nil {
		req.Header = f.Option.Header
	}
	return f.makeResponse(f.Client.Do(req))
}

// Get do request with HTTP using HTTP Verb GET
func (f *Fetch) Get(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request GET: %s", err)
	}
	return f.Do(req)
}

// Post do request with HTTP using HTTP Verb POST
func (f *Fetch) Post(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request POST: %s", err)
	}
	return f.Do(req)
}

// Put do request with HTTP using HTTP Verb PUT
func (f *Fetch) Put(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request PUT: %s", err)
	}
	return f.Do(req)
}

// Delete do request with HTTP using HTTP Verb DELETE
func (f *Fetch) Delete(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request DELETE: %s", err)
	}
	return f.Do(req)
}

// Patch do request with HTTP using HTTP Verb PATCH
func (f *Fetch) Patch(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPatch, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request PATCH: %s", err)
	}
	return f.Do(req)
}

// Options do request with HTTP using HTTP Verb OPTIONS
func (f *Fetch) Options(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodOptions, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request OPTIONS: %s", err)
	}
	return f.Do(req)
}

// Global Instance
var self = New()

func Get(url string, reader io.Reader) (*Response, error) {
	return self.Get(url, reader)
}
func Post(url string, reader io.Reader) (*Response, error) {
	return self.Post(url, reader)
}
func Put(url string, reader io.Reader) (*Response, error) {
	return self.Put(url, reader)
}
func Delete(url string, reader io.Reader) (*Response, error) {
	return self.Delete(url, reader)
}

func IsJSON() *Fetch {
	return New().IsJSON()
}
