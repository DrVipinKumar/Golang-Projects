package main

import "fmt"

func main() {
	var m map[string]int
	//m := make(map[string]int)
	m["r"] = 23
	fmt.Println(m["r"])
}
