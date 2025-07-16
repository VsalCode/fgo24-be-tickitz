package models

import (
	"be-cinevo/dto"
	"be-cinevo/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Movie struct {
	ID           int       `json:"-"`
	Title        string    `json:"-"`
	Overview     string    `json:"-"`
	VoteAverage  int       `json:"-"`
	PosterPath   string    `json:"-"`
	BackdropPath string    `json:"-"`
	ReleaseDate  time.Time `json:"-"`
	Runtime      int       `json:"-"`
	Popularity   int       `json:"-"`
	AdminID      int       `json:"-"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	Genres       []string  `json:"-"`
	Directors    []string  `json:"-"`
	Casts        []string  `json:"-"`
}

func toMovieResponse(m Movie) dto.MovieResponse {
	return dto.MovieResponse{
		ID:           m.ID,
		Title:        m.Title,
		Overview:     m.Overview,
		VoteAverage:  m.VoteAverage,
		PosterPath:   m.PosterPath,
		BackdropPath: m.BackdropPath,
		ReleaseDate:  m.ReleaseDate.Format("2006-01-02"),
		Runtime:      m.Runtime,
		Popularity:   m.Popularity,
		AdminID:      m.AdminID,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		Genres:       m.Genres,
		Directors:    m.Directors,
		Casts:        m.Casts,
	}
}

type Genre struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Cast struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Director struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func HandleShowAllMovies(key string, limit int, offset int, filter string) ([]dto.MovieResponse, int, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, 0, err
	}
	defer conn.Close()

	redisKey := fmt.Sprintf("movies:%s:%d:%d:%s", key, limit, offset, filter)

	cachedMovies, err := utils.RedisClient.Get(context.Background(), redisKey).Result()
	if err == nil {
		var movies []dto.MovieResponse
		err = json.Unmarshal([]byte(cachedMovies), &movies)
		if err == nil {
			return movies, len(movies), nil
		}
	}

	query := `
	SELECT 
		m.id, m.title, m.overview, m.vote_average, m.poster_path, m.backdrop_path, 
		m.release_date, m.runtime, m.popularity, m.admin_id, m.created_at, m.updated_at,
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
	WHERE m.title ILIKE $1
	AND ($2 = '' OR EXISTS (
		SELECT 1 
		FROM movie_genres mg2
		JOIN genres g2 ON mg2.genre_id = g2.id
		WHERE mg2.movie_id = m.id
		AND g2.name ILIKE $2
	))
	GROUP BY m.id
	ORDER BY m.created_at DESC
	LIMIT $3 OFFSET $4`

	rows, err := conn.Query(context.Background(), query, "%"+key+"%", "%"+filter+"%", limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	dbMovies, err := pgx.CollectRows[Movie](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, 0, err
	}

	movies := make([]dto.MovieResponse, len(dbMovies))
	for i, m := range dbMovies {
		movies[i] = toMovieResponse(m)
	}

	countQuery := `
	SELECT COUNT(DISTINCT m.id)
	FROM movies m
	LEFT JOIN movie_genres mg ON m.id = mg.movie_id
	LEFT JOIN genres g ON mg.genre_id = g.id
	WHERE m.title ILIKE $1
	AND ($2 = '' OR g.name ILIKE $2)`

	var total int
	err = conn.QueryRow(context.Background(), countQuery, "%"+key+"%", "%"+filter+"%").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	moviesJSON, err := json.Marshal(movies)
	if err == nil {
		utils.RedisClient.Set(context.Background(), redisKey, moviesJSON, 0)
	}

	return movies, total, nil
}

func HandleNowShowingMovies() ([]dto.MovieResponse, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := `
	SELECT 
		m.id, m.title, m.overview, m.vote_average, m.poster_path, m.backdrop_path, 
		m.release_date, m.runtime, m.popularity, m.admin_id, m.created_at, m.updated_at,
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
	WHERE m.release_date <= CURRENT_DATE
	AND m.release_date >= CURRENT_DATE - INTERVAL '1 month'
	GROUP BY m.id
	ORDER BY m.release_date DESC`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dbMovies, err := pgx.CollectRows[Movie](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}

	movies := make([]dto.MovieResponse, len(dbMovies))
	for i, m := range dbMovies {
		movies[i] = toMovieResponse(m)
	}

	return movies, nil
}

func HandleUpComingMovies() ([]dto.MovieResponse, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := `
	SELECT 
		m.id, m.title, m.overview, m.vote_average, m.poster_path, m.backdrop_path, 
		m.release_date, m.runtime, m.popularity, m.admin_id, m.created_at, m.updated_at,
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
	ORDER BY m.release_date ASC`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dbMovies, err := pgx.CollectRows[Movie](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}

	movies := make([]dto.MovieResponse, len(dbMovies))
	for i, m := range dbMovies {
		movies[i] = toMovieResponse(m)
	}

	return movies, nil
}

func FindMovieById(id int) (*dto.MovieResponse, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := `
    SELECT 
        m.id, m.title, m.overview, m.vote_average, m.poster_path, m.backdrop_path, 
        m.release_date, m.runtime, m.popularity, m.admin_id, m.created_at, m.updated_at,
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
    WHERE m.id = $1
    GROUP BY m.id`

	rows, err := conn.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dbMovies, err := pgx.CollectRows[Movie](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}

	if len(dbMovies) == 0 {
		return nil, errors.New("movie not found")
	}

	movie := toMovieResponse(dbMovies[0])
	return &movie, nil
}

func FindAllGenres() ([]Genre, error) {

	conn, err := utils.DBConnect()
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}
	defer conn.Close()

	query := `SELECT id, name FROM genres ORDER BY id`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	genres, err := pgx.CollectRows(rows, pgx.RowToStructByName[Genre])
	if err != nil {
		return nil, fmt.Errorf("collect rows failed: %w", err)
	}

	return genres, nil
}

func FindAllCasts() ([]Cast, error) {

	conn, err := utils.DBConnect()
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}
	defer conn.Close()

	query := `SELECT id, name FROM casts ORDER BY id`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	casts, err := pgx.CollectRows(rows, pgx.RowToStructByName[Cast])
	if err != nil {
		return nil, fmt.Errorf("collect rows failed: %w", err)
	}

	return casts, nil
}

func FindAllDirectors() ([]Director, error) {

	conn, err := utils.DBConnect()
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}
	defer conn.Close()

	query := `SELECT id, name FROM directors ORDER BY id`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	directors, err := pgx.CollectRows(rows, pgx.RowToStructByName[Director])
	if err != nil {
		return nil, fmt.Errorf("collect rows failed: %w", err)
	}

	return directors, nil
}
