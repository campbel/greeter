package amazon

import (
	"log"
	"net/smtp"
	"os"
)

const server = "email-smtp.us-west-2.amazonaws.com"
const port = "587"

var username string
var password string

func init() {
	username = os.Getenv("AMAZON_USERNAME")
	password = os.Getenv("AMAZON_PASSWORD")
}

// Provider for amazon
type Provider struct {
}

// NewProvider creates an amazon provider
func NewProvider() *Provider {
	return new(Provider)
}

// Send the email with the amazon provider
func (p Provider) Send(to, from, subject, message string) error {
	auth := smtp.PlainAuth("", username, password, server)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "Subject: " + subject + "\n"
	msg := []byte(body + mime + message)
	serverPort := server + ":" + port
	err := smtp.SendMail(serverPort, auth, from, []string{to}, msg)
	if err != nil {
		log.Printf("Error sending email with Amazon, %s", err)
		return err
	}
	log.Println("Succeeded sending email with Amazon")
	return nil
}
