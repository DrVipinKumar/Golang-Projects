package main

import "fmt"

func main() {
	a := [2]int{1, 1}
	for _, x := range a[:] {
		fmt.Println("x=", x)
		a[1] = 8
	}
	fmt.Println("a=", a)
}
