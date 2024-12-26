package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"cat-connect/controllers"
	"cat-connect/models"
	"cat-connect/utils"

	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

func init() {
	web.LoadAppConfig("ini", "../../conf/app.conf")

	// Set up routes for testing
	web.Router("/api/cat-images", &controllers.VotesController{}, "get:GetCatImages")
	web.Router("/api/votes", &controllers.VotesController{}, "post:Vote")
	web.Router("/api/favorites", &controllers.VotesController{}, "post:AddFavorite")
	web.Router("/api/votes", &controllers.VotesController{}, "get:GetVotes")
}

func TestVotesController_GetCatImages(t *testing.T) {
	t.Run("Successful Response", func(t *testing.T) {
		r, _ := http.NewRequest("GET", "/api/cat-images", nil)
		w := httptest.NewRecorder()

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.CatImage
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("Error Response", func(t *testing.T) {
		// Simulating an error scenario (like DB failure or other unexpected issues)
		originalMakeAPIRequest := utils.MakeAPIRequest
		defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()
		utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
			responseChan := make(chan utils.APIResponse, 1)
			responseChan <- utils.APIResponse{
				Error: http.ErrServerClosed,
			}
			return responseChan
		}

		r, _ := http.NewRequest("GET", "/api/cat-images", nil)
		w := httptest.NewRecorder()

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestVotesController_Vote(t *testing.T) {
	vote := models.Vote{
		ImageID: "test-image",
		SubID:   "test-user",
		Value:   1,
	}

	t.Run("Successful Vote", func(t *testing.T) {
		body, _ := json.Marshal(vote)
		r, _ := http.NewRequest("POST", "/api/votes", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Invalid JSON Input", func(t *testing.T) {
		invalidBody := []byte(`invalid-json`)
		r, _ := http.NewRequest("POST", "/api/votes", bytes.NewBuffer(invalidBody))
		w := httptest.NewRecorder()

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code) // Expecting a BadRequest error for invalid JSON
	})

	t.Run("Missing Required Fields (ImageID)", func(t *testing.T) {
		invalidVote := models.Vote{
			SubID: "test-user",
			Value: 1,
		} // Missing ImageID
		body, _ := json.Marshal(invalidVote)
		r, _ := http.NewRequest("POST", "/api/votes", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		web.BeeApp.Handlers.ServeHTTP(w, r)

	})

	t.Run("Missing Required Fields (Value)", func(t *testing.T) {
		invalidVote := models.Vote{
			ImageID: "test-image",
			SubID:   "test-user",
		} // Missing Value
		body, _ := json.Marshal(invalidVote)
		r, _ := http.NewRequest("POST", "/api/votes", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		web.BeeApp.Handlers.ServeHTTP(w, r)

	})
}

func TestVotesController_AddFavorite(t *testing.T) {
	favorite := models.Favorite{
		ImageID: "test-image",
		SubID:   "test-user",
	}

	t.Run("Successful Favorite", func(t *testing.T) {
		body, _ := json.Marshal(favorite)
		r, _ := http.NewRequest("POST", "/api/favorites", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		// Mock the API request
		originalMakeAPIRequest := utils.MakeAPIRequest
		defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()
		utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
			responseChan := make(chan utils.APIResponse, 1)
			responseChan <- utils.APIResponse{
				Body: []byte(`{"id": 123, "message": "SUCCESS"}`),
			}
			return responseChan
		}

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "SUCCESS", response["message"])
		assert.Equal(t, float64(123), response["id"]) // JSON unmarshals numbers to float64 by default
	})

	t.Run("Invalid JSON Input", func(t *testing.T) {
		invalidBody := []byte(`invalid-json`)
		r, _ := http.NewRequest("POST", "/api/favorites", bytes.NewBuffer(invalidBody))
		w := httptest.NewRecorder()

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Missing Required Fields (SubID)", func(t *testing.T) {
		invalidFavorite := models.Favorite{
			ImageID: "test-image",
		} // Missing SubID
		body, _ := json.Marshal(invalidFavorite)
		r, _ := http.NewRequest("POST", "/api/favorites", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Favorite Service Error", func(t *testing.T) {
		body, _ := json.Marshal(favorite)
		r, _ := http.NewRequest("POST", "/api/favorites", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		// Mock the API request to simulate an error
		originalMakeAPIRequest := utils.MakeAPIRequest
		defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()
		utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
			responseChan := make(chan utils.APIResponse, 1)
			responseChan <- utils.APIResponse{
				Error: http.ErrServerClosed,
			}
			return responseChan
		}

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestVotesController_GetVotes(t *testing.T) {
	limit := "10"
	order := "asc"
	subID := "test-sub-id"
	page := "1"

	url := "/api/votes?limit=" + limit + "&order=" + order + "&sub_id=" + subID + "&page=" + page

	// Mock the API request to simulate various scenarios
	originalMakeAPIRequest := utils.MakeAPIRequest
	defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()

	t.Run("Successful Response", func(t *testing.T) {
		utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
			responseChan := make(chan utils.APIResponse, 1)
			responseChan <- utils.APIResponse{
				Body: []byte(`[{"id": 1, "image_id": "abc", "value": 1}]`),
			}
			return responseChan
		}

		r, _ := http.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("Error Response", func(t *testing.T) {
		utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
			responseChan := make(chan utils.APIResponse, 1)
			responseChan <- utils.APIResponse{
				Error: http.ErrServerClosed,
			}
			return responseChan
		}

		r, _ := http.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

	})

	t.Run("Timeout Response", func(t *testing.T) {
		utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
			return make(chan utils.APIResponse) // Simulate no response
		}

		r, _ := http.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()

		web.BeeApp.Handlers.ServeHTTP(w, r)

		assert.Equal(t, http.StatusGatewayTimeout, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Request timed out", response["error"])
	})

	t.Run("Invalid Query Parameters", func(t *testing.T) {
		// Passing invalid parameters to see how the endpoint responds
		url := "/api/votes?limit=not-a-number&order=unknown-order&sub_id=" + subID + "&page=-1"
		r, _ := http.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()

		web.BeeApp.Handlers.ServeHTTP(w, r)

	})
}
