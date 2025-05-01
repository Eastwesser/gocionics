package httpclient

import (
	"net/http"
	"time"
)

type HttpClient struct {
	client *http.Client
}

func New(timeout time.Duration) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *HttpClient) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
