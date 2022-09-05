package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to first web page in golang  by url %s", r.URL.Path)

}
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":9999", nil)
}
