package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main()  {
	r := mux.NewRouter()

	r.HandleFunc("/todo/{id}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)

		writer.Write([]byte("Selecting a ToDo with ID:" + vars["id"]))
	}).Methods("GET")

	// serve todo app static folder
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.ListenAndServe(":8030", r)
}
