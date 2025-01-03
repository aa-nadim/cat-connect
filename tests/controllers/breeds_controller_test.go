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
	// Error during JSON unmarshalling
	t.Run("Error Parsing Breeds", func(t *testing.T) {
		r, _ := http.NewRequest("GET", "/api/breeds", nil)
		w := httptest.NewRecorder()

		// Mock the API request
		originalMakeAPIRequest := utils.MakeAPIRequest
		defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()
		utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
			responseChan := make(chan utils.APIResponse, 1)
			responseChan <- utils.APIResponse{
				Body: []byte(`Invalid JSON`),
			}
			return responseChan
		}

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["error"], "Error parsing breeds")
	})

	// Request timeout case
	t.Run("Request Timeout", func(t *testing.T) {
		r, _ := http.NewRequest("GET", "/api/breeds", nil)
		w := httptest.NewRecorder()

		// Mock the API request with a timeout
		originalMakeAPIRequest := utils.MakeAPIRequest
		defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()
		utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
			return make(chan utils.APIResponse)
		}

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusGatewayTimeout, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Request timed out", response["error"])
	})
}

func TestBreedsController_GetCatImagesByBreed(t *testing.T) {
	// Error during JSON unmarshalling
	t.Run("Error Parsing Cat Images", func(t *testing.T) {
		r, _ := http.NewRequest("GET", "/api/cat-images/by-breed?breed_id=abys", nil)
		w := httptest.NewRecorder()

		// Mock the API request
		originalMakeAPIRequest := utils.MakeAPIRequest
		defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()
		utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
			responseChan := make(chan utils.APIResponse, 1)
			responseChan <- utils.APIResponse{
				Body: []byte(`Invalid JSON`),
			}
			return responseChan
		}

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["error"], "Error parsing cat images")
	})

	// Request timeout case
	t.Run("Request Timeout", func(t *testing.T) {
		r, _ := http.NewRequest("GET", "/api/cat-images/by-breed?breed_id=abys", nil)
		w := httptest.NewRecorder()

		// Mock the API request with a timeout
		originalMakeAPIRequest := utils.MakeAPIRequest
		defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()
		utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
			return make(chan utils.APIResponse)
		}

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusGatewayTimeout, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Request timed out", response["error"])
	})
}
