package truelayer

import "time"

// StandingOrder ...
type StandingOrder struct {
	Frequency          string                 `json:"frequency"`
	Status             string                 `json:"status"`
	Timestamp          time.Time              `json:"timestamp"`
	Currency           string                 `json:"currency"`
	Metadata           map[string]interface{} `json:"meta"`
	NextPaymentDate    time.Time              `json:"next_payment_date"`
	NextPaymentAmount  float32                `json:"next_payment_amount"`
	FirstPaymentDate   time.Time              `json:"first_payment_date"`
	FirstPaymentAmount float32                `json:"first_payment_amount"`
	LastPaymentDate    time.Time              `json:"last_payment_date"`
	LastPaymentAmount  float32                `json:"last_payment_amount"`
	Reference          string                 `json:"reference"`
	Payee              string                 `json:"payee"`
}

// GetStandingOrders ...
func (client *Client) GetStandingOrders(accountID string) ([]*StandingOrder, error) {

	truelayerURL := client.baseURL + "data/v1/accounts/" + accountID + "standing_orders"

	var standingOrders struct {
		Results []*StandingOrder `json:"results"`
	}

	err := client.get(truelayerURL, &standingOrders)
	if err != nil {
		return nil, err
	}

	return standingOrders.Results, nil
}
