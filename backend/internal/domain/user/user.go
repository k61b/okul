package domain

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt"

	"github.com/k61b/okul/config"
)

// User represents a user entity
type User struct {
	ID              int
	Email           string
	IsEmailVerified bool
	Password        string
	Name            string
	Surname         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// NewUser creates a new User instance
func NewUser(email, password, name, surname string) *User {
	return &User{
		Email:           email,
		IsEmailVerified: false,
		Password:        password,
		Name:            name,
		Surname:         surname,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// CheckPassword checks if the password is correct
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Create JWT token
func GenerateJWTToken(email string) (string, error) {
	cfg, err := config.LoadConfig("dev")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	secret := cfg.Utils.JWT_Secret

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Parse JWT token
func ParseToken(token string) (*jwt.Token, error) {
	cfg, err := config.LoadConfig("dev")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	secret := cfg.Utils.JWT_Secret

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
}

// Create Cookie
func GenerateCookie(token string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}
}

// Get Email from JWT token
func GetEmailFromToken(token string) (string, error) {
	parsedToken, err := ParseToken(token)
	if err != nil {
		return "", err
	}

	claims := parsedToken.Claims.(jwt.MapClaims)

	return claims["email"].(string), nil
}

// Verify Email
func VerifyEmail(email, token string) error {
	parsedToken, err := ParseToken(token)
	if err != nil {
		return err
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	fmt.Println(claims["email"], email)
	if claims["email"] != email {
		return fiber.ErrUnauthorized
	}

	return nil
}
