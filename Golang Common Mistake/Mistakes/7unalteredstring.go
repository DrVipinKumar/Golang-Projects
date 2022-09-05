package main

import "fmt"

func main() {
	//s := "Mr Vipin Classes"
	buf := []rune("Mr Vipin Classes")
	buf[0] = 'D'
	s := string(buf)
	fmt.Println(s)
	fmt.Println(buf)
}
