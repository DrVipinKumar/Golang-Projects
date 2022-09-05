package main

import "fmt"

type calculator struct {
	num1, num2 int
}
type Calculation interface {
	sum() int
	mul() int
}

func (c calculator) sum() int {
	return c.num1 + c.num2
}
func (c calculator) mul() int {
	return c.num1 * c.num2
}
func useCalculation(cc Calculation) {
	fmt.Println("Sum of two number is=", cc.sum())
	fmt.Println("Multiply of two number is=", cc.mul())
}
func main() {
	c := calculator{23, 45}
	useCalculation(c)
}
