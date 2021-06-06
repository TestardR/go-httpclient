package gohttp

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequestBody(t *testing.T) {

}

func TestGetRequestHeaders(t *testing.T) {
	client := &httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	assert.Len(t, finalHeaders, 3)
	assert.Equal(t, finalHeaders.Get("X-Request-Id"), "ABC-123")
	assert.Equal(t, finalHeaders.Get("Content-Type"), "application/json")
	assert.Equal(t, finalHeaders.Get("User-Agent"), "cool-http-client")
}

func TestGetReuestBody(t *testing.T) {
	client := httpClient{}
	requestBody := []string{"one", "two"}
	t.Run("noBodyNilResponse", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)

		assert.Nil(t, body)
		assert.Nil(t, err)
	})

	t.Run("jsonBodyResponse", func(t *testing.T) {
		body, err := client.getRequestBody("application/json", requestBody)

		assert.Nil(t, err)
		assert.Equal(t, string(body), `["one","two"]`)
	})

	t.Run("xmlBodyResponse", func(t *testing.T) {
		body, err := client.getRequestBody("application/xml", requestBody)

		assert.Nil(t, err)
		assert.Equal(t, string(body), `<string>one</string><string>two</string>`)
	})

	t.Run("jsonBodyResponseDefault", func(t *testing.T) {
		body, err := client.getRequestBody("", requestBody)

		assert.Nil(t, err)
		assert.Equal(t, string(body), `["one","two"]`)
	})
}
