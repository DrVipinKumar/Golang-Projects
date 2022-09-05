package main

import "oops/student"

func main() {
	s := student.New("Dr. Vipin Kumar", "Ph.D", 1)
	s.GetStudentInfo()
	var s1 = student.New("Dr. Sachin Kumar", "Ph.D", 4)
	s1.GetStudentInfo()

}
