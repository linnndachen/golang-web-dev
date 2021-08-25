package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/me/", name)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func bar(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bar ran")
}

func name(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("hw.gohtml")
	if err != nil {
		log.Fatalln("error parsing tempalte", err)
	}

	err = tpl.ExecuteTemplate(w, "hw.gohtml", "Linda")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
