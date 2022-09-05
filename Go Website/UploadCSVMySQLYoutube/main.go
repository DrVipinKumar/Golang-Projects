package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func getMySQL() *sql.DB {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/studentinfo?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

var temp *template.Template
var db *sql.DB

func init() {
	temp = template.Must(template.ParseGlob("template/*.html"))
}
func homeHandle(w http.ResponseWriter, r *http.Request) {
	check := false
	temp.ExecuteTemplate(w, "index.html", nil)
	if r.Method == http.MethodPost {
		file, handle, err := r.FormFile("filename")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		fp := csv.NewReader(file)
		data, err := fp.ReadAll()
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		filenew, err := os.OpenFile("assets/"+handle.Filename, os.O_CREATE|os.O_RDWR, os.ModePerm)
		f, err := filenew.Stat()
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		defer filenew.Close()
		filecsvnew := csv.NewWriter(filenew)
		if f.Size() == 0 {
			filecsvnew.Write([]string{"SID", "NAME", "COURSE"})
		}

		filedata := "<table>"
		for _, line := range data {
			db = getMySQL()
			sid, _ := strconv.Atoi(line[0])
			_, err := db.Exec("insert into student(sid,name,course) values(?,?,?)", sid, line[1], line[2])
			if err != nil {
				log.Fatal(err.Error())
			} else {
				check = true
			}

			filecsvnew.Write([]string{line[0], line[1], line[2]})
			filedata = filedata + "<tr>"
			filedata = filedata + "<td>" + line[0] + "</td>" + "<td>" + line[1] + "</td>" + "<td>" + line[2] + "</td>"
			filedata = filedata + "</tr>"
		}
		filecsvnew.Flush()
		filedata = filedata + "</table>"
		fmt.Fprintf(w, filedata)
		fmt.Fprintf(w, "File uploaded")
		if check {
			fmt.Fprintf(w, "Data also stored in MySQL")
		} else {
			fmt.Fprintf(w, "Check MySQL connection")
		}
	}
}
func main() {
	http.HandleFunc("/", homeHandle)
	http.ListenAndServe(":9999", nil)
}
