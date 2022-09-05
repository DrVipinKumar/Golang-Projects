package main

import (
	"enc/student"
	"fmt"
)

func main() {
	s := student.Student{Rollno: 1}
	s.SetName("Rahul Kumar")
	s.SetCourse("B.Tech")
	fmt.Println("Student Name=", s.GetName())
	fmt.Println("Student Course=", s.GetCourse())
	fmt.Println("Student Roll No=", s.GetRollno())
	s1 := student.Student{Rollno: 2}
	s1.SetName("Ravi Kumar")
	s1.SetCourse("M.Tech")
	fmt.Println("Student Name=", s1.GetName())
	fmt.Println("Student Course=", s1.GetCourse())
	fmt.Println("Student Roll No=", s1.GetRollno())
}
