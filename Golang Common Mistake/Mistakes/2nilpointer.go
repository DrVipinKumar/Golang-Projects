package main

import "fmt"

type Rectangle struct {
	l, b int
}

func (r *Rectangle) Area() int {
	return r.l * r.b
}
func main() {
	var r *Rectangle
	//r := &Rectangle{l: 10, b: 20}
	//r := &Rectangle{}
	//r := new(Rectangle)
	fmt.Println(r.Area())
}
