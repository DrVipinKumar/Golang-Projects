package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Finally it is running now")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleFunc)
	http.ListenAndServe(":9999", r)
}
