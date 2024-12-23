// tests/models/breed_test.go
package models

import (
	"encoding/json"
	"testing"

	"cat-connect/models"

	"github.com/stretchr/testify/assert"
)

func TestBreedModel(t *testing.T) {
	breed := models.Breed{
		ID:          "beng",
		Name:        "Bengal",
		Origin:      "United States",
		Description: "Bengals are athletic cats with a strong, muscular build.",
	}

	jsonData, err := json.Marshal(breed)
	assert.NoError(t, err)

	var unmarshaledBreed models.Breed
	err = json.Unmarshal(jsonData, &unmarshaledBreed)
	assert.NoError(t, err)
	assert.Equal(t, breed, unmarshaledBreed)
}
