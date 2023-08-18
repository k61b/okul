package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/go-gomail/gomail"
	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/config"
	verficationService "github.com/k61b/okul/internal/application/verificationservice"
	userDomain "github.com/k61b/okul/internal/domain/user"
	verificationDomain "github.com/k61b/okul/internal/domain/verification"
)

type EmailHandler struct {
	dialer *gomail.Dialer
}

func NewEmailHandler(dialer *gomail.Dialer) *EmailHandler {
	return &EmailHandler{dialer: dialer}
}

func (h *EmailHandler) SendVerificationEmailHandler(to string, token string) error {
	cfg, err := config.LoadConfig("dev")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	verificationURL := cfg.Server.Base_Url + "/verify-email?token=" + token

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

func (h *EmailHandler) SendVerificationEmailAndStoreTokenHandler(c *fiber.Ctx) error {
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
	expiresAt := fmt.Sprintf("%d", time.Now().Add(time.Hour*time.Duration(tokenDuration)).Unix())

	if err := verificationService.Create(email, verificationToken, expiresAt); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "success"})
}
