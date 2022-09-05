package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var tmpl *template.Template

type studentInfo struct {
	Sid    string
	Name   string
	Course string
}

var db *sql.DB

func getMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/studentinfo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func init() {
	tmpl = template.Must(template.ParseFiles("forms.html"))
}
func formhandler(w http.ResponseWriter, r *http.Request) {
	db = getMySQLDB()
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	student := studentInfo{
		Sid:    r.FormValue("sid"),
		Name:   r.FormValue("name"),
		Course: r.FormValue("course"),
	}
	if r.FormValue("submit") == "Read" {
		rows, err := db.Query("select * from student")
		if err != nil {
			tmpl.Execute(w, struct {
				Success bool
				Message string
			}{Success: true, Message: err.Error()})
		} else {
			data := []string{}
			s := studentInfo{}
			data = append(data, "<table><tr><th>SID</th><th>Student Name</th><th>Student Course</th></tr>")
			for rows.Next() {
				rows.Scan(&s.Sid, &s.Name, &s.Course)
				data = append(data, fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td></tr>", s.Sid, s.Name, s.Course))
			}
			defer rows.Close()
			data = append(data, "<table>")
			tmpl.Execute(w, struct {
				Success bool
				Message string
			}{Success: true, Message: strings.Trim(fmt.Sprint(data), "[]")})
		}
	}
	if r.FormValue("submit") == "Insert" {
		sid, _ := strconv.Atoi(student.Sid)
		_, err := db.Exec("insert into student(sid,name,course) values(?,?,?)", sid, student.Name, student.Course)
		if err != nil {
			tmpl.Execute(w, struct {
				Success bool
				Message string
			}{Success: true, Message: err.Error()})
		} else {
			tmpl.Execute(w, struct {
				Success bool
				Message string
			}{Success: true, Message: "Record Inserted"})
		}
	}

}
func main() {

	http.HandleFunc("/", formhandler)
	http.ListenAndServe(":8080", nil)
}
