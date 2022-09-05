package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	fmt.Println("Before timer start")
	<-timer1.C
	fmt.Println("After 2 second timer expired")
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Inside function timer 2 expired")
	}()
	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer stoped before expire")
	}
}
