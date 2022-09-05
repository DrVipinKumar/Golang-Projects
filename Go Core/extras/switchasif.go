package main

import "fmt"

func main() {
	i := 4
	switch {
	case i > 5:
		fmt.Println(i, " greater than 5")
	case i < 5:
		fmt.Println(i, "less than 5")
	case i == 5:
		fmt.Println(i, " is equal to 5")
	default:
		fmt.Println("default value")

	}
}
