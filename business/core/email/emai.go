package email

import (
	"context"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

type EmailSender interface {
	Send(ctx context.Context, email Email) error
}

type Core struct {
	adminAddress string
	dialer       *gomail.Dialer
}

func NewCore(adminAddress string, dialer *gomail.Dialer) *Core {

	return &Core{
		adminAddress: adminAddress,
		dialer:       dialer,
	}
}

func (c *Core) Send(ctx context.Context, email Email) error {

	m := gomail.NewMessage()

	m.SetHeader("From", c.adminAddress)
	m.SetHeader("To", email.Recipient)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/plain", email.Body)

	if err := c.dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}
