package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"cat-connect/controllers"
	"cat-connect/utils"

	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

func init() {
	web.LoadAppConfig("ini", "../../conf/app.conf")

	// Set up routes for testing
	web.Router("/api/breeds", &controllers.BreedsController{}, "get:GetBreeds")
	web.Router("/api/cat-images/by-breed", &controllers.BreedsController{}, "get:GetCatImagesByBreed")
}

func TestBreedsController_GetBreeds(t *testing.T) {
	// Valid response case
	t.Run("Valid Response", func(t *testing.T) {
		r, _ := http.NewRequest("GET", "/api/breeds", nil)
		w := httptest.NewRecorder()

		// Mock the API request
		originalMakeAPIRequest := utils.MakeAPIRequest
		defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()
		utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
			responseChan := make(chan utils.APIResponse, 1)
			responseChan <- utils.APIResponse{
				Body: []byte(`[{"id": "abys", "name": "Abyssinian"}]`),
			}
			return responseChan
		}

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
		assert.Equal(t, "Abyssinian", response[0]["name"])
	})

	// Error response from API
	t.Run("API Error Response", func(t *testing.T) {
		r, _ := http.NewRequest("GET", "/api/breeds", nil)
		w := httptest.NewRecorder()

		// Mock the API request
		originalMakeAPIRequest := utils.MakeAPIRequest
		defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()
		utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
			responseChan := make(chan utils.APIResponse, 1)
			responseChan <- utils.APIResponse{
				Error: assert.AnError,
			}
			return responseChan
		}

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestBreedsController_GetCatImagesByBreed(t *testing.T) {
	// Valid response case
	t.Run("Valid Response", func(t *testing.T) {
		r, _ := http.NewRequest("GET", "/api/cat-images/by-breed?breed_id=abys", nil)
		w := httptest.NewRecorder()

		// Mock the API request
		originalMakeAPIRequest := utils.MakeAPIRequest
		defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()
		utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
			responseChan := make(chan utils.APIResponse, 1)
			responseChan <- utils.APIResponse{
				Body: []byte(`[{"id": "image1", "url": "https://example.com/image1.jpg"}]`),
			}
			return responseChan
		}

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
		assert.Equal(t, "image1", response[0]["id"])
		assert.Equal(t, "https://example.com/image1.jpg", response[0]["url"])
	})

	// Missing breed_id
	t.Run("Missing Breed ID", func(t *testing.T) {
		r, _ := http.NewRequest("GET", "/api/cat-images/by-breed", nil)
		w := httptest.NewRecorder()

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Breed ID is required", response["error"])
	})

	// API error response
	t.Run("API Error Response", func(t *testing.T) {
		r, _ := http.NewRequest("GET", "/api/cat-images/by-breed?breed_id=abys", nil)
		w := httptest.NewRecorder()

		// Mock the API request
		originalMakeAPIRequest := utils.MakeAPIRequest
		defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()
		utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
			responseChan := make(chan utils.APIResponse, 1)
			responseChan <- utils.APIResponse{
				Error: assert.AnError,
			}
			return responseChan
		}

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	// Timeout scenario
	t.Run("Request Timeout", func(t *testing.T) {
		r, _ := http.NewRequest("GET", "/api/cat-images/by-breed?breed_id=abys", nil)
		w := httptest.NewRecorder()

		// Mock the API request with a timeout
		originalMakeAPIRequest := utils.MakeAPIRequest
		defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()
		utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
			responseChan := make(chan utils.APIResponse)
			return responseChan
		}

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusGatewayTimeout, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Request timed out", response["error"])
	})
}
