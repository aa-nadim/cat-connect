// controllers/favorites_controller.go

package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"cat-connect/utils"

	"github.com/beego/beego/v2/server/web"
)

type FavoritesController struct {
	web.Controller
}

func (c *FavoritesController) GetFavorites() {
	apiKey, _ := web.AppConfig.String("cat_api_key")
	subID := c.GetString("sub_id")

	// Add cache-busting parameter to ensure fresh data
	timestamp := time.Now().UnixNano()
	url := fmt.Sprintf("https://api.thecatapi.com/v1/favourites?sub_id=%s&_=%d", subID, timestamp)

	responseChan := utils.MakeAPIRequest("GET", url, nil, apiKey)

	select {
	case response := <-responseChan:
		if response.Error != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": response.Error.Error()}
		} else {
			// Parse the response to ensure it's valid JSON before sending
			var favorites []map[string]interface{}
			if err := json.Unmarshal(response.Body, &favorites); err != nil {
				c.Ctx.Output.SetStatus(500)
				c.Data["json"] = map[string]string{"error": "Error parsing favorites"}
			} else {
				c.Data["json"] = favorites
			}
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
