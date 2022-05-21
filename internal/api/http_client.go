package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/corpix/uarand"
	"io"
	"io/ioutil"
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

func (c *httpClient) PostRequest(path string, data interface{}) (response []byte, err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	preparedData := bytes.NewBuffer(jsonData)

	url := fmt.Sprintf("%s/%s", apiUrl, path)
	resp, err := c.post(url, contentType, preparedData)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, err
}
