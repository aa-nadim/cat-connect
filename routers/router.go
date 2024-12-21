package routers

import (
	"cat-connect/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/api/cat-images", &controllers.BreedsController{}, "get:GetCatImages")
	beego.Router("/api/breeds", &controllers.BreedsController{}, "get:GetBreeds")
	beego.Router("/api/cat-images/by-breed", &controllers.BreedsController{}, "get:GetCatImagesByBreed")

	// Favorites routes
	beego.Router("/api/favorites", &controllers.FavoritesController{}, "post:AddFavorite")
	beego.Router("/api/favorites", &controllers.FavoritesController{}, "get:GetFavorites")
	beego.Router("/api/favorites/:id", &controllers.FavoritesController{}, "delete:DeleteFavorite")

	// New route for voting
	beego.Router("/*", &controllers.MainController{}, "get:Get")
	beego.Router("/api/votes", &controllers.VotesController{}, "post:Vote")
	beego.Router("/api/votes", &controllers.VotesController{}, "get:GetVotes")
}
