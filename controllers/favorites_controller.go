package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"cat-connect/models"
	"cat-connect/utils"

	"github.com/beego/beego/v2/server/web"
)

type FavoritesController struct {
	web.Controller
}

func (c *FavoritesController) AddFavorite() {
	apiKey, _ := web.AppConfig.String("cat_api_key")

	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": fmt.Sprintf("Error reading request body: %v", err)}
		c.ServeJSON()
		return
	}

	var favorite models.Favorite
	if err := json.Unmarshal(body, &favorite); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": fmt.Sprintf("Error parsing request body: %v", err)}
		c.ServeJSON()
		return
	}

	url := "https://api.thecatapi.com/v1/favourites"
	responseChan := utils.MakeAPIRequest("POST", url, body, apiKey)

	select {
	case response := <-responseChan:
		if response.Error != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": response.Error.Error()}
		} else {
			c.Ctx.Output.SetStatus(200)
			c.Ctx.Output.Body(response.Body)
		}
	case <-time.After(15 * time.Second):
		c.Ctx.Output.SetStatus(504)
		c.Data["json"] = map[string]string{"error": "Request timed out"}
	}

	c.ServeJSON()
}

func (c *FavoritesController) GetFavorites() {
	apiKey, _ := web.AppConfig.String("cat_api_key")
	subID := c.GetString("sub_id")
	url := fmt.Sprintf("https://api.thecatapi.com/v1/favourites?sub_id=%s", subID)

	responseChan := utils.MakeAPIRequest("GET", url, nil, apiKey)

	select {
	case response := <-responseChan:
		if response.Error != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": response.Error.Error()}
		} else {
			c.Ctx.Output.SetStatus(200)
			c.Ctx.Output.Body(response.Body)
		}
	case <-time.After(15 * time.Second):
		c.Ctx.Output.SetStatus(504)
		c.Data["json"] = map[string]string{"error": "Request timed out"}
	}

	c.ServeJSON()
}

func (c *FavoritesController) DeleteFavorite() {
	apiKey, _ := web.AppConfig.String("cat_api_key")
	favoriteID := c.Ctx.Input.Param(":id")
	url := fmt.Sprintf("https://api.thecatapi.com/v1/favourites/%s", favoriteID)

	responseChan := utils.MakeAPIRequest("DELETE", url, nil, apiKey)

	select {
	case response := <-responseChan:
		if response.Error != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": response.Error.Error()}
		} else {
			c.Ctx.Output.SetStatus(200)
			c.Ctx.Output.Body(response.Body)
		}
	case <-time.After(15 * time.Second):
		c.Ctx.Output.SetStatus(504)
		c.Data["json"] = map[string]string{"error": "Request timed out"}
	}

	c.ServeJSON()
}
