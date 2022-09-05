package main

import (
	"fmt"
	"time"
)

func main() {
	//file,err := os.Open("abc.txt")
	//defer file.Close()
	t, err := time.Parse(time.RFC3339, "2018-04-06T10:49:05Z")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
}
