package main

import (
	"net/http"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("StudentInfo.html"))
}

type studentInfo struct {
	Sid    string
	Name   string
	Course string
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	student := studentInfo{
		Sid:    r.FormValue("sid"),
		Name:   r.FormValue("name"),
		Course: r.FormValue("course"),
	}
	tmpl.Execute(w, struct {
		Success bool
		Student studentInfo
	}{true, student})
}
func main() {
	http.HandleFunc("/", StudentHandler)
	http.ListenAndServe(":8080", nil)
}
