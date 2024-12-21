package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type APIResponse struct {
	Body  []byte
	Error error
}

func MakeAPIRequest(method, url string, body []byte, apiKey string) <-chan APIResponse {
	responseChan := make(chan APIResponse)

	go func() {
		defer close(responseChan)

		req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			responseChan <- APIResponse{nil, fmt.Errorf("error creating request: %v", err)}
			return
		}

		req.Header.Set("x-api-key", apiKey)
		if method == "POST" {
			req.Header.Set("Content-Type", "application/json")
		}

		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			responseChan <- APIResponse{nil, fmt.Errorf("error making request: %v", err)}
			return
		}
		defer resp.Body.Close()

		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			responseChan <- APIResponse{nil, fmt.Errorf("error reading response body: %v", err)}
			return
		}

		responseChan <- APIResponse{responseBody, nil}
	}()

	return responseChan
}
