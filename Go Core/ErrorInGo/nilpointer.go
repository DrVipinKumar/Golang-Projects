package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
func main() {
	var p *Point
	//var p *Point = new(Point)
	//var p Point
	fmt.Println(p.Abs())
}
