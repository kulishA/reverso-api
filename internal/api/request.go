package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func postRequest(path string, data interface{}) (response []byte, err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	preparedData := bytes.NewBuffer(jsonData)

	url := fmt.Sprintf("%s/%s", apiUrl, path)

	client := newHttpClient(&http.Client{Timeout: 5 * time.Second})

	resp, err := client.post(url, contentType, preparedData)
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
