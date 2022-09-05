package main

import (
	"fmt"
)

type student struct {
	firstName string
	lastName  string
	bio       string
}

func (a student) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type topics struct {
	title   string
	content string
	student
}

func (p topics) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Student: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

type website struct {
	topic []topics
}

func (w website) contents() {
	fmt.Println("Contents of Website\n")
	for _, v := range w.topic {
		v.details()
		fmt.Println()
	}
}

func main() {
	s1 := student{
		"Anil Kumar",
		"Sunil Kumar",
		"Ravi Kumar",
	}
	t1 := topics{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		s1,
	}
	t2 := topics{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		s1,
	}
	t3 := topics{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		s1,
	}
	w := website{
		topic: []topics{t1, t2, t3},
	}
	w.contents()
}
