package utils

import (
	"cat-connect/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockRequestBody struct {
	Message string `json:"message"`
}

func TestMakeAPIRequest(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate the request headers
		if r.Header.Get("x-api-key") != "test-api-key" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		if r.Method == http.MethodPost {
			if r.Header.Get("Content-Type") != "application/json" {
				http.Error(w, "unsupported media type", http.StatusUnsupportedMediaType)
				return
			}

			// Validate the request body
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			var requestBody MockRequestBody
			if err := json.Unmarshal(body, &requestBody); err != nil {
				http.Error(w, "invalid JSON", http.StatusBadRequest)
				return
			}
			if requestBody.Message != "Hello, World!" {
				http.Error(w, "invalid message", http.StatusBadRequest)
				return
			}
		}

		// Respond with success
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success"}`))
	}))
	defer mockServer.Close()

	// Prepare the test data
	testBody := MockRequestBody{Message: "Hello, World!"}
	bodyBytes, _ := json.Marshal(testBody)

	// Call the function
	responseChan := utils.MakeAPIRequest(http.MethodPost, mockServer.URL, bodyBytes, "test-api-key")

	// Wait for the response
	select {
	case response := <-responseChan:
		if response.Error != nil {
			t.Fatalf("unexpected error: %v", response.Error)
		}

		// Validate the response body
		var responseBody map[string]string
		if err := json.Unmarshal(response.Body, &responseBody); err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}
		if responseBody["status"] != "success" {
			t.Errorf("expected status 'success', got '%s'", responseBody["status"])
		}
	case <-time.After(5 * time.Second):
		t.Fatal("test timed out waiting for response")
	}
}
