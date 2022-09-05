package main

import "fmt"

func main() {
	var m map[string]float64
	m["pi"] = 3.1416
	fmt.Println(m)
	//solution
	m := make(map[string]float64)
}
