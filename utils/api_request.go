package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

type APIResponse struct {
	Body  []byte
	Error error
}

var MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan APIResponse {
	responseChan := make(chan APIResponse, 1)
	go func() {
		req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			responseChan <- APIResponse{Error: err}
			return
		}
		req.Header.Set("x-api-key", apiKey)
		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			responseChan <- APIResponse{Error: err}
			return
		}
		defer resp.Body.Close()
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			responseChan <- APIResponse{Error: err}
			return
		}
		responseChan <- APIResponse{Body: responseBody}
	}()
	return responseChan
}
