package models

import (
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
	Phone string `json:"phone"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func GetNewUser(req RegisterRequest) (User, error) {
	conn, err := utils.DBConnect()

	if err != nil {
		return User{}, err
	}

	tempName := strings.Split(req.Email, "@")

	rows, err := conn.Query(
		context.Background(),
		`
		INSERT INTO users (fullname, email, password, phone) VALUES ($1, $2, $3, $4) RETURNING *
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