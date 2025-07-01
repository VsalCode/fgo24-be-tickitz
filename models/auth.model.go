package models

import (
	"be-cinevo/dto"
	"be-cinevo/utils"
	"context"
	"strings"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func GetNewUser(req dto.RegisterRequest) (User, error) {
	conn, err := utils.DBConnect()

	if err != nil {
		return User{}, err
	}

	tempName := strings.Split(req.Email, "@")

	rows, err := conn.Query(
		context.Background(),
		`
		INSERT INTO users (fullname, email, password, phone, roles) VALUES ($1, $2, $3, $4, user) RETURNING *
		`,
		tempName[0],
	)

	if err != nil {
		return User{}, err
	}

	user, err := pgx.CollectOneRow[User](rows, pgx.RowToStructByName)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
