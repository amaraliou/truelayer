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

// GetDirectDebits ...
func (client *Client) GetDirectDebits(accountID string) ([]*DirectDebit, error) {

	truelayerURL := client.baseURL + "data/v1/accounts/" + accountID + "/direct_debits"

	var directDebits struct {
		Results []*DirectDebit `json:"results"`
	}

	err := client.get(truelayerURL, &directDebits)
	if err != nil {
		return nil, err
	}

	return directDebits.Results, nil
}
