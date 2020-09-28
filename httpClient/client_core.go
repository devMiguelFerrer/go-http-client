package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func (c *client) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	reqBody, err := c.getReqBody(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, errors.New("Wrong params to create a new Request")
	}

	reqHeaders := c.getReqHeaders(headers)
	req.Header = reqHeaders

	return client.Do(req)
}

func (c *client) getReqHeaders(headers http.Header) http.Header {
	reqHeaders := make(http.Header)
	for header, value := range c.Headers {
		if len(value) > 0 {
			reqHeaders.Set(header, value[0])
		}
	}
	for header, value := range headers {
		if len(value) > 0 {
			reqHeaders.Set(header, value[0])
		}
	}
	return reqHeaders
}

func (c *client) getReqBody(body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	return json.Marshal(body)
}
