package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func getMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/studentinfo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type studentInfo struct {
	Sid    string
	Name   string
	Course string
}

func getStudents(w http.ResponseWriter, r *http.Request) {

	db = getMySQLDB()
	defer db.Close()
	ss := []studentInfo{}
	s := studentInfo{}
	rows, err := db.Query("select * from student")
	if err != nil {
		fmt.Fprintf(w, ""+err.Error())
	} else {
		for rows.Next() {
			rows.Scan(&s.Sid, &s.Name, &s.Course)
			ss = append(ss, s)
		}
		json.NewEncoder(w).Encode(ss)

	}

}
func addStudents(w http.ResponseWriter, r *http.Request) {
	db = getMySQLDB()
	defer db.Close()
	s := studentInfo{}
	json.NewDecoder(r.Body).Decode(&s)
	sid, _ := strconv.Atoi(s.Sid)
	result, err := db.Exec("insert into student(sid,name,course) values(?,?,?)", sid, s.Name, s.Course)
	if err != nil {
		fmt.Fprintf(w, ""+err.Error())
	} else {
		_, err := result.LastInsertId()
		if err != nil {
			json.NewEncoder(w).Encode("error:Record not inserted")
		} else {
			json.NewEncoder(w).Encode(s)
		}

	}

}
func updateStudents(w http.ResponseWriter, r *http.Request) {
	db = getMySQLDB()
	defer db.Close()
	s := studentInfo{}
	json.NewDecoder(r.Body).Decode(&s)
	sid, _ := strconv.Atoi(s.Sid)
	result, err := db.Exec("update student set name=?,course=? where sid=?", s.Name, s.Course, sid)
	if err != nil {
		fmt.Fprintf(w, ""+err.Error())
	} else {
		_, err := result.RowsAffected()
		if err != nil {
			json.NewEncoder(w).Encode("error:Record not updated")
		} else {
			json.NewEncoder(w).Encode(s)
		}

	}

}
func deleteStudents(w http.ResponseWriter, r *http.Request) {
	db = getMySQLDB()
	defer db.Close()
	vars := mux.Vars(r)
	sid, _ := strconv.Atoi(vars["sid"])
	result, err := db.Exec("delete from student where sid=?", sid)
	if err != nil {
		fmt.Fprintf(w, ""+err.Error())
	} else {
		_, err := result.RowsAffected()
		if err != nil {
			json.NewEncoder(w).Encode("{error:Record not deleted}")
		} else {
			json.NewEncoder(w).Encode("{result:Record deleted}")
		}

	}

}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/students", addStudents).Methods("POST")
	r.HandleFunc("/students/{sid}", updateStudents).Methods("PUT")
	r.HandleFunc("/students/{sid}", deleteStudents).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
