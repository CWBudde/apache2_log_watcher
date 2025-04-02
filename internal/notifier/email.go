package notifier

import (
	"fmt"
	"net/smtp"
)

// SendEmail sends an email using the given config.
func SendEmail(cfg Config, body string) error {
	from := cfg.EmailFrom
	to := cfg.EmailTo
	pass := cfg.SMTPPass
	smtpHost := cfg.SMTPServer
	smtpPort := cfg.SMTPPort
	username := cfg.SMTPUser

	msg := []byte("Subject: Visitor Alert\r\n" +
		"To: " + to + "\r\n" +
		"From: " + from + "\r\n" +
		"\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", username, pass, smtpHost)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpHost, smtpPort),
		auth,
		from,
		[]string{to},
		msg)

	if err != nil {
		return fmt.Errorf("email send error: %w", err)
	}

	return nil
}