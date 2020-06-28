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

// CardBalance ...
type CardBalance struct {
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
