package common

import (
	"fmt"
	"net/smtp"
)

type EmailClient struct {
	Account  string
	Password string
	Host     string
	Port     string
}

func NewEmailClient(account, password, host, port string) *EmailClient {
	return &EmailClient{
		Account:  account,
		Password: password,
		Host:     host,
		Port:     port,
	}
}

func NewGmailClient(account, password string) *EmailClient {
	return NewEmailClient(
		account,
		password,
		"smtp.gmail.com",
		"587",
	)
}

func (self *EmailClient) Address() string {
	return fmt.Sprintf("%s:%s", self.Host, self.Port)
}

func (self *EmailClient) SendMessage(message []byte, recipients ...string) error {
	// Authentication.
	auth := smtp.PlainAuth("", self.Account, self.Password, self.Host)
	// Sending email.
	return smtp.SendMail(self.Address(), auth, self.Account, recipients, message)
}
