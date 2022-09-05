package main

import "fmt"

func main() {
	var data = [6]int{2, 4, 5, 7, 8, 10}
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
	var data2 = [2][3]int{{12, 34, 56}, {56, 78, 98}}
	for i := 0; i < len(data2); i++ {
		for j := 0; j < len(data2[i]); j++ {
			fmt.Print(data2[i][j], " ")
		}
		fmt.Println()
	}

}
