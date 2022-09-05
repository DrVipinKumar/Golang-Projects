package main

import "fmt"

func main() {
	var num1 = 20
	var num2 = 30
	//var num3 = 45
	if num3 := 45; num1 > num2 && num1 > num3 {
		fmt.Println("Num1 is greater than num2 and num3")
	} else if num2 > num1 && num2 > num3 {
		fmt.Println("Num2 is greater than num1 and num3")
	} else {
		fmt.Println("Num3 is greater than num2 and num1")
	}
}
