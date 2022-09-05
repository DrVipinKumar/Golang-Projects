package main

import "fmt"

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	sum := func(num1, num2 int) int {
		return num1 + num2
	}
	fmt.Println("Sum of two number is=", sum(23, 56))
	count := counter()
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	count1 := counter()
	fmt.Println(count1())
	fmt.Println(count1())
}
