package main

import (
	"fmt"
	"time"
)

func message1(ch1 chan string) {
	for i := 0; i < 5; i++ {
		ch1 <- fmt.Sprintf("Message 1 value %d", i)
		time.Sleep(time.Second)
	}
}
func message2(ch2 chan string) {
	for i := 0; i < 5; i++ {
		ch2 <- fmt.Sprintf("Message 2 value %d", i)
		time.Sleep(time.Second)
	}
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go message1(ch1)
	go message2(ch2)
	for i := 0; i < 10; i++ {
		select {
		case r := <-ch1:
			fmt.Println(r)
		case r := <-ch2:
			fmt.Println(r)
		}
	}

}
