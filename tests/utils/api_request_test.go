// tests/utils/api_request_test.go
package utils

import (
	"testing"
	"time"

	"cat-connect/utils"

	"github.com/stretchr/testify/assert"
)

func TestMakeAPIRequest(t *testing.T) {
	// Test GET request
	responseChan := utils.MakeAPIRequest("GET", "https://api.thecatapi.com/v1/images/search?limit=1", nil, "")

	select {
	case response := <-responseChan:
		assert.NoError(t, response.Error)
		assert.NotNil(t, response.Body)
	case <-time.After(15 * time.Second):
		t.Fatal("Request timed out")
	}
}
