package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	for _, n := range a {
		fmt.Println(n)
	}
}
