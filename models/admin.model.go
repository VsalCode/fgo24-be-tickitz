package models

import (
	"be-cinevo/dto"
	"be-cinevo/utils"
	"context"
	"time"
)

type UpdatedMovie struct {
	Title        *string    `json:"title"`
	Overview     *string    `json:"overview"`
	VoteAverage  *int       `json:"vote_average"`
	PosterPath   *string    `json:"poster_path"`
	BackdropPath *string    `json:"backdrop_path"`
	ReleaseDate  *time.Time `json:"release_date"`
	Runtime      *int       `json:"runtime"`
}


func CreateNewMovie(req dto.MovieRequest) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	query := `
	INSERT INTO movies (title, overview, vote_average, poster_path, backdrop_path, release_date, runtime, admin_id)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
	`

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

func GetUpdatedMovie(id int, req UpdatedMovie) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	var (
		title        string
		overview     string
		voteAverage  int
		posterPath   string
		backdropPath string
		releaseDate  time.Time
		runtime      int
	)
	query := `
        SELECT title, overview, vote_average, poster_path, backdrop_path, release_date, runtime
        FROM movies WHERE id = $1
    `
	err = conn.QueryRow(context.Background(), query, id).Scan(
		&title, &overview, &voteAverage, &posterPath, &backdropPath, &releaseDate, &runtime,
	)
	if err != nil {
		return err
	}

	newTitle := &title
	newOverview := &overview
	newVoteAverage := &voteAverage
	newPosterPath := &posterPath
	newBackdropPath := &backdropPath
	newReleaseDate := &releaseDate
	newRuntime := &runtime

	if req.Title != nil {
		newTitle = req.Title
	}
	if req.Overview != nil {
		newOverview = req.Overview
	}
	if req.VoteAverage != nil {
		newVoteAverage = req.VoteAverage
	}
	if req.PosterPath != nil {
		newPosterPath = req.PosterPath
	}
	if req.BackdropPath != nil {
		newBackdropPath = req.BackdropPath
	}
	if req.ReleaseDate != nil {
		newReleaseDate = req.ReleaseDate
	}
	if req.Runtime != nil {
		newRuntime = req.Runtime
	}

	_, err = conn.Exec(
		context.Background(),
		`
    UPDATE movies
		SET
			title = $1,
			overview = $2,
			vote_average = $3,
			poster_path = $4,
			backdrop_path = $5,
			release_date = $6,
			runtime = $7,
			updated_at = NOW()
		WHERE id = $8
        `,
		*newTitle, *newOverview, *newVoteAverage, *newPosterPath, *newBackdropPath, *newReleaseDate, *newRuntime, id,
	)
	if err != nil {
		return err
	}

	return nil
}

func GetTicketSales(filter string) ([]map[string]interface{}, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var query string
	if filter == "genre" {
		query = `
			SELECT g.name AS genre, COUNT(t.movie_id) AS total_sales
			FROM transactions t
			JOIN movies m ON t.movie_id = m.id
			JOIN movie_genres mg ON mg.movie_id = m.id
			JOIN genres g ON g.id = mg.genre_id
			GROUP BY g.name
		`
	} else if filter == "movie" {
		query = `
			SELECT m.title AS movie_name, COUNT(t.movie_id) AS total_sales
			FROM transactions t
			JOIN movies m ON t.movie_id = m.id
			GROUP BY m.title
		`
	}

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}

	for rows.Next() {
		var movieName string
		var totalSales int
		if err := rows.Scan(&movieName, &totalSales); err != nil {
			return nil, err
		}

		result := map[string]interface{}{
			"movie_name":  movieName,
			"total_sales": totalSales,
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func GetSalesChart(filter string) ([]map[string]interface{}, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var query string
	if filter == "location" {
		query = `
			SELECT t.location AS genre, COUNT(t.id) AS total_sales
			FROM transactions t
			GROUP BY t.location
		`
	} else if filter == "cinema" {
		query = `
			SELECT t.cinema AS genre, COUNT(t.id) AS total_sales
			FROM transactions t
			GROUP BY t.cinema
		`
	}

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}

	for rows.Next() {
		var movieName string
		var totalSales int
		if err := rows.Scan(&movieName, &totalSales); err != nil {
			return nil, err
		}

		result := map[string]interface{}{
			"movie_name":  movieName,
			"total_sales": totalSales,
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
