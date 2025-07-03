package models

import (
	"be-cinevo/dto"
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

func GetUpdatedUserInfo(id int, req dto.UpdatedUser) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	var email, password string
	var profileID int
	err = conn.QueryRow(context.Background(),
		`SELECT email, password, profile_id FROM users WHERE id = $1`, id).
		Scan(&email, &password, &profileID)
	if err != nil {
		return err
	}

	var fullname, phone string
	if profileID != 0 {
		err = conn.QueryRow(context.Background(),
			`SELECT fullname, phone FROM profiles WHERE id = $1`, profileID).
			Scan(&fullname, &phone)
		if err != nil {
			return err
		}
	}

	newEmail := email
	if req.Email != "" {
		newEmail = req.Email
	}
	newPassword := password
	if req.Password != "" {
		newPassword = req.Password
	}
	newFullname := fullname
	if req.Fullname != "" {
		newFullname = req.Fullname
	}
	newPhone := phone
	if req.Phone != "" {
		newPhone = req.Phone
	}

	_, err = conn.Exec(context.Background(),
		`UPDATE users SET email = $1, password = $2 WHERE id = $3`,
		newEmail, newPassword, id)
	if err != nil {
		return err
	}

	if profileID != 0 {
		_, err = conn.Exec(context.Background(),
			`UPDATE profiles SET fullname = $1, phone = $2 WHERE id = $3`,
			newFullname, newPhone, profileID)
		if err != nil {
			return err
		}
	}

	return nil
}
