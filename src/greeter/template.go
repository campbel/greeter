package greeter

import (
	"bytes"
	"log"
	"text/template"
)

var greetingTemplate = template.Must(template.New("greeting").Parse(greetingTemplateHTML))

const greetingTemplateHTML = `
<h1>Greetings {{.FirstName}} {{.LastName}}!</h1>
<p>I hope you have a great time on this special day!</p>
`

func getGreetingMessage(g greeting) (string, error) {
	var doc bytes.Buffer
	err := greetingTemplate.ExecuteTemplate(&doc, "greeting", g)
	if err != nil {
		log.Println(err)
		return "", err
	}
	message := doc.String()
	return message, nil
}
