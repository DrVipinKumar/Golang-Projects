package student

import "fmt"

type student struct {
	Name   string
	Course string
	Rollno int
}

func New(name string, course string, rollno int) student {
	s := student{Name: name, Course: course, Rollno: rollno}
	return s
}

func (s *student) GetStudentInfo() {
	fmt.Printf("\nStudent Name =%s, Course=%s, and Rollno=%d", s.Name, s.Course, s.Rollno)
}
