package main

import (
	"net/http"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", serveStatic)

	http.ListenAndServe(":8080", nil)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	if data, err := ioutil.ReadFile(path); err != nil {
		w.WriteHeader(404)

		w.Write([]byte("404 Mi Amor - " + http.StatusText(404)))
	} else {
		w.Write(data)
	}
}
