package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

type student struct {
	Name   string `json:"name,omitempty"`
	Course string `json:"course,omitempty"`
}

func main() {
	db := []student{
		{"Ravi Kumar", "B.Tech"},
		{"Rahul Kumar", "B.Tech"},
		{"Sanjiv Kumar", "M.Tech"},
		{"Sanjiv Kumar", "P.hD"},
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(db)
	f, err := os.Create("studentdb.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, buf)
}
