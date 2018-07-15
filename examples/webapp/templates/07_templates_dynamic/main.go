package main

import (
	"net/http"
	"html/template"
	"log"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Books []string
}

const doc = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Hello World!</title>
</head>
<body>
	<h1>Hello there {{.Name}}!</h1>
	<div>Your age is {{.Age}}</div>
	<br />

	<dir>
		Your favorite books are:
		<ul>
		{{range .Books}}
		<li>{{.}}</li>
		{{end}}
	</div>
</body>
</html>
`

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("MyNewTemplate").Parse(doc)

	if err != nil {
		log.Fatal(err)

		os.Exit(1)

		return
	}

	person := Person{
		"Bob",
		55,
		[]string{
			"The hitchhikers guide to the galaxy",
			"I.T",
		},
	}

	tpl.Execute(w, person)
}

func main() {
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8080", nil)
}
