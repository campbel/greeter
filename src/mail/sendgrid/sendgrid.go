package sendgrid

import (
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

var username string
var password string

func init() {
	username = os.Getenv("SENDGRID_USERNAME")
	password = os.Getenv("SENDGRID_PASSWORD")
}

// Provider for send grid
type Provider struct {
}

// NewProvider creates a send grid provider
func NewProvider() *Provider {
	return new(Provider)
}

// Send the email with the send grid provider
func (p Provider) Send(to, from, subject, message string) error {
	sg := sendgrid.NewSendGridClient(username, password)
	msg := sendgrid.NewMail()
	msg.AddTo(to)
	msg.SetSubject(subject)
	msg.SetHTML(message)
	msg.SetFrom(from)
	err := sg.Send(msg)
	if err != nil {
		log.Println("Error sending email with SendGrid,", err)
		return err
	}
	log.Println("Succeeded sending email with SendGrid")
	return nil
}
