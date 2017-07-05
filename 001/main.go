package main

import ("net/http"
		"io"
		"html/template")

var tpl *template.Template

func init() {
	tpl = template.Must(template, ParseGlob("templates/*.gohtml"))
}


func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index (w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `)
}