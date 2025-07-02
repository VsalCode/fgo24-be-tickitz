package models

type Movies struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Overview     string `json:"overview"`
	VoteAverage  string `json:"voteAverage"`
	PosterPath   string `json:"posterPath"`
	BackdropPath string `json:"backdropPath"`
	ReleaseDate  string `json:"releaseDate"`
	Runtime      int    `json:"runtime"`
	Popularity   int    `json:"popularity"`
}

func GetNowShowingMovies() {}

func GetUpComingMovies() {}
