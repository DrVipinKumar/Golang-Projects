package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}
type userdb struct {
	Users []user `json:"users,omitempty"`
	Types string `json:"types,omitempty"`
}

func main() {
	db := userdb{}
	f, err := os.Open("./usersdb.json")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(f)
	dec.Decode(&db)
	fmt.Println(db)
}
