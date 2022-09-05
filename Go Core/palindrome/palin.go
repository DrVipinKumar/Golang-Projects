package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInput() string {
	var data string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		data = scanner.Text()
	}
	return data
}
func main() {
	fmt.Println("Enter string to check palindrome")
	value := getInput()
	l := len(value) - 1
	check := true
	for i := 0; i < len(value)/2; i++ {
		if value[i] != value[l] {
			check = false
		}
		l--
	}
	if check == true {
		fmt.Println("String is palindrome")
	} else {
		fmt.Println("String is not palindrome")
	}
}
