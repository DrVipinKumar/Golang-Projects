package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type svalue struct {
	sid    string
	name   string
	course string
}

func getMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/studentinfo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func main() {
	var db = getMySQLDB()
	courses := []svalue{}
	file, err := os.Open("course.csv")
	if err != nil {
		log.Fatal(err)
	}
	df := csv.NewReader(file)
	data, _ := df.ReadAll()
	for _, value := range data {
		courses = append(courses, svalue{sid: value[0], name: value[1], course: value[2]})
	}
	for i := 1; i < len(courses); i++ {
		sid, _ := strconv.Atoi(courses[i].sid)
		_, err := db.Exec("insert into student(sid,name,course) values(?,?,?)", sid, courses[i].name, courses[i].course)
		if err != nil {
			fmt.Printf("" + err.Error())
		}
	}
	fmt.Println("All record inserted")
}
