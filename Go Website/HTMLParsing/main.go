package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("template/*.html"))
}
func handleFunc(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome <h1>to Go Web</h1> Page")
	temp.ExecuteTemplate(w, "index.html", nil)
}
func handleAbout(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome <h1>to Go Web</h1> Page")
	temp.ExecuteTemplate(w, "about.html", nil)
}
func handleContact(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome <h1>to Go Web</h1> Page")
	temp.ExecuteTemplate(w, "contact.html", nil)
}
func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fs))
	//http.Handle("/assets/", http.StripPrefix("/assets", fs))
	r.HandleFunc("/", handleFunc)
	r.HandleFunc("/about", handleAbout)
	r.HandleFunc("/contact", handleContact)
	http.ListenAndServe(":9999", r)
}
