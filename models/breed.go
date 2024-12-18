// models/breed.go
package models

type Breed struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Origin      string `json:"origin"`
}

type CatImage struct {
	ID     string  `json:"id"`
	URL    string  `json:"url"`
	Breeds []Breed `json:"breeds"`
}
