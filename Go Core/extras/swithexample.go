package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekdays")
	}
	var i = 4
	switch i {
	case 0:
		fmt.Println("Sunday")
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
		fallthrough
	case 5:
		fmt.Println("Friday")
		fallthrough
	case 6:
		fmt.Println("Saturday")
	}
}
