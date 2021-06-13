package examples

import (
	"time"

	"github.com/TestardR/go-httpclient/go-httpclient/gohttp"
)

var (
	httpClient = getHttpClient()
)

// getHttpClient singleton
func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		Build()
	return client
}
