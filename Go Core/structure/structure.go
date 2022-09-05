package main

import "fmt"

type person struct {
	name          string
	qualification string
}

func getPerson() *person {
	p := person{"Ankit", "Ph.D"}
	return &p
}
func main() {
	p2 := getPerson()
	fmt.Println(p2.name)
	p1 := person{"Dr. Vipin Kumar", "Ph.D"}
	fmt.Println(p1)
	fmt.Println("Name=", p1.name)
	fmt.Println("Qualification=", p1.qualification)
	fmt.Println(person{"Anil Kumar", "M.Tech"})
	sp1 := &p1
	fmt.Println(sp1.name)
	sp1.name = "Dr. Sachin Kumar"
	fmt.Println(p1.name)
}
