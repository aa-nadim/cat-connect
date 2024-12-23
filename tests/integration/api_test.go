package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"cat-connect/models"

	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

func TestFullAPIFlow(t *testing.T) {
	// 1. Get cat images
	r1, _ := http.NewRequest("GET", "/api/cat-images", nil)
	w1 := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(w1, r1)
	assert.Equal(t, http.StatusOK, w1.Code)

	var catImages []models.CatImage
	err := json.Unmarshal(w1.Body.Bytes(), &catImages)
	assert.NoError(t, err)
	assert.NotEmpty(t, catImages)

	// 2. Add favorite
	if len(catImages) > 0 {
		favorite := models.Favorite{
			ImageID: catImages[0].ID,
			SubID:   "test-user",
		}

		body, _ := json.Marshal(favorite)
		r2, _ := http.NewRequest("POST", "/api/favorites", bytes.NewBuffer(body))
		w2 := httptest.NewRecorder()
		web.BeeApp.Handlers.ServeHTTP(w2, r2)
		assert.Equal(t, http.StatusOK, w2.Code)

		var favoriteResponse map[string]interface{}
		err = json.Unmarshal(w2.Body.Bytes(), &favoriteResponse)
		assert.NoError(t, err)
		assert.Equal(t, "Success", favoriteResponse["message"])
	}

	// 3. Get favorites
	r3, _ := http.NewRequest("GET", "/api/favorites?sub_id=test-user", nil)
	w3 := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(w3, r3)
	assert.Equal(t, http.StatusOK, w3.Code)

	var favorites []models.Favorite
	err = json.Unmarshal(w3.Body.Bytes(), &favorites)
	assert.NoError(t, err)
	assert.NotEmpty(t, favorites)

	// 4. Delete favorite
	if len(favorites) > 0 {
		favoriteID := favorites[0].ID
		r4, _ := http.NewRequest("DELETE", "/api/favorites/"+favoriteID, nil)
		w4 := httptest.NewRecorder()
		web.BeeApp.Handlers.ServeHTTP(w4, r4)
		assert.Equal(t, http.StatusOK, w4.Code)

		var deleteResponse map[string]interface{}
		err = json.Unmarshal(w4.Body.Bytes(), &deleteResponse)
		assert.NoError(t, err)
		assert.Equal(t, "Success", deleteResponse["message"])
	}
}
