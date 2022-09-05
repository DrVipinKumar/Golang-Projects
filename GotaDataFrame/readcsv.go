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

var db *sql.DB

type sdata struct {
	sno    string
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
	db = getMySQLDB()
	var svalue = []sdata{}
	file, err := os.Open("course.csv")
	if err != nil {
		log.Fatal(err)
	}
	df := csv.NewReader(file)
	data, _ := df.ReadAll()
	for _, value := range data {

		svalue = append(svalue, sdata{sno: value[0], name: value[1], course: value[2]})

	}
	for i := 1; i < len(svalue); i++ {
		sid, _ := strconv.Atoi(svalue[i].sno)
		db.Exec("insert into student(sid,name,course) values(?,?,?)", sid, svalue[i].name, svalue[i].course)
	}
	fmt.Println("Inserted All")

}
