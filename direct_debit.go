package truelayer

import "time"

// DirectDebit ...
type DirectDebit struct {
	ID                    string                 `json:"direct_debit_id"`
	Date                  time.Time              `json:"timestamp"`
	Name                  string                 `json:"name"`
	Status                string                 `json:"status"`
	PreviousPaymentDate   time.Time              `json:"previous_payment_timestamp"`
	PreviousPaymentAmount string                 `json:"previous_payment_amount"`
	Currency              string                 `json:"currency"`
	Metadata              map[string]interface{} `json:"meta"`
}
