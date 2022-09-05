package main

import "fmt"

func Calculator(num1, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}
func sum(num1 int, num2 int) int {
	return (num1 + num2)
}
func mul(num1, num2 int) int {
	return (num1 * num2)
}
func main() {
	s := sum(23, 45)
	fmt.Println("Sum of two number is=", s)
	m := mul(23, 56)
	fmt.Println("Mulitply of two number is =", m)
	sum, mul := Calculator(23, 45)
	fmt.Println("Sum and multiplication of two number is=", sum, mul)
	sum1, _ := Calculator(24, 5)
	fmt.Println("Sum of two number is=", sum1)

}
