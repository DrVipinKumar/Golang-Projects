package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbm *sql.DB

func connectDB() {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/studentinfo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MySQl Connected....")
	dbm = db
}
func createTable() {
	query := `Create table student (
		     sid int auto_increment,
			 name text not null, 
			 course text not null,
			 create_at datetime,
			 primary key(sid)
			 );`
	_, err := dbm.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Student table created...")
}
func insertInfo() {
	name := "Dr. Vipin Kumar"
	course := "Ph.D"
	created_at := time.Now()
	result, err := dbm.Exec(`insert into student (name,course,create_at) values(?,?,?)`,
		name, course, created_at)
	if err != nil {
		log.Fatal(err)
	} else {
		value, _ := result.LastInsertId()
		fmt.Println("Record insert for Student ID=", value)
	}
}
func updateInfo() {
	sid := 1
	name := "Dr. Sachin Kumar"
	_, err := dbm.Exec(`update student set name=? where sid=?`,
		name, sid)
	if err != nil {
		log.Fatal(err)
	} else {

		fmt.Println("Record Updated")
	}
}
func getAll() {
	type student struct {
		sid       int
		name      string
		course    string
		create_at time.Time
	}
	var s student
	rows, err := dbm.Query(`select * from student`)
	if err != nil {
		log.Fatal(err)
	} else {
		for rows.Next() {
			rows.Scan(&s.sid, &s.name, &s.course, &s.create_at)
			fmt.Println("Student Id=", s.sid)
			fmt.Println("Student Name=", s.name)
			fmt.Println("Student Course=", s.course)
			fmt.Println("Student ID Create at time=", s.create_at)
		}
	}
}
func deleteStudent() {
	_, err := dbm.Exec(`delete from student where sid=?`, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Record delete for sid", 2)
}
func main() {
	connectDB()
	//createTable()
	//insertInfo()
	//updateInfo()
	deleteStudent()
	getAll()
}
