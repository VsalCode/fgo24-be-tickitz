package models;

import (
	"be-cinevo/dto"
	"be-cinevo/utils"
	"context"
	"github.com/jackc/pgx/v5"
)

func CreateNewMovie(req dto.MovieRequest) (*Movie, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	query := `INSERT INTO movies (title, overview, vote_average, poster_path, backdrop_path, release_date, runtime, admin_id)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var movieID int
	err = tx.QueryRow(context.Background(), query,
		req.Title, req.Overview, req.VoteAverage, req.PosterPath, req.BackdropPath, req.ReleaseDate, req.Runtime, 1,
	).Scan(&movieID)
	if err != nil {
		return nil, err
	}

	for _, genreID := range req.Genres {
		query = `INSERT INTO movie_genres (movie_id, genre_id) VALUES ($1, $2)`
		_, err = tx.Exec(context.Background(), query, movieID, genreID)
		if err != nil {
			return nil, err
		}
	}

	for _, directorID := range req.Directors {
		query = `INSERT INTO movie_directors (movie_id, director_id) VALUES ($1, $2)`
		_, err = tx.Exec(context.Background(), query, movieID, directorID)
		if err != nil {
			return nil, err
		}
	}

	for _, castID := range req.Casts {
		query = `INSERT INTO movie_casts (movie_id, cast_id) VALUES ($1, $2)`
		_, err = tx.Exec(context.Background(), query, movieID, castID)
		if err != nil {
			return nil, err
		}
	}

	rows, err := tx.Query(context.Background(),
		`SELECT id, title, overview, vote_average, poster_path, backdrop_path, release_date, runtime, popularity, admin_id, created_at, updated_at
         FROM movies WHERE id = $1`, movieID)
	if err != nil {
		return nil, err
	}

	movie, err := pgx.CollectOneRow[Movie](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	return &movie, nil
}
