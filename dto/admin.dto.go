package dto

type GenreRequest struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PersonRequest struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type MovieRequest struct {
	Title        string           `json:"title"`
	Overview     string           `json:"overview"`
	VoteAverage  float64          `json:"vote_average"`
	PosterPath   string           `json:"poster_path"`
	BackdropPath string           `json:"backdrop_path"`
	ReleaseDate  string           `json:"release_date"`
	Runtime      int              `json:"runtime"`
	Genres       []GenreRequest   `json:"genres"`
	Directors    []PersonRequest  `json:"directors"`
	Casts        []PersonRequest  `json:"casts"`
}