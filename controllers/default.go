// controllers/default.go
package controllers

import (
	"cat-connect/models"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "Cat Connect"
	c.TplName = "index.tpl"
}

func (c *MainController) GetBreeds() {
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

func (c *MainController) GetBreedImages() {
	breedId := c.GetString("breed_id")
	if breedId == "" {
		c.Data["json"] = map[string]string{"error": "breed_id is required"}
		c.ServeJSON()
		return
	}

	apiKey := "live_rtO7Nhjpuo0DmEaWTsE0J41ytL3FlYxLkJbSZNDG557WGS09hgLR2w0rjAWyNO5m" // Replace with your actual API key

	imageChan := make(chan []models.CatImage)
	errChan := make(chan error)

	go func() {
		client := &http.Client{}
		url := "https://api.thecatapi.com/v1/images/search?breed_ids=" + breedId + "&limit=5"
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
