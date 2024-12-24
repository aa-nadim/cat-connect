package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "cat-connect/routers" // Import the routers package to initialize the routes

	beego "github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

func init() {
	beego.TestBeegoInit("../")
}

func TestMainRoute(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
}

func TestGetCatImages(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/cat-images", nil)
	w := httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
}

func TestAddFavorite(t *testing.T) {
	r, _ := http.NewRequest("POST", "/api/favorites", nil)
	w := httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
}

func TestVote(t *testing.T) {
	r, _ := http.NewRequest("POST", "/api/votes", nil)
	w := httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
}

func TestGetVotes(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/votes", nil)
	w := httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
}

func TestGetBreeds(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/breeds", nil)
	w := httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
}

func TestGetCatImagesByBreed(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/cat-images/by-breed", nil)
	w := httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
}

func TestGetFavorites(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/favorites", nil)
	w := httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
}

func TestDeleteFavorite(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/api/favorites/1", nil)
	w := httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
}
