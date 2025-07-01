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
	Roles    string `json:"roles"`
}

func GetNewUser(req dto.RegisterRequest) (User, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return User{}, err
	}

	tempName := strings.Split(req.Email, "@")
	fullname := tempName[0]

	var profileID int
	err = conn.QueryRow(
		context.Background(),
		`INSERT INTO profiles (fullname, phone) VALUES ($1, $2) RETURNING id`,
		fullname, "",
	).Scan(&profileID)
	if err != nil {
		return User{}, err
	}

	var userId int
	err = conn.QueryRow(
		context.Background(),
		`
				INSERT INTO users (email, password, roles, profile_id) VALUES ($1, $2, $3, $4) 
				RETURNING id
				`,
		req.Email, req.Password, "user", profileID,
	).Scan(&userId)

	if err != nil {
		return User{}, err
	}

	rows, err := conn.Query(
		context.Background(),
		`
		SELECT u.id, p.fullname, u.email, u.password, p.phone, u.roles
    FROM users u
    JOIN profiles p ON u.profile_id = p.id
    WHERE u.id = $1
		`,
		userId,
	)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow[User](rows, pgx.RowToStructByName)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func ValidateLogin(req dto.LoginRequest) (int, string, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return 0, "", err
	}

	var userId int
	var roles string
	err = conn.QueryRow(
		context.Background(),
		`
		SELECT id, roles FROM users WHERE email = $1 AND password = $2
		
		`,
		req.Email, req.Password,
	).Scan(&userId, &roles)

	if err != nil {
		return 0, "", err
	}

	return userId, roles, nil
}