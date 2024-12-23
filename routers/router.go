package routers

import (
	"cat-connect/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// route for voting
	beego.Router("/*", &controllers.MainController{}, "get:Get")

	beego.Router("/api/cat-images", &controllers.VotesController{}, "get:GetCatImages")
	beego.Router("/api/favorites", &controllers.VotesController{}, "post:AddFavorite")
	beego.Router("/api/votes", &controllers.VotesController{}, "post:Vote")

	beego.Router("/api/votes", &controllers.VotesController{}, "get:GetVotes")

	// route for breeds
	beego.Router("/api/breeds", &controllers.BreedsController{}, "get:GetBreeds")
	beego.Router("/api/cat-images/by-breed", &controllers.BreedsController{}, "get:GetCatImagesByBreed")

	// route for favorites
	beego.Router("/api/favorites", &controllers.FavoritesController{}, "get:GetFavorites")
	beego.Router("/api/favorites/:id", &controllers.FavoritesController{}, "delete:DeleteFavorite")
}
