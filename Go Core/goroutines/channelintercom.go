package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRand() int {
	return rand.Intn(100)
}
func getSum(result chan string) {
	n1 := getRand()
	n2 := getRand()
	sum := n1 + n2
	result <- fmt.Sprintf("The sum of two random number %d and %d is=%d", n1, n2, sum)
	time.Sleep(time.Millisecond * 500)
}

func getResult(result chan string) {
	fmt.Println(<-result)
	time.Sleep(time.Millisecond * 500)
}
func main() {
	result := make(chan string)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		go getSum(result)
		go getResult(result)
	}
	close(result)
}
