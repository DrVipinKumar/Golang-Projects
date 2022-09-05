package main

import "fmt"

func Odd(n int) bool {
	//return n%2 == 1
	//return n%2 != 0
	return n&1 == 1

}
func main() {
	fmt.Println(Odd(-4))
}
