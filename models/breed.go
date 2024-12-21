package models

type Breed struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CatImage struct {
	URL string `json:"url"`
}
