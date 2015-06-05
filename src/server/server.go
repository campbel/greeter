package main

import (
	"greeter"
	"log"
	"mail"
	"mail/amazon"
	"mail/sendgrid"
	"net/http"
)

func main() {
	handler := greeter.NewHandler(
		mail.NewSender(
			sendgrid.NewProvider(),
			amazon.NewProvider(),
		),
	)

	http.Handle("/greeting", handler)

	log.Println("Starting web server on port 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
