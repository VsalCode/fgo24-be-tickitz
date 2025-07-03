package models

import (
	"be-cinevo/utils"
	"context"

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

func FindUserById(id int) (User, error) {
	conn, err := utils.DBConnect()

	if err != nil {
		return User{}, err
	}

	query := `
	SELECT u.id, p.fullname, u.email, u.password, p.phone, u.roles FROM users u 
	LEFT JOIN profiles p ON u.profile_id = p.id 
	WHERE u.id = $1 AND roles = 'user'
	`

	row, err := conn.Query(
		context.Background(),
		query,
		id,
	)
	if err != nil {
		return User{}, err
	}

	user, err := pgx.CollectOneRow[User](row, pgx.RowToStructByName)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
