package models

import (
	"be-cinevo/utils"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Movie struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Overview     string    `json:"overview"`
	VoteAverage  int       `json:"vote_average"`
	PosterPath   string    `json:"poster_path"`
	BackdropPath string    `json:"backdrop_path"`
	ReleaseDate  time.Time `json:"release_date"`
	Runtime      int       `json:"runtime"`
	Popularity   int       `json:"popularity"`
	AdminID      int       `json:"admin_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Genres       []string  `json:"genres"`
	Directors    []string  `json:"directors"`
	Casts        []string  `json:"casts"`
}

func FindAllMovies(param string) ([]Movie, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}

	var query string
	if param == "upcoming" {
		query = `
				SELECT m.id, m.title, m.overview, m.vote_average, m.poster_path, m.backdrop_path, m.release_date, m.runtime, m.popularity, m.admin_id, m.created_at, m.updated_at, 
               COALESCE(array_agg(DISTINCT g.name) FILTER (WHERE g.name IS NOT NULL), '{}') AS genres,
               COALESCE(array_agg(DISTINCT d.name) FILTER (WHERE d.name IS NOT NULL), '{}') AS directors,
               COALESCE(array_agg(DISTINCT c.name) FILTER (WHERE c.name IS NOT NULL), '{}') AS casts
        FROM movies m
        LEFT JOIN movie_genres mg ON m.id = mg.movie_id
        LEFT JOIN genres g ON mg.genre_id = g.id
        LEFT JOIN movie_directors md ON m.id = md.movie_id
        LEFT JOIN directors d ON md.director_id = d.id
        LEFT JOIN movie_casts mc ON m.id = mc.movie_id
        LEFT JOIN casts c ON mc.cast_id = c.id
        WHERE m.release_date > CURRENT_DATE
            GROUP BY m.id
            ORDER BY m.release_date ASC
				`

	} else if param == "showing" {
		query = `
        SELECT m.id, m.title, m.overview, m.vote_average, m.poster_path, m.backdrop_path, m.release_date, m.runtime, m.popularity, m.admin_id, m.created_at, m.updated_at, 
               COALESCE(array_agg(DISTINCT g.name) FILTER (WHERE g.name IS NOT NULL), '{}') AS genres,
               COALESCE(array_agg(DISTINCT d.name) FILTER (WHERE d.name IS NOT NULL), '{}') AS directors,
               COALESCE(array_agg(DISTINCT c.name) FILTER (WHERE c.name IS NOT NULL), '{}') AS casts
        FROM movies m
        LEFT JOIN movie_genres mg ON m.id = mg.movie_id
        LEFT JOIN genres g ON mg.genre_id = g.id
        LEFT JOIN movie_directors md ON m.id = md.movie_id
        LEFT JOIN directors d ON md.director_id = d.id
        LEFT JOIN movie_casts mc ON m.id = mc.movie_id
        LEFT JOIN casts c ON mc.cast_id = c.id
        GROUP BY m.id
        ORDER BY m.release_date DESC`

	} else {

		query = `
				SELECT m.id, m.title, m.overview, m.vote_average, m.poster_path, m.backdrop_path, m.release_date, m.runtime, m.popularity, m.admin_id, m.created_at, m.updated_at, 
               COALESCE(array_agg(DISTINCT g.name) FILTER (WHERE g.name IS NOT NULL), '{}') AS genres,
               COALESCE(array_agg(DISTINCT d.name) FILTER (WHERE d.name IS NOT NULL), '{}') AS directors,
               COALESCE(array_agg(DISTINCT c.name) FILTER (WHERE c.name IS NOT NULL), '{}') AS casts
        FROM movies m
        LEFT JOIN movie_genres mg ON m.id = mg.movie_id
        LEFT JOIN genres g ON mg.genre_id = g.id
        LEFT JOIN movie_directors md ON m.id = md.movie_id
        LEFT JOIN directors d ON md.director_id = d.id
        LEFT JOIN movie_casts mc ON m.id = mc.movie_id
        LEFT JOIN casts c ON mc.cast_id = c.id
        GROUP BY m.id
        ORDER BY m.created_at DESC
				`
	}

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	movies, err := pgx.CollectRows[Movie](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}

	return movies, nil
}
