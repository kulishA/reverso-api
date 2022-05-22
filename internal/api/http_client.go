package api

import (
	"github.com/corpix/uarand"
	"io"
	"net/http"
)

type httpClient struct {
	c *http.Client
}

func newHttpClient(c *http.Client) *httpClient {
	return &httpClient{c: c}
}

func (c *httpClient) post(url, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", uarand.GetRandom())

	return c.do(req)
}

func (c *httpClient) do(req *http.Request) (*http.Response, error) {
	return c.c.Do(req)
}
