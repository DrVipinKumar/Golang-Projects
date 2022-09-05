package main

import "fmt"

func main() {
	var mdata = make(map[string]int)
	mdata["key1"] = 34
	mdata["key2"] = 45
	mdata["key3"] = 56
	for key, value := range mdata {
		fmt.Println(key, "=>", value)
	}
	var data = [4]int{11, 12, 33, 34}
	for _, value := range data {
		fmt.Print(value, " ")
	}
	fmt.Println()
	var sdata = make([]int, 3)
	sdata[0] = 14
	sdata[1] = 45
	sdata[2] = 56
	for _, value := range sdata {
		fmt.Print(value, " ")
	}
}
