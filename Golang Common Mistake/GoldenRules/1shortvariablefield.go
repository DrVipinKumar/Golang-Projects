//Can't Use Short Variable Declarations to Set Field Values
package main

import (
	"fmt"
)

type info struct {
	result int
}

func work() (int, error) {
	return 15, nil
}

func main() {
	var data info
	var err error
	data.result, err = work() //error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("info: %+v\n", data)
}
