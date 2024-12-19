package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type VotingController struct {
	BaseController
}

func (c *VotingController) Get() {
	c.TplName = "voting.html"
}

func (c *VotingController) GetVotingImages() {
	resp, err := http.Get(fmt.Sprintf("%s/images/search?limit=10", CAT_API_BASE))
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	var images []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&images); err != nil {
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	} else {
		c.Data["json"] = images
	}
	c.ServeJSON()
}
