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
var db *sql.DB

func getMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/studentinfo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func init() {
	tmpl = template.Must(template.ParseFiles("crudForm.html"))
}

type studentInfo struct {
	Sid    string
	Name   string
	Course string
}

func crudHandler(w http.ResponseWriter, r *http.Request) {
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
	if r.FormValue("submit") == "Insert" {
		sid, _ := strconv.Atoi(student.Sid)
		_, err := db.Exec("insert into student (sid,name,course) values(?,?,?)", sid, student.Name, student.Course)
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

	} else if r.FormValue("submit") == "Read" {
		data := []string{}
		rows, err := db.Query("select * from student")
		if err != nil {
			tmpl.Execute(w, struct {
				Success bool
				Message string
			}{Success: true, Message: err.Error()})
		} else {
			s := studentInfo{}
			data = append(data, "<table border=1>")
			data = append(data, "<tr><th>Student Id</th><th>Student Name</td><th>Student Course</th></tr>")
			for rows.Next() {
				rows.Scan(&s.Sid, &s.Name, &s.Course)
				data = append(data, fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td>", s.Sid, s.Name, s.Course))
			}
			data = append(data, "</table>")
			tmpl.Execute(w, struct {
				Success bool
				Message string
			}{Success: true, Message: strings.Trim(fmt.Sprint(data), "[]")})
		}

	} else if r.FormValue("submit") == "Update" {
		sid, _ := strconv.Atoi(student.Sid)
		result, err := db.Exec("update student set name=?, course=? where sid=?", student.Name, student.Course, sid)
		if err != nil {
			tmpl.Execute(w, struct {
				Success bool
				Message string
			}{Success: true, Message: err.Error()})
		} else {
			_, err := result.RowsAffected()
			if err != nil {
				tmpl.Execute(w, struct {
					Success bool
					Message string
				}{Success: true, Message: "Record not updated"})
			} else {
				tmpl.Execute(w, struct {
					Success bool
					Message string
				}{Success: true, Message: "Record updated"})
			}
		}

	} else if r.FormValue("submit") == "Delete" {
		sid, _ := strconv.Atoi(student.Sid)
		result, err := db.Exec("delete from student where sid=?", sid)
		if err != nil {
			tmpl.Execute(w, struct {
				Success bool
				Message string
			}{Success: true, Message: err.Error()})
		} else {
			_, err := result.RowsAffected()
			if err != nil {
				tmpl.Execute(w, struct {
					Success bool
					Message string
				}{Success: true, Message: "Record not deleted"})
			} else {
				tmpl.Execute(w, struct {
					Success bool
					Message string
				}{Success: true, Message: "Record deleted"})
			}
		}

	}

}
func main() {
	http.HandleFunc("/", crudHandler)
	http.ListenAndServe(":8080", nil)
}
