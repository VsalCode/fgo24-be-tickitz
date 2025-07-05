package models

import (
	"be-cinevo/utils"
	"context"

	"time"

	"github.com/jackc/pgx/v5"
)

type Transactions struct {
	Fullname        string   `json:"customer_fullname"`
	Email           string   `json:"customer_email"`
	Phone           string   `json:"customer_phone"`
	Amount          float64  `json:"amount"`
	Cinema          string   `json:"cinema"`
	Location        string   `json:"location"`
	Time            string   `json:"show_time"`
	Date            string   `json:"show_date"`
	MovieID         int      `json:"movie_id"`
	PaymentMethodID int      `json:"payment_method_id"`
	Seats           []string `json:"seat"`
}

type HistoryTransactions struct {
	Title  string    `json:"title"`
	Genre  []string  `json:"genre"`
	Date   time.Time `json:"date"`
	Time   time.Time `json:"time"`
	Seat   []string  `json:"seat"`
	Cinema string    `json:"cinema"`
}

func HandleBookingTicket(id int, req Transactions) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	var transactionID int
	query := `
        INSERT INTO transactions (
            customer_fullname, customer_email, customer_phone, amount,
            cinema, location, show_time, show_date,
            user_id, movie_id, payment_method_id
        )
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
        RETURNING id
    `
	err = tx.QueryRow(
		context.Background(),
		query,
		req.Fullname, req.Email, req.Phone, req.Amount,
		req.Cinema, req.Location, req.Time, req.Date,
		id, req.MovieID, req.PaymentMethodID,
	).Scan(&transactionID)
	if err != nil {
		return err
	}

	for _, seat := range req.Seats {
		_, err := tx.Exec(
			context.Background(),
			`INSERT INTO transaction_details (seat, transaction_id) VALUES ($1, $2)`,
			seat, transactionID,
		)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}

func FindHistoryByUserId(id int) ([]HistoryTransactions, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}

	query := `
  SELECT 
    m.title AS title,
    COALESCE(array_agg(DISTINCT g.name)) AS genre,
    t.show_date AS date,
    t.show_time AS time,
    COALESCE(array_agg(DISTINCT td.seat)) AS seat,
    t.cinema
  FROM transactions t
  LEFT JOIN payment_method pm ON pm.id = t.payment_method_id
  LEFT JOIN transaction_details td ON td.transaction_id = t.id 
  LEFT JOIN movies m ON m.id = t.movie_id
  LEFT JOIN movie_genres mg ON m.id = mg.movie_id
  LEFT JOIN genres g ON mg.genre_id = g.id
  WHERE t.user_id = $1
  GROUP BY t.id, m.title, t.show_date, t.show_time, t.cinema
  ORDER BY t.created_at DESC
    `

	rows, err := conn.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}

	historys, err := pgx.CollectRows[HistoryTransactions](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}

	return historys, nil
}
