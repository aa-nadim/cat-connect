package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"cat-connect/models"
	"cat-connect/utils"

	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
)

type VotesController struct {
	web.Controller
}

func (c *VotesController) GetCatImages() {
	apiKey, _ := beego.AppConfig.String("cat_api_key")
	url := "https://api.thecatapi.com/v1/images/search?limit=10"

	responseChan := utils.MakeAPIRequest("GET", url, nil, apiKey)

	select {
	case response := <-responseChan:
		if response.Error != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": response.Error.Error()}
		} else {
			var catImages []models.CatImage
			if err := json.Unmarshal(response.Body, &catImages); err != nil {
				c.Ctx.Output.SetStatus(500)
				c.Data["json"] = map[string]string{"error": fmt.Sprintf("Error parsing cat images: %v", err)}
			} else {
				c.Data["json"] = catImages
			}
		}
	case <-time.After(15 * time.Second):
		c.Ctx.Output.SetStatus(504)
		c.Data["json"] = map[string]string{"error": "Request timed out"}
	}

	c.ServeJSON()
}

func (c *VotesController) AddFavorite() {
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
			// Parse the response to get the new favorite's data
			var favoriteResponse struct {
				ID      int    `json:"id"`
				Message string `json:"message"`
			}
			if err := json.Unmarshal(response.Body, &favoriteResponse); err != nil {
				c.Ctx.Output.SetStatus(500)
				c.Data["json"] = map[string]string{"error": "Error parsing response"}
			} else {
				// Return both the success message and the new favorite's data
				c.Data["json"] = favoriteResponse
			}
		}
	case <-time.After(15 * time.Second):
		c.Ctx.Output.SetStatus(504)
		c.Data["json"] = map[string]string{"error": "Request timed out"}
	}

	c.ServeJSON()
}

func (c *VotesController) Vote() {
	apiKey, _ := web.AppConfig.String("cat_api_key")

	// Read the request body
	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": fmt.Sprintf("Error reading request body: %v", err)}
		c.ServeJSON()
		return
	}

	// Parse the request body into the Vote model
	var vote models.Vote
	if err := json.Unmarshal(body, &vote); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": fmt.Sprintf("Error parsing request body: %v", err)}
		c.ServeJSON()
		return
	}

	// Call the external API to create the vote
	url := "https://api.thecatapi.com/v1/votes"
	responseChan := utils.MakeAPIRequest("POST", url, body, apiKey)

	select {
	case response := <-responseChan:
		if response.Error != nil {
			// Handle API error
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": response.Error.Error()}
			c.ServeJSON()
			return
		}

		// Log and confirm successful vote creation
		fmt.Println("Vote created successfully. Response:", string(response.Body))

		// After successful vote creation, call GetVotes to retrieve the updated votes list
		c.GetVotes()
		return

	case <-time.After(15 * time.Second):
		// Handle timeout
		c.Ctx.Output.SetStatus(504)
		c.Data["json"] = map[string]string{"error": "Request timed out"}

		println("I am here................", c.Data)

		c.ServeJSON()
		return
	}
}

func (c *VotesController) GetVotes() {
	apiKey, _ := beego.AppConfig.String("cat_api_key")
	limit := c.GetString("limit")
	order := c.GetString("order")
	subID := c.GetString("sub_id")
	page := c.GetString("page")

	url := fmt.Sprintf("https://api.thecatapi.com/v1/votes?limit=%s&order=%s&sub_id=%s&page=%s", limit, order, subID, page)

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

	println("I am here................", c.Data)

	c.ServeJSON()
}
