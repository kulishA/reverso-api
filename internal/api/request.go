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

	client := http.Client{
		Timeout: time.Second * 5,
	}

	url := fmt.Sprintf("%s/%s", apiUrl, path)
	resp, err := client.Post(url, contentType, preparedData)
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
