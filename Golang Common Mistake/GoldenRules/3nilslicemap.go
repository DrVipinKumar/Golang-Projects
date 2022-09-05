package main

import "fmt"

func main() {
	var s []int
	s = append(s, 1)
	s = append(s, 2)
	fmt.Println(s)
	fmt.Println(cap(s))
	//var m map[string]int
	var m = make(map[string]int)
	m["one"] = 1 //error
	m["two"] = 2
	fmt.Println(cap(m)) //not to use cap with map
	fmt.Println(m)
}
