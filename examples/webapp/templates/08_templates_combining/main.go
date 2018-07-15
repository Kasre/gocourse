package main

import (
	"net/http"
	"html/template"
)

type Person struct {
	Name  string
	Age   int
	Books []string
}

const body = `
{{template "header"}}
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
{{template "footer"}}
`

const head = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8" />
    <title>Hey there!</title>
</head>
`

const foot = `
</html>
`

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.New("MyNewTemplate")

	tpl.New("body").Parse(body)
	tpl.New("header").Parse(head)
	tpl.New("footer").Parse(foot)

	person := Person{
		"Bob",
		55,
		[]string{
			"The hitchhikers guide to the galaxy",
			"I.T",
		},
	}

	tpl.Lookup("body").Execute(w, person)
}

func main() {
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8080", nil)
}
