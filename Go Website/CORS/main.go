//golang preflight request error
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome World of CORS")
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", GetUsers).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH"},
		AllowedHeaders: []string{"Bearer", "Content_Type"}})

	// c := cors.New(cors.Options{
	// 	AllowedMethods: []string{"GET","POST", "OPTIONS"},
	// 	AllowedOrigins: []string{"*"},
	// 	AllowCredentials: true,
	// 	AllowedHeaders: []string{"Content-Type","Bearer","Bearer ","content-type","Origin","Accept"},
	// 	OptionsPassthrough: true,
	// })
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
