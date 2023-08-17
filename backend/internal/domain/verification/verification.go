package domain

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/k61b/okul/config"
)

type Verification struct {
	ID        int
	email     string
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewVerification(email string, token string, expiresAt time.Time) *Verification {
	return &Verification{
		email:     email,
		Token:     token,
		ExpiresAt: expiresAt,
	}
}

func (v *Verification) GenerateVerificationToken() (string, error) {
	cfg, err := config.LoadConfig("dev")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	secret := cfg.Utils.JWT_Secret
	tokenDuration := cfg.Utils.JWT_TokenDuration

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": v.email,
		"exp":   time.Now().Add(time.Hour * time.Duration(tokenDuration)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return tokenString, nil
}
