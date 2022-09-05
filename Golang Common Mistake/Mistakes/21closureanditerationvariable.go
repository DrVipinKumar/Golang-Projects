package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Print(n, " ")
			wg.Done()
		}(i)

	}
	wg.Wait()
	fmt.Println()
}
