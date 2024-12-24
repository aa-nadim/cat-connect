package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
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

	fmt.Printf("Using API key: %s\n", "****"+apiKey[len(apiKey)-4:])

	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": fmt.Sprintf("Error reading request body: %v", err)}
		c.ServeJSON()
		return
	}

	fmt.Printf("Incoming request body: %s\n", string(body))

	var favorite models.Favorite
	if err := json.Unmarshal(body, &favorite); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": fmt.Sprintf("Error parsing request body: %v", err)}
		c.ServeJSON()
		return
	}

	fmt.Printf("Parsed favorite: %+v\n", favorite)

	favoriteJSON, err := json.Marshal(favorite)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Error preparing request"}
		c.ServeJSON()
		return
	}

	url := "https://api.thecatapi.com/v1/favourites"
	fmt.Printf("Making request to: %s\n", url)
	fmt.Printf("Request payload: %s\n", string(favoriteJSON))

	responseChan := utils.MakeAPIRequest("POST", url, favoriteJSON, apiKey)

	select {
	case response := <-responseChan:
		if response.Error != nil {
			fmt.Printf("API Error: %v\n", response.Error)
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": response.Error.Error()}
		} else {
			rawResponse := string(response.Body)
			fmt.Printf("Raw API Response: %s\n", rawResponse)

			// If response starts with a quote, it's probably an error message
			if strings.HasPrefix(rawResponse, "\"") {
				// Remove quotes and return as error
				errorMsg := strings.Trim(rawResponse, "\"")
				c.Ctx.Output.SetStatus(400)
				c.Data["json"] = map[string]string{"error": errorMsg}
			} else {
				var favoriteResponse struct {
					ID      int    `json:"id"`
					Message string `json:"message"`
				}
				if err := json.Unmarshal(response.Body, &favoriteResponse); err != nil {
					fmt.Printf("Error parsing response: %v\n", err)
					c.Ctx.Output.SetStatus(500)
					c.Data["json"] = map[string]string{"error": fmt.Sprintf("Failed to parse API response: %v", err)}
				} else {
					c.Ctx.Output.SetStatus(200)
					c.Data["json"] = favoriteResponse
				}
			}
		}
	case <-time.After(15 * time.Second):
		fmt.Println("Request timed out")
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
