package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getRand() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}
func main() {
	ch := make(chan string)
	//channel are the special variable that use to exchange information
	//amoung goroutines (lightweight threads)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 15; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				r := getRand()
				ch <- fmt.Sprintf("Message %d...", r)
			}()
		}

		wg.Wait()
		close(ch)
	}()
	for vch := range ch {
		fmt.Println(vch)
	}
}
