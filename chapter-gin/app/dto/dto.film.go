package dto

import "time"

type Movies []Movie
type Movie struct {
	FilmID          int       `json:"film_id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	ReleaseYear     int       `json:"release_year"`
	LanguageID      int       `json:"language_id"`
	RentalDuration  int       `json:"rental_duration"`
	RentalRate      float64   `json:"rental_rate"`
	Length          int       `json:"length"`
	ReplacementCost float64   `json:"replacement_cost"`
	Rating          string    `json:"rating"`
	LastUpdate      time.Time `json:"last_update"`
	SpecialFeatures []string  `json:"special_features"`
	Fulltext        string    `json:"fulltext"`
}
