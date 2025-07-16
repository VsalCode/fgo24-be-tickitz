package dto

import (
	"time"
)
type MovieResponse struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Overview     string    `json:"overview"`
	VoteAverage  int       `json:"vote_average"`
	PosterPath   string    `json:"poster_path"`
	BackdropPath string    `json:"backdrop_path"`
	ReleaseDate  string    `json:"release_date"`
	Runtime      int       `json:"runtime"`
	Popularity   int       `json:"popularity"`
	AdminID      int       `json:"admin_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Genres       []string  `json:"genres"`
	Directors    []string  `json:"directors"`
	Casts        []string  `json:"casts"`
}
