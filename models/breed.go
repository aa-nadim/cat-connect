package models

type Breed struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Origin      string `json:"origin"`
	Description string `json:"description"`
}

type CatImage struct {
	URL string `json:"url"`
}
