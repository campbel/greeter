package mail

import (
	"errors"
	"log"
)

const from = "greeter@mailinator.com"
const subject = "Greetings on your special occasion!"

// Provider defines the functionality expected from a mail provider
type Provider interface {
	Send(to, from, subject, message string) error
}

// Sender is a application specific abstraction for sending emails
type Sender struct {
	Providers []Provider
}

// Send will iterate through the senders providers until one succeeds
func (f Sender) Send(email, message string) error {
	var err error
	for _, p := range f.Providers {
		err = p.Send(email, from, subject, message)
		if err == nil {
			return nil
		}
	}
	errMsg := "Failed sending mail after trying all the providers"
	log.Println(errMsg)
	return errors.New(errMsg)
}

// NewSender creates a new mail sender out of the given providers
func NewSender(providers ...Provider) Sender {
	f := Sender{
		Providers: make([]Provider, len(providers)),
	}

	for i := range providers {
		f.Providers[i] = providers[i]
	}

	return f
}
