package main

import (
	"fmt"
)

func main() {
	var x string = "" //error

	if x == "" { //error
		x = "This is string variable"
	}
	fmt.Println(x)
}
