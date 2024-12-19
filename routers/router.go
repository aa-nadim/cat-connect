package routers

import (
	"cat-connect/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.BreedsController{})
	beego.Router("/breeds", &controllers.BreedsController{})
	beego.Router("/voting", &controllers.VotingController{})
	beego.Router("/favs", &controllers.FavsController{})

	beego.Router("/api/breeds", &controllers.BreedsController{}, "get:GetBreeds")
	beego.Router("/api/breeds/:id", &controllers.BreedsController{}, "get:GetBreedImages")
	beego.Router("/api/voting", &controllers.VotingController{}, "get:GetVotingImages")
}
