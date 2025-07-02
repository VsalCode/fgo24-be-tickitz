package utils

import (
    "fmt"
		"time"
		"math/rand"
    "net/smtp"
		"os"
)

func SendEmailOTP(toEmail, otp string) error {

    smtpHost := "smtp.gmail.com"
    smtpPort := "587"
    senderEmail := "cinevoticket@gmail.com"      
    senderPassword := os.Getenv("APP_PASSWORD")    

    auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

    subject := "Subject: Your OTP Code\r\n"
    body := fmt.Sprintf("Your OTP code is: %s", otp)
    msg := []byte(subject + "\r\n" + body)

    err := smtp.SendMail(
        smtpHost+":"+smtpPort,
        auth,
        senderEmail,
        []string{toEmail},
        msg,
    )
    return err
}

func GenerateOTP(length int) string {
	rand.Seed(time.Now().UnixNano())
    digits := "0123456789"
    code := make([]byte, length)
    for i := range code {
        code[i] = digits[rand.Intn(len(digits))]
    }
    return string(code)
}