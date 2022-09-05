package main

import "fmt"

type calculator struct {
	num1, num2 int
}

func (c *calculator) sum() int {
	return c.num1 + c.num2
}
func (c calculator) mul() int {
	return c.num1 * c.num2
}
func main() {
	cc := calculator{23, 56}
	fmt.Println("Sum of two number", cc.sum())
	fmt.Println("Multiply of two number", cc.mul())
}
