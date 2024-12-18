// routers/router.go
package routers

import (
	"cat-connect/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.Router("/api/breeds", &controllers.MainController{}, "get:GetBreeds")
	web.Router("/api/breed-images", &controllers.MainController{}, "get:GetBreedImages")
}
