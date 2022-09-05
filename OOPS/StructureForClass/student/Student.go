package student

import (
	"fmt"
)

type studentInfo struct {
	Name   string
	Course string
	Rollno int
}

func New(Name string, Course string, Rollno int) studentInfo {
	s := studentInfo{Name, Course, Rollno}
	return s
}
func (s *studentInfo) GetStudentInfo() {
	fmt.Printf("\nStudent Name =%s, Course = %s and Roll No=%d", s.Name, s.Course, s.Rollno)
}
