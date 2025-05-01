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

/*
client := httpclient.New(10 * time.Second)
req, _ := http.NewRequest("GET", "https://api.com", nil)
resp, err := client.Do(req)
*/
