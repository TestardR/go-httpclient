package gohttp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequestBody(t *testing.T) {
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
