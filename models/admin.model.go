package models;

import (
	"be-cinevo/dto"
	"be-cinevo/utils"
	"context"
)

func CreateNewMovie(req dto.MovieRequest) (error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	query := `INSERT INTO movies (title, overview, vote_average, poster_path, backdrop_path, release_date, runtime, admin_id)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var movieID int
	err = tx.QueryRow(context.Background(), query,
		req.Title, req.Overview, req.VoteAverage, req.PosterPath, req.BackdropPath, req.ReleaseDate, req.Runtime, 1,
	).Scan(&movieID)
	if err != nil {
		return err
	}

	for _, genreID := range req.Genres {
		query = `INSERT INTO movie_genres (movie_id, genre_id) VALUES ($1, $2)`
		_, err = tx.Exec(context.Background(), query, movieID, genreID)
		if err != nil {
			return err
		}
	}

	for _, directorID := range req.Directors {
		query = `INSERT INTO movie_directors (movie_id, director_id) VALUES ($1, $2)`
		_, err = tx.Exec(context.Background(), query, movieID, directorID)
		if err != nil {
			return err
		}
	}

	for _, castID := range req.Casts {
		query = `INSERT INTO movie_casts (movie_id, cast_id) VALUES ($1, $2)`
		_, err = tx.Exec(context.Background(), query, movieID, castID)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}

func DeleteMovieById(id int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}

	defer conn.Close()

	query := `DELETE FROM movies WHERE id = $1`

	_, err = conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}