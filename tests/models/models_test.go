// tests/models/models_test.go
package models

import (
	"encoding/json"
	"testing"

	"cat-connect/models"

	"github.com/stretchr/testify/assert"
)

func TestVoteModel(t *testing.T) {
	vote := models.Vote{
		ImageID: "test-image",
		SubID:   "test-user",
		Value:   1,
	}

	jsonData, err := json.Marshal(vote)
	assert.NoError(t, err)

	var unmarshaledVote models.Vote
	err = json.Unmarshal(jsonData, &unmarshaledVote)
	assert.NoError(t, err)
	assert.Equal(t, vote, unmarshaledVote)
}
