package routes

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/expenses", handlers.addExpense).Methods("POST")
	router.HandleFunc("/expenses", handlers.getExpense).Methods("GET")
	router.HandleFunc("/expenses/summary", handlers.getExpenseSummary).Methods("GET")
	router.HandleFunc("/expenses/{id}", handlers.updateExpense).Methods("PUT")
	router.HandleFunc("/expenses/{id}", handlers.deleteExpense).Methods("DELETE")
}
