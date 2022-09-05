package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getRand() int {
	return rand.Intn(100)
}
func displayMessage(ch chan string) {
	wg := sync.WaitGroup{}
	for i := 1; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
			ch <- fmt.Sprintf("This is value %d in channel", getRand())
		}()
	}
	wg.Wait()
	close(ch)
}
func main() {
	ch := make(chan string, 3)
	//special variable that is use to exchange information amoung goroutines
	go displayMessage(ch)
	for chv := range ch {
		fmt.Println(chv)
	}

}
