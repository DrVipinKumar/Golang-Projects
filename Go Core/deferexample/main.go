package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//defer fmt.Println("World!")
	//fmt.Println("Hello")

	//for i := 1; i < 6; i++ {
	//		defer fmt.Println(i)  //LIFO order stack
	//}
	f, err := os.Open("./goroutines/main.go")
	defer f.Close()
	if err != nil {
		fmt.Println("fil not found")
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
