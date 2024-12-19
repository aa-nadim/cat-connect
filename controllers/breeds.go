// controllers/breeds.go
package controllers

import (
	"cat-connect/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BreedsController struct {
	web.Controller
}

type Breed struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Origin       string `json:"origin"`
	Description  string `json:"description"`
	WikipediaURL string `json:"wikipedia_url"`
}

func (c *BreedsController) Get() {
	c.Layout = "layout.html"
	c.TplName = "breeds.html"
}

func (c *BreedsController) GetBreeds() {
	apiKey := "live_rtO7Nhjpuo0DmEaWTsE0J41ytL3FlYxLkJbSZNDG557WGS09hgLR2w0rjAWyNO5m" // Replace with your actual API key

	breedChan := make(chan []models.Breed)
	errChan := make(chan error)

	go func() {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/breeds", nil)
		if err != nil {
			errChan <- err
			return
		}

		req.Header.Add("x-api-key", apiKey)
		resp, err := client.Do(req)
		if err != nil {
			errChan <- err
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errChan <- err
			return
		}

		var breeds []models.Breed
		if err := json.Unmarshal(body, &breeds); err != nil {
			errChan <- err
			return
		}

		breedChan <- breeds
	}()

	select {
	case breeds := <-breedChan:
		c.Data["json"] = breeds
	case err := <-errChan:
		c.Data["json"] = map[string]string{"error": err.Error()}
	}

	c.ServeJSON()
}

func (c *BreedsController) GetBreedImages() {
	breedID := c.Ctx.Input.Param(":id") // Change this line
	if breedID == "" {
		c.Data["json"] = map[string]string{"error": "breed_id is required"}
		c.ServeJSON()
		return
	}

	apiKey := "live_rtO7Nhjpuo0DmEaWTsE0J41ytL3FlYxLkJbSZNDG557WGS09hgLR2w0rjAWyNO5m"

	imageChan := make(chan []models.CatImage)
	errChan := make(chan error)

	go func() {
		client := &http.Client{}
		url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s&limit=2&api_key=%s", breedID, apiKey)
		// println(url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			errChan <- err
			return
		}

		req.Header.Add("x-api-key", apiKey)
		resp, err := client.Do(req)
		if err != nil {
			errChan <- err
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errChan <- err
			return
		}

		var images []models.CatImage
		if err := json.Unmarshal(body, &images); err != nil {
			errChan <- err
			return
		}

		imageChan <- images
	}()

	select {
	case images := <-imageChan:
		c.Data["json"] = images
	case err := <-errChan:
		c.Data["json"] = map[string]string{"error": err.Error()}
	}

	c.ServeJSON()
}
