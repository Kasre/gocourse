package main

import (
	"net/http"
	"log"
	"time"
	"fmt"
)

type myInt int

func (m *myInt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 5);

	w.Write([]byte("Finished after 5 seconds"))
}

func main() {
	timeout := time.Second * 6
	http.Handle("/404", http.NotFoundHandler())

	http.Handle("/", http.TimeoutHandler(new (myInt), timeout, fmt.Sprintf("Timedout after %v", timeout)))

	http.HandleFunc("/foo", indexHandler)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Some one is being greeted!")

	w.Write([]byte("Hello, world!"))
}
