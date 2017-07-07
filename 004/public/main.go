package main

import ("net/http"
		"html/template"
		)

type person struct {
	First_name string
	Last_name string
	Saying string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}


func main() {
	http.HandleFunc("/", index)
	http.Handle("/assets", http.StripPrefix("/assets", http.FileServer(http.Dir("public/"))))
	http.ListenAndServe(":8080", nil)
}

func index (w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First_name: "James",
		Last_name: "Bond",
		Saying: "Shaken, not stirred",
	}

	p2 := person{
		First_name: "Al",
		Last_name: "Pacino",
		Saying: "Cunho",
	}

	xp := []person{p1, p2}


	tpl.ExecuteTemplate(w, "index.gohtml", xp )
}