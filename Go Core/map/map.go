package main

import "fmt"

func main() {
	var temp = make(map[string]int)
	temp["delhi"] = 35
	temp["muradnagar"] = 31
	temp["Ghaziabad"] = 32
	fmt.Println(temp)
	var delhi = temp["delhi"]
	fmt.Println("Delhi Temperature=", delhi)
	delete(temp, "delhi")
	fmt.Println(temp)
	_, check := temp["Ghaziabad"]
	fmt.Println(check)
}
