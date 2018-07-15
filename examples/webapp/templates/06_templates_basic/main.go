package main

import (
	"net/http"
	"html/template"
	"log"
	"os"
)

const doc = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Hello World!</title>
</head>
<body>
	<h1>Hello, World! I'm a template</h1>
</body>
</html>
`

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("FooBar").Parse(doc)

	if err != nil {
		log.Fatal(err)

		os.Exit(1)

		return
	}

	w.Header().Add("Content-Type", "text/html")
	tpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8080", nil)
}