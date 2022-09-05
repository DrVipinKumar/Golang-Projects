package main

import "fmt"

func swap(num1, num2 *int) {
	temp := *num1
	*num1 = *num2
	*num2 = temp
}
func main() {
	n1 := 20
	n2 := 30
	fmt.Println(n1, n2)
	swap(&n1, &n2)
	fmt.Println(n1, n2)

}
