package main

import (
	"fmt"
	"time"
)

func main() {
	limit := 1000
	start := time.Now()
	var mylist []int
	for i := 0; i < limit; i++ {
		mylist = append(mylist, i)
		fmt.Println(mylist[i])
	}

	duration := time.Since(start)
	fmt.Println(duration)
}
