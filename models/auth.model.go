package models

import (
	"be-cinevo/dto"
	"be-cinevo/utils"
	"context"
	"strings"
	"time"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

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

	passwordHash, err := utils.HashPassword(req.Password)
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
		req.Email, passwordHash, "user", profileID,
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
	var password string

	err = conn.QueryRow(
		context.Background(),
		`
    SELECT id, password, roles FROM users WHERE email = $1
    `,
		req.Email,
	).Scan(&userId, &password, &roles)

	if err != nil {
		return 0, "", fmt.Errorf("user not found")
	}

	exist := utils.CheckPasswordHash(req.Password, password)
	if !exist {
		return 0, "", fmt.Errorf("invalid password")
	}

	return userId, roles, nil
}

func SendVerificationCode(email string) (dto.OTP, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return dto.OTP{}, err
	}

	var exists bool
	err = conn.QueryRow(
		context.Background(),
		`SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)`,
		email,
	).Scan(&exists)
	if err != nil {
		return dto.OTP{}, err
	}
	if !exists {
		return dto.OTP{}, err
	}

	verificationCode := utils.GenerateOTP(6)

	result := dto.OTP{
		Email:            email,
		VerificationCode: verificationCode,
	}

	ctx := context.Background()
	key := "otp:" + email
	err = utils.RedisClient.Set(ctx, key, verificationCode, 2*time.Minute).Err()
	if err != nil {
		return dto.OTP{}, err
	}

	return result, nil
}

func SendNewPassword(req dto.ForgotPasswordRequest) error {
	ctx := context.Background()
	key := "otp:" + req.Email

	storedCode, err := utils.RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return err
	}
	if err != nil {
		return err
	}
	if storedCode != req.VerificationCode {
		return err
	}

	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	_, err = conn.Exec(
		context.Background(),
		`UPDATE users SET password=$1 WHERE email=$2`,
		req.NewPassword, req.Email,
	)
	if err != nil {
		return err
	}

	utils.RedisClient.Del(ctx, key)

	return nil
}
