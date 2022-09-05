package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	p := Person{"Dr Vipin", 22}
	jsonData, _ := json.Marshal(p)
	fmt.Println(string(jsonData))
}
