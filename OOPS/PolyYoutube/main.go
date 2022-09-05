package main

import "fmt"

type Calculation interface {
	calculate() int
	displayResult()
}

func (s sum) calculate() int {
	return s.num1 + s.num2
}
func (m multiply) calculate() int {
	return m.num1 * m.num2
}

func (s sum) displayResult() {
	fmt.Printf("\nSum =%d", s.calculate())
}
func (m multiply) displayResult() {
	fmt.Printf("\nMultiply =%d", m.calculate())
}

type sum struct {
	num1 int
	num2 int
}
type multiply struct {
	num1 int
	num2 int
}

func main() {
	s := sum{23, 45}
	m := multiply{56, 78}
	cc := []Calculation{s, m}
	for _, c := range cc {
		c.displayResult()
	}
}
