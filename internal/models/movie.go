package models

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseYear int       `json:"release_year"`
	Rating      float64   `json:"rating"`
	ImageURL    string    `json:"image_url"`
	Director    string    `json:"director"`
	Genres      []string  `json:"genres"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PipelineStats struct {
	Processed int `json:"processed"`
	Enriched  int `json:"enriched"`
	Errors    int `json:"errors"`
}

type EnrichmentData struct {
	IMDbRating  float64  `json:"imdb_rating"`
	BoxOffice   string   `json:"box_office"`
	Awards      string   `json:"awards"`
	Plot        string   `json:"plot"`
	Actors      []string `json:"actors"`
	Runtime     string   `json:"runtime"`
}
