package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type VotingController struct {
	web.Controller
}

type FavoriteRequest struct {
	ImageID string `json:"image_id"`
	SubID   string `json:"sub_id"`
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

	client := &http.Client{}
	jsonBody, _ := json.Marshal(map[string]string{
		"image_id": req.ImageID,
		"sub_id":   req.SubID,
	})

	apiReq, err := http.NewRequest("POST", "https://api.thecatapi.com/v1/favourites", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	apiReq.Header.Set("Content-Type", "application/json")
	apiReq.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(apiReq)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	// Log the response for debugging
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		c.Ctx.Output.SetStatus(resp.StatusCode)
		c.Data["json"] = map[string]string{"error": string(body)}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = json.RawMessage(body)
	c.ServeJSON()
}
