package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Parse(time.RFC3339, "2018-04-06T10:49:05Z")
	//	if err != nil {

	//}
	fmt.Println(t)
}
