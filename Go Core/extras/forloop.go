package main

import "fmt"

func main() {
	for i := 1; i <= 8; i++ {
		if i == 4 || i == 6 {
			continue
		}
		fmt.Println(i)
	}
}
