package main
import (
	"sync"
	"fmt"
)
var count=0
func increment(wg *sync.WaitGroup, m *sync.Mutex){
	defer wg.Done()
	m.Lock()
	count=count+1  
	m.Unlock()
}
func main(){
	var wg sync.WaitGroup
	var m sync.Mutex
	for i:=0;i<10000;i++{
		wg.Add(1)
		go increment(&wg,&m)
	}
	wg.Wait()
fmt.Println("Count:",count)
}