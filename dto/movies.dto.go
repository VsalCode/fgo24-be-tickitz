package dto;

type MovieRequest struct {
    Title        string   `json:"title" binding:"required"`
    Overview     string   `json:"overview"`
    VoteAverage  int      `json:"vote_average" binding:"required,min=0,max=10"`
    PosterPath   string   `json:"poster_path"`
    BackdropPath string   `json:"backdrop_path"`
    ReleaseDate  string   `json:"release_date"`
    Runtime      int      `json:"runtime" binding:"required,min=1"`
    Genres       []int    `json:"genres" binding:"required"`
    Directors    []int    `json:"directors" binding:"required"`
    Casts        []int    `json:"casts" binding:"required"`
}
