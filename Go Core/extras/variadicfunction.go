package main

import (
	"fmt"
	"time"
)

func sum(nums ...int) int {
	s := 0
	for _, num := range nums {
		s = s + num
	}
	return s
}
func main() {
	fmt.Println(time.Now().Format("02-01-2006"))
	result := sum(12, 45)
	fmt.Println("Sum of two number is=", result)
	result = sum(12, 45, 34, 56, 45, 678)
	fmt.Println("Sum of numbers are=", result)

}
