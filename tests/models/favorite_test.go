// tests/models/favorite_test.go
package models

import (
	"encoding/json"
	"testing"

	"cat-connect/models"

	"github.com/stretchr/testify/assert"
)

func TestFavoriteModel(t *testing.T) {
	favorite := models.Favorite{
		ImageID: "test-image",
		SubID:   "test-user",
	}

	jsonData, err := json.Marshal(favorite)
	assert.NoError(t, err)

	var unmarshaledFavorite models.Favorite
	err = json.Unmarshal(jsonData, &unmarshaledFavorite)
	assert.NoError(t, err)
	assert.Equal(t, favorite, unmarshaledFavorite)
}
