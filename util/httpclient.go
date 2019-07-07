package util

import (
	"net/http"
	"time"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

const (
	retryWaitMin = 500 * time.Millisecond
	retryWaitMax = 1000 * time.Millisecond
)

var httpClient = http.Client{
	Transport: &http.Transport{MaxIdleConnsPerHost: 4},
}

// Httpclient retryablehttp.Clientを生成し、設定値を反映させて返す
func Httpclient() *retryablehttp.Client {
	rh := retryablehttp.NewClient()

	rh.HTTPClient = &httpClient
	rh.RetryMax = 3
	rh.RetryWaitMin = retryWaitMin
	rh.RetryWaitMax = retryWaitMax

	return rh
}
