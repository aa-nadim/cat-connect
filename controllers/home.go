package controllers

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type HomeController struct {
	beego.Controller
}

type HomeData struct {
	Breeds    interface{} `json:"breeds"`
	Favorites interface{} `json:"favorites"`
	Error     string      `json:"error,omitempty"`
}

func (c *HomeController) Get() {
	var wg sync.WaitGroup
	var data HomeData
	apiKey, _ := beego.AppConfig.String("APIKey")

	// Create channels for data and errors
	breedsChan := make(chan interface{}, 1)
	favoritesChan := make(chan interface{}, 1)
	errorChan := make(chan error, 2)

	// Fetch breeds
	wg.Add(1)
	go func() {
		defer wg.Done()
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/breeds", nil)
		if err != nil {
			errorChan <- err
			return
		}

		req.Header.Add("x-api-key", apiKey)
		resp, err := client.Do(req)
		if err != nil {
			errorChan <- err
			return
		}
		defer resp.Body.Close()

		var breeds interface{}
		if err := json.NewDecoder(resp.Body).Decode(&breeds); err != nil {
			errorChan <- err
			return
		}

		breedsChan <- breeds
	}()

	// Fetch favorites
	wg.Add(1)
	go func() {
		defer wg.Done()
		subID := "user-123" // You might want to make this dynamic
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/favourites?sub_id="+subID, nil)
		if err != nil {
			errorChan <- err
			return
		}

		req.Header.Add("x-api-key", apiKey)
		resp, err := client.Do(req)
		if err != nil {
			errorChan <- err
			return
		}
		defer resp.Body.Close()

		var favorites interface{}
		if err := json.NewDecoder(resp.Body).Decode(&favorites); err != nil {
			errorChan <- err
			return
		}

		favoritesChan <- favorites
	}()

	// Wait for all goroutines to complete
	go func() {
		wg.Wait()
		close(breedsChan)
		close(favoritesChan)
		close(errorChan)
	}()

	// Collect results
	select {
	case breeds := <-breedsChan:
		data.Breeds = breeds
	case <-time.After(5 * time.Second):
		data.Error = "Timeout while fetching breeds"
	}

	select {
	case favorites := <-favoritesChan:
		data.Favorites = favorites
	case <-time.After(5 * time.Second):
		data.Error = "Timeout while fetching favorites"
	}

	// Check for any errors
	select {
	case err := <-errorChan:
		data.Error = err.Error()
	default:
	}

	c.Data["APIKey"] = apiKey
	c.Data["InitialData"] = data
	c.TplName = "home.html"
}
