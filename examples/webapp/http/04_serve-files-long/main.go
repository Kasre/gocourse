package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", staticHandler)

	http.ListenAndServe(":8080", nil)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	path := string(r.URL.Path[1:])

	log.Printf("Requesting file: %s \n", path)

	fileData, err := ioutil.ReadFile(path)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintf(w, http.StatusText(500))

		return
	}

	var contentType string

	switch {
	case strings.HasSuffix(path, ".html"):
		contentType = "text/html"
	case strings.HasSuffix(path, ".js"):
		contentType = "application/javascript"
	case strings.HasSuffix(path, ".css"):
		contentType = "text/css"
	default:
		contentType = "text/plain"
	}

	log.Printf("Suggested content type is: %s \n", contentType)

	w.Header().Add("Content-Type", contentType)

	w.Write(fileData)
}
