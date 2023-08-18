package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/go-gomail/gomail"
	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/config"
	UpdateUserEmailVerificationStatus "github.com/k61b/okul/internal/application/userservice"
	verficationService "github.com/k61b/okul/internal/application/verificationservice"
	userDomain "github.com/k61b/okul/internal/domain/user"
	verificationDomain "github.com/k61b/okul/internal/domain/verification"
)

type EmailHandlers struct {
	dialer *gomail.Dialer
}

func NewEmailHandler(dialer *gomail.Dialer) *EmailHandlers {
	return &EmailHandlers{dialer: dialer}
}

func (h *EmailHandlers) SendVerificationEmailHandler(to string, token string) error {
	cfg, err := config.LoadConfig("dev")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	verificationURL := cfg.Server.Base_Url + "/verification/verify-email?token=" + token

	m := gomail.NewMessage()
	m.SetHeader("From", cfg.Email.User)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Verify your email")

	verificationLink := fmt.Sprintf("%s?token=%s", verificationURL, token)
	m.SetBody("text/html", fmt.Sprintf("Click <a href=\"%s\">here</a> to verify your email.", verificationLink))

	d := gomail.NewDialer(cfg.Email.Host, cfg.Email.Port, cfg.Email.User, cfg.Email.Password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func (h *EmailHandlers) SendVerificationEmailAndStoreTokenHandler(c *fiber.Ctx) error {
	cfg, err := config.LoadConfig("dev")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	token := c.Cookies("token")

	email, err := userDomain.GetEmailFromToken(token)
	if err != nil {
		return err
	}

	verificationToken, err := verificationDomain.GenerateVerificationToken(email)
	if err != nil {
		return err
	}

	if err := h.SendVerificationEmailHandler(email, verificationToken); err != nil {
		return err
	}

	verificationService := verficationService.NewVerificationService(nil)
	tokenDuration := cfg.Utils.JWT_TokenDuration
	expiresAt := time.Now().Add(time.Hour * time.Duration(tokenDuration)).Unix()

	if err := verificationService.Create(email, verificationToken, expiresAt); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "success"})
}

func (h *EmailHandlers) VerifyEmailHandler(c *fiber.Ctx) error {
	token := c.Query("token")

	verificationService := verficationService.NewVerificationService(nil)
	email, err := verificationService.GetEmailFromToken(token)
	if err != nil {
		return err
	}

	if err := userDomain.VerifyEmail(email, token); err != nil {
		return err
	}

	if err := verificationService.Delete(token); err != nil {
		return err
	}

	updateUserEmailVerificationStatus := UpdateUserEmailVerificationStatus.NewUserService(nil)
	if err := updateUserEmailVerificationStatus.UpdateUserEmailVerificationStatus(email); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "success"})
}
