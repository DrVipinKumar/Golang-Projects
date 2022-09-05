package main

import "fmt"

func main() {
	const (
		Num1 = 100
		Num2 = 10
		Num3 = 1
	)
	//base 8 (octal) start with 0
	//base 10 (decimal) never start with 0
	//base 16 (hexa) start with 0x
	fmt.Println("Total =", Num1+2*Num2+2*Num3)
}
