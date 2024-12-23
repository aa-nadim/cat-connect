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
