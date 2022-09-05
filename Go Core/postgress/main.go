package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func connectPostgresDB() *sql.DB {
	connstring := "host=localhost port=5432 user=postgres password=123 dbname=studentinfo sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
func insertIntoProtgres(db *sql.DB, rollno int, name string, course string) {
	_, err := db.Exec("insert into student(rollno, name, course) values($1, $2, $3)", rollno, name, course)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Record Inserted")
	}
}
func updateData(db *sql.DB, rollno int, name string) {
	_, err := db.Exec("update student set name=$1 where rollno=$2", name, rollno)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Record Updated")
	}

}
func deleteData(db *sql.DB, rollno int) {
	_, err := db.Exec("delete from student where rollno=$1", rollno)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Record Deleted")
	}
}
func getInput() string {
	var data string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		data = scanner.Text()
	}
	return data
}
func main() {
	var choice string
	var rollno int
	var name string
	var course string
	db := connectPostgresDB()
	for {
		fmt.Println("\nEnter your choice")
		fmt.Println("1. Insert data in Prosgres DB")
		fmt.Println("2. Read data from Prosgres DB")
		fmt.Println("3. Update data in Prosgres DB")
		fmt.Println("4. Delete data from Prosgres DB")
		fmt.Println("5. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case "1":

			fmt.Println("Enter rollno")
			roll := getInput()
			rollno, _ = strconv.Atoi(roll)
			fmt.Println("Enter name")
			name = getInput()
			fmt.Println("Enter course")
			course = getInput()
			insertIntoProtgres(db, rollno, name, course)
		case "2":
			rows, err := db.Query("select * from student")
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Rollno\tName\tCourse")
				fmt.Println("--------------------------")
				for rows.Next() {
					rows.Scan(&rollno, &name, &course)
					fmt.Printf("%d\t%s\t%s", rollno, name, course)
					fmt.Println()
				}

			}

		case "3":
			fmt.Println("Enter roll which you want to update")
			roll := getInput()
			rollno, _ = strconv.Atoi(roll)
			fmt.Println("Enter name to update")
			name = getInput()
			updateData(db, rollno, name)
		case "4":
			fmt.Println("Enter roll to delete")
			roll := getInput()
			rollno, _ = strconv.Atoi(roll)
			deleteData(db, rollno)
		case "5":
			os.Exit(0)
		}
	}
}
