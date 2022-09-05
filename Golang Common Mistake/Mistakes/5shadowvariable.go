package main

import "fmt"

func main() {
	num := 0 //shadow variable
	num++
	if true {
		num = 1 //shadow variable
		num++
		fmt.Println(num)
	}
	fmt.Println(num)
}
