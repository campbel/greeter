package greeter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const failureMessage = "Failed to send the greeting, sorry :("

// Sender defines the functionality expected from an email sender
type Sender interface {
	Send(email, message string) error
}

type greeting struct {
	FirstName string
	LastName  string
	Email     string
}

// Handler is an http handler for sending greetings
type Handler struct {
	Sender Sender
}

// NewHandler returns a new greeter http.Handler
func NewHandler(sender Sender) Handler {
	return Handler{
		Sender: sender,
	}
}

// ServeHTTP handles http requests for the handler
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("POST - %s\n", r.URL.Path)

	var g greeting
	readBody(r, &g)

	if g.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message, err := getGreetingMessage(g)
	if err != nil {
		http.Error(w, failureMessage, http.StatusInternalServerError)
		return
	}

	err = h.Sender.Send(g.Email, message)
	if err != nil {
		http.Error(w, failureMessage, http.StatusInternalServerError)
		return
	}
}

func readBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}
