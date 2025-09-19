package main

import (
	"fmt"
	"net/http"

	"github.com/divyathakkar3112/Library/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting Library")

	// writing api
	r := mux.NewRouter()
	r.HandleFunc("/home", handlers.HomePage).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUsers).Methods("POST")
	r.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")

	http.ListenAndServe(":8000", r)
}
