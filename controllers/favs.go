package controllers

import (
	"cat-connect/utils"
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type FavsController struct {
	web.Controller
}

func (c *FavsController) Get() {
	c.Layout = "layout.html"
	c.TplName = "favs.html"
}

func (c *FavsController) GetFavorites() {
	apiKey, _ := web.AppConfig.String("APIKey")
	subID := c.GetString("sub_id")

	url := fmt.Sprintf("https://api.thecatapi.com/v1/favourites?sub_id=%s&order=DESC", subID)
	resp, err := utils.MakeRequest("GET", url, apiKey, nil)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = json.RawMessage(resp)
	c.ServeJSON()
}
