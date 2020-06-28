package truelayer

import "time"

// AccountBalance ...
type AccountBalance struct {
	Currency  string    `json:"currency"`
	Available float32   `json:"available"`
	Current   float32   `json:"current"`
	Overdraft float32   `json:"overdraft"`
	UpdatedAt time.Time `json:"update_timestamp"`
}

// TransactionBalance ...
type TransactionBalance struct {
	Amount   float32 `json:"amount"`
	Currency string  `json:"currency"`
}
