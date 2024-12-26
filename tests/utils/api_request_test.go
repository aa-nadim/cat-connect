package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"cat-connect/utils"

	"github.com/stretchr/testify/assert"
)

func TestMakeAPIRequest(t *testing.T) {
	t.Run("Successful GET request", func(t *testing.T) {
		// Create a test server
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "GET", r.Method)
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
			w.Write([]byte(`{"status": "success"}`))
		}))
		defer ts.Close()

		responseChan := utils.MakeAPIRequest("GET", ts.URL, nil, "test-api-key")

		select {
		case response := <-responseChan:
			assert.NoError(t, response.Error)
			assert.NotNil(t, response.Body)
			assert.Contains(t, string(response.Body), "success")
		case <-time.After(5 * time.Second):
			t.Fatal("Request timed out")
		}
	})

	t.Run("Invalid URL", func(t *testing.T) {
		responseChan := utils.MakeAPIRequest("GET", "invalid-url", nil, "")

		select {
		case response := <-responseChan:
			assert.Error(t, response.Error)
			assert.Nil(t, response.Body)
		case <-time.After(5 * time.Second):
			t.Fatal("Request timed out")
		}
	})

	t.Run("Server error response", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer ts.Close()

		responseChan := utils.MakeAPIRequest("GET", ts.URL, nil, "")

		select {
		case response := <-responseChan:
			assert.NoError(t, response.Error) // The request itself succeeded
			assert.NotNil(t, response.Body)   // Should contain error response body
		case <-time.After(5 * time.Second):
			t.Fatal("Request timed out")
		}
	})

	t.Run("POST request with body", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "POST", r.Method)
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
			assert.Equal(t, "test-api-key", r.Header.Get("x-api-key"))
			w.Write([]byte(`{"status": "created"}`))
		}))
		defer ts.Close()

		requestBody := []byte(`{"test": "data"}`)
		responseChan := utils.MakeAPIRequest("POST", ts.URL, requestBody, "test-api-key")

		select {
		case response := <-responseChan:
			assert.NoError(t, response.Error)
			assert.NotNil(t, response.Body)
			assert.Contains(t, string(response.Body), "created")
		case <-time.After(5 * time.Second):
			t.Fatal("Request timed out")
		}
	})

	t.Run("Request timeout", func(t *testing.T) {
		// Create a server that sleeps longer than the client timeout
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(11 * time.Second)
		}))
		defer ts.Close()

		responseChan := utils.MakeAPIRequest("GET", ts.URL, nil, "")

		select {
		case response := <-responseChan:
			assert.Error(t, response.Error)
			assert.Nil(t, response.Body)
		case <-time.After(12 * time.Second):
			t.Fatal("Test didn't complete in time")
		}
	})
}
