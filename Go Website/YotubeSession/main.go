package main

import (
	"net/http"
	"text/template"

	"github.com/icza/session"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("template/*.html"))
}
func indexHandle(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}
func loginHandle(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "Vipin" && password == "Teotia" {
		sess := session.NewSessionOptions(&session.SessOptions{
			CAttrs: map[string]interface{}{"username": username},
		})
		session.Add(sess, w)
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	} else {
		data := map[string]interface{}{
			"error": "Invalid Username and Password",
		}
		temp.ExecuteTemplate(w, "index.html", data)
	}

}
func welcomeHandle(w http.ResponseWriter, r *http.Request) {
	sess := session.Get(r)
	username := sess.CAttr("username")
	data := map[string]interface{}{
		"username": username,
	}
	temp.ExecuteTemplate(w, "welcome.html", data)
}
func logoutHandle(w http.ResponseWriter, r *http.Request) {
	sess := session.Get(r)
	if sess != nil {
		session.Remove(sess, w)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func main() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/login", loginHandle)
	http.HandleFunc("/welcome", welcomeHandle)
	http.HandleFunc("/logout", logoutHandle)
	http.ListenAndServe(":9999", nil)
}
