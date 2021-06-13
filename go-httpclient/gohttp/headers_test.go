package gohttp

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequestHeaders(t *testing.T) {
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "mocked-http-client")
	client.builder = &clientBuilder{
		headers:   commonHeaders,
		userAgent: "cool-user-agent",
	}

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	assert.Len(t, finalHeaders, 3)
	assert.Equal(t, finalHeaders.Get("X-Request-Id"), "ABC-123")
	assert.Equal(t, finalHeaders.Get("Content-Type"), "application/json")
	assert.Equal(t, finalHeaders.Get("User-Agent"), "mocked-http-client")
}
