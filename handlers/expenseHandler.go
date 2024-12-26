package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Vanamali-Sims/tracks_expenses_brownies/models"
)

var expenses = []models.Expense{}

func addExpense(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil || expense.Category == "" || expense.Amount < 0 || expense.Date == "" {
		http.Error(w, "Bad brownie batch, invalid expense data!", http.StatusBadRequest)
		return
	}

}

func getExpense() {
	json.NewEncoder(w).Encode(expenses)
}

func getExpenseSummary() {

}

func updateExpense() {

}

func deleteExpense() {

}
