package models

import (
	"be-cinevo/dto"
	"be-cinevo/utils"
	"context"
	"math/rand"
	"strings"
	"time"
	"github.com/redis/go-redis/v9"
	"github.com/jackc/pgx/v5"
	"fmt"
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

func SendVerificationCode(email string) error {
    conn, err := utils.DBConnect()
    if err != nil {
        return err
    }

    var exists bool
    err = conn.QueryRow(
        context.Background(),
        `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)`,
        email,
    ).Scan(&exists)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("email not found")
    }

    rand.Seed(time.Now().UnixNano())
    digits := "0123456789"
    code := make([]byte, 6)
    for i := range code {
        code[i] = digits[rand.Intn(len(digits))]
    }
    verificationCode := string(code)

    ctx := context.Background()
    key := "otp:" + email
    err = utils.RedisClient.Set(ctx, key, verificationCode, 2*time.Minute).Err()
    if err != nil {
        return fmt.Errorf("failed to save OTP to Redis")
    }

    // _, err = conn.Exec(
    //     context.Background(),
    //     `UPDATE users SET verification_code=$1 WHERE email=$2`,
    //     verificationCode, email,
    // )
    // if err != nil {
    //     return fmt.Errorf("failed to save OTP to DB: %v", err)
    // }

    return nil
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