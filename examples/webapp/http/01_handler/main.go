package main

import "net/http"

type MyHandler struct {
	greeting string
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(m.greeting))
}

func main()  {
	handler := MyHandler{"Hello, WORLD!!!"}

	http.Handle("/", &handler)

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello there mister foo"))
	})

	http.ListenAndServe(":8080", nil)
}