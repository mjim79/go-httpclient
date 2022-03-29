package examples

import (
	"github.com/mjim79/go-httpclient/gohttp"
	"time"
)

var (
	httpClient = getHttpclient()
)

func getHttpclient() gohttp.Client {

	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		Build()
	return client
}
