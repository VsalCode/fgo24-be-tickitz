package models

import (
	"be-cinevo/utils"
	"context"
	// "time"
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
