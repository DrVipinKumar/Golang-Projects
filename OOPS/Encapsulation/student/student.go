package student

type Student struct {
	name   string
	course string
	Rollno int
}

func (s *Student) GetRollno() int {
	return s.Rollno
}
func (s *Student) GetCourse() string {
	return s.course
}
func (s *Student) SetCourse(course string) {
	s.course = course
}
func (s *Student) GetName() string {
	return s.name
}
func (s *Student) SetName(name string) {
	s.name = name
}
