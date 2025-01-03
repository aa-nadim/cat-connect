package controllers

import (
	"cat-connect/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.LoadAppConfig("ini", "../../conf/app.conf")

	// Set the view path
	web.BConfig.WebConfig.ViewsPath = "../../views"

	// Set up routes for testing
	web.Router("/", &controllers.MainController{})
}

func TestMainController_Get(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	web.BeeApp.Handlers.ServeHTTP(w, r)
}
