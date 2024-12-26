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
	r, _ := http.NewRequest("GET", "/api/cat-images", nil)
	w := httptest.NewRecorder()

	web.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.CatImage
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response)
}

func TestVotesController_GetCatImages_Error(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/cat-images", nil)
	w := httptest.NewRecorder()

	// Mock the API request to return an error
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
	t.Logf("Response Code: %d, Body: %s", w.Code, w.Body.String())
}

func TestVotesController_Vote(t *testing.T) {
	vote := models.Vote{
		ImageID: "test-image",
		SubID:   "test-user",
		Value:   1,
	}

	body, _ := json.Marshal(vote)
	r, _ := http.NewRequest("POST", "/api/votes", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	web.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestVotesController_Vote_InvalidInput(t *testing.T) {
	r, _ := http.NewRequest("POST", "/api/votes", bytes.NewBuffer([]byte(`invalid-json`)))
	w := httptest.NewRecorder()

	web.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	t.Logf("Response Code: %d, Body: %s", w.Code, w.Body.String())
}

func TestVotesController_Vote_MissingFields(t *testing.T) {
	vote := models.Vote{
		ImageID: "test-image", // SubID is missing
		Value:   1,
	}

	body, _ := json.Marshal(vote)
	r, _ := http.NewRequest("POST", "/api/votes", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	web.BeeApp.Handlers.ServeHTTP(w, r)

}

func TestVotesController_AddFavorite(t *testing.T) {
	favorite := models.Favorite{
		ImageID: "test-image",
		SubID:   "test-user",
	}

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
}

func TestVotesController_AddFavorite_APIFailure(t *testing.T) {
	favorite := models.Favorite{
		ImageID: "test-image",
		SubID:   "test-user",
	}

	body, _ := json.Marshal(favorite)
	r, _ := http.NewRequest("POST", "/api/favorites", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	// Mock the API request to fail
	originalMakeAPIRequest := utils.MakeAPIRequest
	defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()
	utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
		responseChan := make(chan utils.APIResponse, 1)
		responseChan <- utils.APIResponse{
			Body:  []byte(`{"message": "ERROR"}`),
			Error: http.ErrServerClosed,
		}
		return responseChan
	}

	web.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	t.Logf("Response Code: %d, Body: %s", w.Code, w.Body.String())

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
}
