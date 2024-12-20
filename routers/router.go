package routers

import (
	"cat-connect/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.BreedsController{})

	// API routes
	beego.Router("/api/breeds", &controllers.BreedsController{}, "get:GetBreeds")
	beego.Router("/api/breeds/:id", &controllers.BreedsController{}, "get:GetBreedImages")

	// Favorites routes
	beego.Router("/api/favorites", &controllers.FavsController{}, "get:GetFavorites")
	beego.Router("/api/favorites/:favorite_id", &controllers.FavsController{}, "delete:DeleteFavorite")
	beego.Router("/api/favorites", &controllers.VotingController{}, "post:AddFavorite")
}
