package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main()  {
	r := mux.NewRouter()

	// serve todo app static folder
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.ListenAndServe(":8030", r)
}
