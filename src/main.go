package main

import (
	"html/template"
	"net/http" //standard package for http connection
)

// User is struct
type User struct {
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	user := User{
		Name: "sample",
		Age:  20,
	}
	tmpl := template.Must(template.ParseFiles("src/views/index.html"))
	tmpl.Execute(responseWriter, user)
}
