package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Id      int
	Message string
}

func main() {

	data := map[string][]Todo{
		"Todos": {
			Todo{
				Id:      1,
				Message: "Learn Go",
			},
			Todo{
				Id:      2,
				Message: "Learn HTMX",
			},
		},
	}

	todosHandler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, data)
	}

	http.HandleFunc("/", todosHandler)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
