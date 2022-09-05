package main

import (
	"fmt"
	"time"
)

func main() {
	//n := 1000
	var n time.Duration = 1000
	time.Sleep(n * time.Millisecond)
	fmt.Println("Welcome to Go Timing System")
}
