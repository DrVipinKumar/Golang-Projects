package main

import (
	"bytes"
	"encoding/json"
	"io"
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
	users := []user{
		{"Dr. Vipin Kumar", "Muradnagar"},
		{"Mr. Sachin Kumar", "Meerut"},
		{"Mr. Ashok Tyagi", "Delhi"},
	}
	db := userdb{Users: users, Types: "Programming"}
	f, err := os.Create("./usersdb.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(&db)
	io.Copy(f, buf)

}
