package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144

func getInput() string {
	var data string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		data = scanner.Text()
	}
	return data
}
func DisplayFibonacci(value int) {
	n1 := 0
	n2 := 1
	fmt.Print(n1, ",", n2)
	for i := 0; i < value-2; i++ {
		t := n1 + n2
		n1 = n2
		n2 = t
		fmt.Print(",", t)
	}
	fmt.Println()
}
func main() {
	fmt.Println("Enter range for Fibonacci series")
	value := getInput()
	num, _ := strconv.Atoi(value)
	DisplayFibonacci(num)
}
