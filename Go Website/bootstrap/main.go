package main

import (
	"net/http"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("template/*.html"))
}
func RunIndex(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}
func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", RunIndex)
	http.ListenAndServe(":8080", nil)
}
