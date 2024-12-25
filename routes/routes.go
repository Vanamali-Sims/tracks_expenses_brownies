package routes

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/expenses", handlers.addExpense).Methods("POST")
}
