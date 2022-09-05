package main

import (
	"fmt"
	"time"
)

func msg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, "=", i)
	}
}
func msg2(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, "=", i)
	}
}
func main() {
	go msg("No threads")
	go msg2("goroutines")
	msg("this is simple method")
	time.Sleep(time.Second)
}
