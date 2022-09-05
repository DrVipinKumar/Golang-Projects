package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	Name   string `json:"name,omitempty"`
	Course string `json:"course,omitempty"`
}

func main() {
	db := []student{}
	f, err := os.Open("studentdb.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&db)
	fmt.Println(db)

}
