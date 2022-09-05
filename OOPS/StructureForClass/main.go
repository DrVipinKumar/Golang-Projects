package main

import (
	"oops/student"
)

func main() {
	s := student.New("Dr. Vipin Kuamr", "Ph.D", 23)
	s.GetStudentInfo()
	var s1 = student.New("Dr. Sachin Kumar", "Ph.D", 34)
	s1.GetStudentInfo()
}
