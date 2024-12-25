package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	log.Println("Expense Tracker API (brownie version) is running on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
