package services

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

// EmailService handles sending emails via SMTP (copied from Rolston project)
type EmailService struct {
	smtpHost     string
	smtpPort     string
	smtpUser     string
	smtpPassword string
	fromEmail    string
}

// NewEmailService creates a new email service
func NewEmailService() *EmailService {
	return &EmailService{
		smtpHost:     getEnv("SMTP_HOST", ""),
		smtpPort:     getEnv("SMTP_PORT", "587"),
		smtpUser:     getEnv("SMTP_USER", ""),
		smtpPassword: getEnv("SMTP_PASS", ""),
		fromEmail:    getEnv("SMTP_FROM", getEnv("SMTP_USER", "noreply@example.com")),
	}
}

// IsConfigured checks if SMTP is properly configured
func (es *EmailService) IsConfigured() bool {
	return es.smtpHost != ""
}

// SendNotificationEmail sends a notification email to the configured recipients
func (es *EmailService) SendNotificationEmail(subject, body string) error {
	if !es.IsConfigured() {
		return fmt.Errorf("SMTP not configured")
	}

	notificationEmails := getEnv("NOTIFICATION_EMAILS", "")
	if notificationEmails == "" {
		return fmt.Errorf("NOTIFICATION_EMAILS not configured")
	}

	recipients := strings.Split(notificationEmails, ",")
	for i, email := range recipients {
		recipients[i] = strings.TrimSpace(email)
	}

	message := fmt.Sprintf("From: %s\r\n", es.fromEmail)
	message += fmt.Sprintf("To: %s\r\n", strings.Join(recipients, ", "))
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += "MIME-Version: 1.0\r\n"
	message += "Content-Type: text/html; charset=UTF-8\r\n"
	message += "\r\n"
	message += body

	var auth smtp.Auth
	if es.smtpUser != "" && es.smtpPassword != "" {
		auth = smtp.PlainAuth("", es.smtpUser, es.smtpPassword, es.smtpHost)
	}

	addr := fmt.Sprintf("%s:%s", es.smtpHost, es.smtpPort)

	// MailHog/local support (no TLS/auth)
	if es.smtpPort == "1025" || es.smtpHost == "mailhog" {
		client, err := smtp.Dial(addr)
		if err != nil {
			return fmt.Errorf("failed to connect to SMTP server: %w", err)
		}
		defer client.Close()

		if err := client.Mail(es.fromEmail); err != nil {
			return fmt.Errorf("failed to set sender: %w", err)
		}

		for _, recipient := range recipients {
			if err := client.Rcpt(recipient); err != nil {
				return fmt.Errorf("failed to set recipient %s: %w", recipient, err)
			}
		}

		writer, err := client.Data()
		if err != nil {
			return fmt.Errorf("failed to open data writer: %w", err)
		}

		if _, err = writer.Write([]byte(message)); err != nil {
			return fmt.Errorf("failed to write message: %w", err)
		}

		if err = writer.Close(); err != nil {
			return fmt.Errorf("failed to close data writer: %w", err)
		}

		return client.Quit()
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         es.smtpHost,
	}

	client, err := smtp.Dial(addr)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer client.Close()

	if ok, _ := client.Extension("STARTTLS"); ok {
		if err := client.StartTLS(tlsConfig); err != nil {
			return fmt.Errorf("failed to start TLS: %w", err)
		}
	}

	if auth != nil {
		if err := client.Auth(auth); err != nil {
			return fmt.Errorf("failed to authenticate: %w", err)
		}
	}

	if err := client.Mail(es.fromEmail); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}

	for _, recipient := range recipients {
		if err := client.Rcpt(recipient); err != nil {
			return fmt.Errorf("failed to set recipient %s: %w", recipient, err)
		}
	}

	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to open data writer: %w", err)
	}

	if _, err = writer.Write([]byte(message)); err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	if err = writer.Close(); err != nil {
		return fmt.Errorf("failed to close data writer: %w", err)
	}

	return client.Quit()
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
