package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"cat-connect/models"
	"cat-connect/utils"

	"github.com/beego/beego/v2/server/web"
)

type BreedsController struct {
	web.Controller
}

func (c *BreedsController) GetBreeds() {
	url := "https://api.thecatapi.com/v1/breeds"
	responseChan := utils.MakeAPIRequest("GET", url, nil, "")

	select {
	case response := <-responseChan:
		if response.Error != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": response.Error.Error()}
		} else {
			var breeds []models.Breed
			if err := json.Unmarshal(response.Body, &breeds); err != nil {
				c.Ctx.Output.SetStatus(500)
				c.Data["json"] = map[string]string{"error": fmt.Sprintf("Error parsing breeds: %v", err)}
			} else {
				// Log the response data
				fmt.Println("Response Data:", breeds)

				// Assign the breeds to the response
				c.Data["json"] = breeds
			}
		}
	case <-time.After(15 * time.Second):
		c.Ctx.Output.SetStatus(504)
		c.Data["json"] = map[string]string{"error": "Request timed out"}
	}

	// Log the final response data being sent
	//fmt.Println("Final Response:", c.Data["json"])

	c.ServeJSON()
}

func (c *BreedsController) GetCatImagesByBreed() {
	apiKey, _ := web.AppConfig.String("cat_api_key")
	breedID := c.GetString("breed_id")
	if breedID == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Breed ID is required"}
		c.ServeJSON()
		return
	}

	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s&limit=5", breedID)

	responseChan := utils.MakeAPIRequest("GET", url, nil, apiKey)

	select {
	case response := <-responseChan:
		if response.Error != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": response.Error.Error()}
		} else {
			var catImages []map[string]interface{}
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
