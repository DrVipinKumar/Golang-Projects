package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		ch1 <- "Message 1"
	}()
	select {
	case res := <-ch1:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("Message 1 timout")
	}
	ch2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message 2"
	}()
	select {
	case res := <-ch2:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("Message 2 timeout")
	}
}
