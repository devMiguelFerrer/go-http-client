package httpclient

import (
	"net/http"
	"testing"
)

func TestGetReqHeaders(t *testing.T) {
	// Arrage
	client := client{}
	defaultHeaders := make(http.Header)
	addHeaders := make(http.Header)
	// Act
	defaultHeaders.Set("Authrozation", "Bearer JWT")
	defaultHeaders.Set("Content-Type", "application/json")
	addHeaders.Set("X-service", "proxy")
	client.Headers = defaultHeaders
	respHeaders := client.getReqHeaders(addHeaders)
	// Assert
	if len(respHeaders) != 3 {
		t.Error("Should be 3 headers")
	}
}
