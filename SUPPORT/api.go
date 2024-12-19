package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type APIController struct {
	web.Controller
}

type FavoriteRequest struct {
	ImageID string `json:"image_id"`
	SubID   string `json:"sub_id"`
}

func (c *APIController) GetRandomCat() {
	apiKey, _ := web.AppConfig.String("APIKey")

	resp, err := makeRequest("GET", "https://api.thecatapi.com/v1/images/search", apiKey, nil)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *APIController) GetBreeds() {
	apiKey, _ := web.AppConfig.String("APIKey")

	resp, err := makeRequest("GET", "https://api.thecatapi.com/v1/breeds", apiKey, nil)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *APIController) GetBreedImages() {
	breedID := c.Ctx.Input.Param(":breed_id")
	apiKey, _ := web.AppConfig.String("APIKey")

	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s&limit=10", breedID)
	resp, err := makeRequest("GET", url, apiKey, nil)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = resp
	c.ServeJSON()
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

	resp, err := makeRequest("POST", "https://api.thecatapi.com/v1/favourites", apiKey, body)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *FavsController) GetFavorites() {
	apiKey, _ := web.AppConfig.String("APIKey")
	subID := c.GetString("sub_id")

	url := fmt.Sprintf("https://api.thecatapi.com/v1/favourites?sub_id=%s&order=DESC", subID)
	resp, err := makeRequest("GET", url, apiKey, nil)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *VotingController) DeleteFavorite() {
	apiKey, _ := web.AppConfig.String("APIKey")
	favoriteID := c.Ctx.Input.Param(":favorite_id")

	url := fmt.Sprintf("https://api.thecatapi.com/v1/favourites/%s", favoriteID)
	resp, err := makeRequest("DELETE", url, apiKey, nil)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

func makeRequest(method, url, apiKey string, body []byte) (interface{}, error) {
	client := &http.Client{}

	var req *http.Request
	var err error

	if body != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(body))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("x-api-key", apiKey)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result interface{}
	if err := json.Unmarshal(responseBody, &result); err != nil {
		return nil, err
	}

	return result, nil
}
