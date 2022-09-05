package main

import "fmt"

func main() {
	values := []int{2, 3, 5, 7}
	for v := range values {
		fmt.Println(v, "=", v)

	}

}
