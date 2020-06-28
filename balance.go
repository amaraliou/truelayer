package truelayer

import "time"

// AccountBalance ...
type AccountBalance struct {
	Available float32   `json:"available"`
	Currency  string    `json:"currency"`
	Current   float32   `json:"current"`
	Overdraft float32   `json:"overdraft"`
	UpdatedAt time.Time `json:"update_timestamp"`
}

// TransactionBalance ...
type TransactionBalance struct {
	Amount   float32 `json:"amount"`
	Currency string  `json:"currency"`
}

// CardBalance ...
type CardBalance struct {
	Available            float32   `json:"available"`
	Currency             string    `json:"currency"`
	Current              float32   `json:"current"`
	CreditLimit          float32   `json:"credit_limit"`
	LastStatementBalance float32   `json:"last_statement_balance"`
	LastStatementDate    string    `json:"last_statement_date"`
	PaymentDue           float32   `json:"payment_due"`
	PaymentDueDate       string    `json:"payment_due_date"`
	UpdatedAt            time.Time `json:"update_timestamp"`
}

// GetAccountBalance ...
func (client *Client) GetAccountBalance(accountID string) (*AccountBalance, error) {

	truelayerURL := client.baseURL + "data/v1/accounts/" + accountID + "/balance"

	var balance struct {
		Results []*AccountBalance `json:"results"`
	}

	err := client.get(truelayerURL, &balance)
	if err != nil {
		return nil, err
	}

	return balance.Results[0], nil
}

// GetCardBalance ...
func (client *Client) GetCardBalance(accountID string) (*CardBalance, error) {

	truelayerURL := client.baseURL + "data/v1/cards/" + accountID + "/balance"

	var balance struct {
		Results []*CardBalance `json:"results"`
	}

	err := client.get(truelayerURL, &balance)
	if err != nil {
		return nil, err
	}

	return balance.Results[0], nil
}
