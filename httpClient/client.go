package httpclient

import (
	"net/http"
)

type Client interface {
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header) (*http.Response, error)
	Put(url string, headers http.Header) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
	SetRequestHeaders(headers http.Header)
}

type client struct {
	Headers http.Header
}

func (c *client) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}
func (c *client) Post(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, nil)
}
func (c *client) Put(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, nil)
}
func (c *client) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}

func (c *client) SetRequestHeaders(headers http.Header) {
	c.Headers = headers
}

func New() Client {
	return &client{}
}
