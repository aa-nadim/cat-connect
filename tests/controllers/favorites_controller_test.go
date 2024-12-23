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
	web.Router("/api/favorites", &controllers.FavoritesController{}, "get:GetFavorites")
	web.Router("/api/favorites/:id", &controllers.FavoritesController{}, "delete:DeleteFavorite")
}

func TestFavoritesController_GetFavorites(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/favorites?sub_id=test-user", nil)
	w := httptest.NewRecorder()

	// Mock the API request
	originalMakeAPIRequest := utils.MakeAPIRequest
	defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()
	utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
		responseChan := make(chan utils.APIResponse, 1)
		responseChan <- utils.APIResponse{
			Body: []byte(`[{"id": "fav1", "image_id": "img1"}]`),
		}
		return responseChan
	}

	web.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, "fav1", response[0]["id"])
	assert.Equal(t, "img1", response[0]["image_id"])
}

func TestFavoritesController_DeleteFavorite(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/api/favorites/fav1", nil)
	w := httptest.NewRecorder()

	// Mock the API request
	originalMakeAPIRequest := utils.MakeAPIRequest
	defer func() { utils.MakeAPIRequest = originalMakeAPIRequest }()
	utils.MakeAPIRequest = func(method, url string, body []byte, apiKey string) chan utils.APIResponse {
		responseChan := make(chan utils.APIResponse, 1)
		responseChan <- utils.APIResponse{
			Body: []byte(`{"message": "SUCCESS"}`),
		}
		return responseChan
	}

	web.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	// Debug print the response body
	t.Logf("Response Body: %s", w.Body.String())

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "SUCCESS", response["message"])
}
