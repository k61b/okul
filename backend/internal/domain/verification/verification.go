package domain

import (
	"log"
	"time"

	"github.com/go-gomail/gomail"
	"github.com/golang-jwt/jwt"
	"github.com/k61b/okul/config"
)

type Verification struct {
	ID        int
	Email     string
	Token     string
	ExpiresAt time.Time
}

func NewVerification(email, token string, expiresAt time.Time) *Verification {
	return &Verification{
		Email:     email,
		Token:     token,
		ExpiresAt: expiresAt,
	}
}

func GenerateVerificationToken(email string, expiresAt time.Time) (string, error) {
	cfg, err := config.LoadConfig("dev")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	secret := cfg.Utils.JWT_Secret

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   expiresAt,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func SendVerificationEmail(email, token string) error {
	cfg, err := config.LoadConfig("dev")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	verificationURL := cfg.Server.Base_Url + "/user/verify-email?token=" + token

	m := gomail.NewMessage()
	m.SetHeader("From", cfg.Email.User)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Verify your email")
	m.SetBody("text/html", "Click <a href=\""+verificationURL+"\">here</a> to verify your email.")

	d := gomail.NewDialer(cfg.Email.Host, cfg.Email.Port, cfg.Email.User, cfg.Email.Password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
