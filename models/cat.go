package models

type Breed struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Origin            string `json:"origin"`
	Description       string `json:"description"`
	WikipediaURL      string `json:"wikipedia_url"`
	TemperamentString string `json:"temperament"`
}

type CatImage struct {
	ID     string  `json:"id"`
	URL    string  `json:"url"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Breeds []Breed `json:"breeds"`
}
