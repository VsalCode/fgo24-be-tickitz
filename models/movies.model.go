package models

import (
	"be-cinevo/utils"
	"context"
	"encoding/json"
	"fmt"
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

func HandleShowAllMovies(key string, limit int, offset int, filter string) ([]Movie, int, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, 0, err
	}
	defer conn.Close()

 redisKey := fmt.Sprintf("movies:%s:%d:%d:%s", key, limit, offset, filter)

	cachedMovies, err := utils.RedisClient.Get(context.Background(), redisKey).Result()
	if err == nil {
		var movies []Movie
		err = json.Unmarshal([]byte(cachedMovies), &movies)
		if err == nil {
			return movies, len(movies), nil
		}
	}

	query := `
    SELECT m.id, m.title, m.overview, m.vote_average, m.poster_path, m.backdrop_path, m.release_date, m.runtime, m.popularity, m.admin_id, m.created_at, m.updated_at, 
        COALESCE(array_agg(DISTINCT g.name)) AS genres,
        COALESCE(array_agg(DISTINCT d.name)) AS directors,
        COALESCE(array_agg(DISTINCT c.name)) AS casts
    FROM movies m
    LEFT JOIN movie_genres mg ON m.id = mg.movie_id
    LEFT JOIN genres g ON mg.genre_id = g.id
    LEFT JOIN movie_directors md ON m.id = md.movie_id
    LEFT JOIN directors d ON md.director_id = d.id
    LEFT JOIN movie_casts mc ON m.id = mc.movie_id
    LEFT JOIN casts c ON mc.cast_id = c.id
    WHERE m.title ILIKE $1
    AND ($2 = '' OR g.name ILIKE $2)
    GROUP BY m.id
    ORDER BY m.created_at DESC
    LIMIT $3 OFFSET $4
    `

	rows, err := conn.Query(context.Background(), query, "%"+key+"%", "%"+filter+"%", limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	movies, err := pgx.CollectRows[Movie](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, 0, err
	}

	countQuery := `SELECT COUNT(*) FROM movies WHERE title ILIKE $1`
	var total int
	err = conn.QueryRow(context.Background(), countQuery, "%"+key+"%").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	moviesJSON, err := json.Marshal(movies)
	if err == nil {
		utils.RedisClient.Set(context.Background(), redisKey, moviesJSON, 0) 
	}

	return movies, total, nil
}

func HandleNowShowingMovies() ([]Movie, error) {
    conn, err := utils.DBConnect()
    if err != nil {
        return nil, err
    }

    query := `
    SELECT 
        m.id, 
        m.title, 
        m.overview, 
        m.vote_average, 
        m.poster_path, 
        m.backdrop_path, 
        m.release_date, 
        m.runtime, 
        m.popularity, 
        m.admin_id, 
        m.created_at, 
        m.updated_at,
        COALESCE(NULLIF(array_agg(DISTINCT g.name), '{NULL}'), '{}') AS genres,
        COALESCE(NULLIF(array_agg(DISTINCT d.name), '{NULL}'), '{}') AS directors,
        COALESCE(NULLIF(array_agg(DISTINCT c.name), '{NULL}'), '{}') AS casts
    FROM movies m
    LEFT JOIN movie_genres mg ON m.id = mg.movie_id
    LEFT JOIN genres g ON mg.genre_id = g.id
    LEFT JOIN movie_directors md ON m.id = md.movie_id
    LEFT JOIN directors d ON md.director_id = d.id
    LEFT JOIN movie_casts mc ON m.id = mc.movie_id
    LEFT JOIN casts c ON mc.cast_id = c.id
    WHERE m.release_date <= CURRENT_DATE
      AND m.release_date >= CURRENT_DATE - INTERVAL '1 month'
    GROUP BY m.id
    ORDER BY m.release_date DESC`

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

func HandleUpComingMovies() ([]Movie, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
		query := `
		SELECT m.id, m.title, m.overview, m.vote_average, m.poster_path, m.backdrop_path, m.release_date, m.runtime, m.popularity, m.admin_id, m.created_at, m.updated_at, 
        COALESCE(array_agg(DISTINCT g.name)) AS genres,
 				COALESCE(array_agg(DISTINCT d.name)) AS directors,
        COALESCE(array_agg(DISTINCT c.name)) AS casts
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
