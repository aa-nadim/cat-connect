package controllers

import (
	"cat-connect/utils"
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type VotingController struct {
	web.Controller
}

type FavoriteRequest struct {
	ImageID string `json:"image_id"`
	SubID   string `json:"sub_id"`
}

func (c *VotingController) Get() {
	apiKey, _ := web.AppConfig.String("APIKey")
	c.Data["APIKey"] = apiKey

	c.Layout = "layout.html"
	c.TplName = "voting.html"
}

func (c *VotingController) AddFavorite() {
	apiKey, _ := web.AppConfig.String("APIKey")

	var req FavoriteRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	body, err := json.Marshal(req)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	resp, err := utils.MakeRequest("POST", "https://api.thecatapi.com/v1/favourites", apiKey, body)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = json.RawMessage(resp)
	c.ServeJSON()
}

func (c *VotingController) DeleteFavorite() {
	apiKey, _ := web.AppConfig.String("APIKey")
	favoriteID := c.Ctx.Input.Param(":favorite_id")

	url := fmt.Sprintf("https://api.thecatapi.com/v1/favourites/%s", favoriteID)
	resp, err := utils.MakeRequest("DELETE", url, apiKey, nil)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = json.RawMessage(resp)
	c.ServeJSON()
}
