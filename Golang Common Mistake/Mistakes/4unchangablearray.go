package main

import "fmt"

func Value(a [2]int) {
	a[0] = 8
}
func main() {
	a := [2]int{1, 2}
	fmt.Println(a)
	Value(a)
	fmt.Println(a)
}
