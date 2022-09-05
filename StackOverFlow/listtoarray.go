package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Returnjson struct {
	list []string `json:"list,omitempty"`
}

func returnjson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var returnjson Returnjson
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&returnjson)
	if err != nil {
		w.WriteHeader(400)
	}
	var arr []string

	for index := 0; index < 3; index++ {
		arr[index] = returnjson.list[index]
	}

	printall := arr[0]
	w.Write([]byte(printall))

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/json", returnjson).Methods("POST")
	http.ListenAndServe(":8080", r)
}
