package models

type Expense struct {
	ID       string  `json:id`
	Category string  `json:category`
	Amount   float64 `json:amount`
	Date     string  `json:"date"`
}
