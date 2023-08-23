package domain

import (
	"log"
	"time"

	"github.com/go-gomail/gomail"
	"github.com/golang-jwt/jwt"
	"github.com/k61b/okul/config"
)

type Verification struct {
	ID               int
	VerificationType string
	Email            string
	Token            string
	ExpiresAt        time.Time
}

func NewVerification(verificationType, email, token string, expiresAt time.Time) *Verification {
	return &Verification{
		VerificationType: verificationType,
		Email:            email,
		Token:            token,
		ExpiresAt:        expiresAt,
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
	m.SetBody("text/html", `
		<html>
			<head>
				<style>
					body {
						font-family: Arial, sans-serif;
					}
					.container {
						background-color: #f9f9f9;
						padding: 20px;
						border-radius: 10px;
						text-align: center;
					}
					.header {
						background-color: #ffa500;
						color: white;
						padding: 10px;
						border-radius: 10px 10px 0 0;
					}
					.content {
						padding: 20px;
					}
					.button {
						display: inline-block;
						background-color: #4caf50;
						color: white;
						padding: 10px 20px;
						text-decoration: none;
						border-radius: 5px;
					}
					.footer {
						background-color: #4caf50;
						color: white;
						padding: 10px;
						border-radius: 0 0 10px 10px;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<div class="header">
						<h2>Welcome to `+cfg.Utils.Project_Name+`!</h2>
					</div>
					<div class="content">
						<p>Hello,</p>
						<p>Thank you for registering with our application. To complete your registration and verify your email, please click the button below:</p>
						<p><a class="button" href="`+verificationURL+`">Verify Email</a></p>
						<p>If you didn't sign up for an account, you can safely ignore this email.</p>
						<p>Regards,<br>The `+cfg.Utils.Project_Name+` Team</p>
					</div>
					<div class="footer">
						<p>© 2023 `+cfg.Utils.Project_Name+`. All rights reserved.</p>
					</div>
				</div>
			</body>
		</html>
	`)

	d := gomail.NewDialer(cfg.Email.Host, cfg.Email.Port, cfg.Email.User, cfg.Email.Password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func SendVerificationEmailForPassword(email, token string) error {
	cfg, err := config.LoadConfig("dev")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	resetPasswordURL := cfg.Server.Base_Url + "/user/reset-password?token=" + token

	m := gomail.NewMessage()
	m.SetHeader("From", cfg.Email.User)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Reset your password")
	m.SetBody("text/html", `
		<html>
			<head>
				<style>
					body {
						font-family: Arial, sans-serif;
					}
					.container {
						background-color: #f9f9f9;
						padding: 20px;
						border-radius: 10px;
						text-align: center;
					}
					.header {
						background-color: #ffa500;
						color: white;
						padding: 10px;
						border-radius: 10px 10px 0 0;
					}
					.content {
						padding: 20px;
					}
					.button {
						display: inline-block;
						background-color: #4caf50;
						color: white;
						padding: 10px 20px;
						text-decoration: none;
						border-radius: 5px;
					}
					.footer {
						background-color: #4caf50;
						color: white;
						padding: 10px;
						border-radius: 0 0 10px 10px;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<div class="header">
						<h2>Reset your password</h2>
					</div>
					<div class="content">
						<p>Hello,</p>
						<p>We received a request to reset your password. If you didn't make the request, just ignore this email. Otherwise, you can reset your password using this link:</p>
						<p><a class="button" href="`+resetPasswordURL+`">Reset Password</a></p>
						<p>If you didn't sign up for an account, you can safely ignore this email.</p>
						<p>Regards,<br>The `+cfg.Utils.Project_Name+` Team</p>
					</div>
					<div class="footer">
						<p>© 2023 `+cfg.Utils.Project_Name+`. All rights reserved.</p>
					</div>
				</div>
			</body>
		</html>
	`)

	d := gomail.NewDialer(cfg.Email.Host, cfg.Email.Port, cfg.Email.User, cfg.Email.Password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
