package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/icza/session"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseGlob("template/*.html"))
}
func welcomeHandle(w http.ResponseWriter, r *http.Request) {
	sess := session.Get(r)
	if sess == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		userName := sess.CAttr("username")
		data := map[string]interface{}{
			"username": userName,
		}
		templ.ExecuteTemplate(w, "welcome.html", data)
	}

}
func logoutHandle(w http.ResponseWriter, r *http.Request) {
	sess := session.Get(r)
	if sess == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		session.Remove(sess, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}
func loginHandle(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "index.html", nil)
}
func loginCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		templ.Execute(w, nil)
		return
	}
	if r.FormValue("username") == "vipin" && r.FormValue("password") == "teotia" {
		//sess := session.NewSession()
		username := strings.ToUpper(r.FormValue("username"))
		sess := session.NewSessionOptions(&session.SessOptions{
			CAttrs: map[string]interface{}{"username": username},
		})
		session.Add(sess, w)
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	} else {
		data := map[string]interface{}{
			"err": "Invalid Username or Password",
		}
		templ.ExecuteTemplate(w, "index.html", data)
	}

}
func main() {

	http.HandleFunc("/", loginHandle)
	http.HandleFunc("/login", loginCheck)
	http.HandleFunc("/welcome", welcomeHandle)
	http.HandleFunc("/logout", logoutHandle)
	http.ListenAndServe(":9999", nil)
}
