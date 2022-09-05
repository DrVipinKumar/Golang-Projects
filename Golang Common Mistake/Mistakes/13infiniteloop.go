package main

import "fmt"

func main() {
	var b int
	for b = 250; b < 256; b++ {
		fmt.Printf("%d %c", b, b)
	}
}
