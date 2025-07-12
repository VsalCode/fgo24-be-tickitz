package dto

import (
	"time"
)

type MovieRequest struct {
	Title        string `json:"title" binding:"required"`
	Overview     string `json:"overview"`
	VoteAverage  int    `json:"vote_average" binding:"required,min=0,max=10"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
	ReleaseDate  string `json:"release_date"`
	Runtime      int    `json:"runtime" binding:"required,min=1"`
	Genres       []int  `json:"genres" binding:"required"`
	Directors    []int  `json:"directors" binding:"required"`
	Casts        []int  `json:"casts" binding:"required"`
}

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
