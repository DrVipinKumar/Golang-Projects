package main

import "fmt"

func main() {
	x := map[string]string{"one": "1", "two": "2", "three": "3"}

	if value, v := x["one"]; !v {
		fmt.Println("Key not Found")
	} else {
		fmt.Println("Key Found", value)
	}

}
