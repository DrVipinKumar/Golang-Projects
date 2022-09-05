package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

type student struct {
	sid    string
	name   string
	course string
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
	var students = []student{}
	file, err := os.OpenFile("students.csv", os.O_APPEND|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err != nil {
		panic(err)
	}
forlabel:
	for {
		fmt.Println("Do you want to write data in CSV file")
		fmt.Scanln(&choice)
		switch choice {
		case "yes":
			var sid, name, course string
			fmt.Println("Enter student id")
			sid = getInput()
			fmt.Println("Enter student name")
			name = getInput()
			fmt.Println("Enter student course")
			course = getInput()
			students = append(students, student{sid, name, course})
		case "no":
			break forlabel
		}
	}
	df := csv.NewWriter(file)
	f, _ := file.Stat()
	if f.Size() == 0 {
		df.Write([]string{"sid", "name", "course"})
	}
	for _, s := range students {
		err := df.Write([]string{s.sid, s.name, s.course})
		if err != nil {
			panic(err)
		}
	}
	df.Flush()
	fmt.Println("Record inserted in CSV file")
}
