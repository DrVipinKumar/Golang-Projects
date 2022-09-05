package main

import "fmt"

func main() {
	//variable is a name of memory location that hold value at runtime
	var num1 int = 20
	fmt.Println(num1)
	fmt.Println(&num1)
	var pnum1 = &num1
	fmt.Println(pnum1)
	fmt.Println(&pnum1)
	fmt.Println(*pnum1)
	var p *int
	var num2 = 45
	p = &num2
	fmt.Println(*p)
	fmt.Println(num2)
}
