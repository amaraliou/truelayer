package truelayer

import "time"

// Account ...
type Account struct {
	ID          string    `json:"account_id"`
	Type        string    `json:"account_type"`
	Currency    string    `json:"currency"`
	DisplayName string    `json:"display_name"`
	UpdatedAt   time.Time `json:"update_timestamp"`
	Description string    `json:"description"`
	Number      *Number   `json:"account_number"`
	Provider    *Provider `json:"provider"`
}

// Number ...
type Number struct {
	IBAN     string `json:"iban"`
	Number   string `json:"number"`
	SortCode string `json:"sort_code"`
	SwiftBic string `json:"swift_bic"`
}

// GetAccounts ...
func (client *Client) GetAccounts() ([]*Account, error) {

	truelayerURL := client.baseURL + "data/v1/accounts"

	var accounts struct {
		Results []*Account `json:"results"`
	}

	err := client.get(truelayerURL, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts.Results, nil
}

// GetAccount ...
func (client *Client) GetAccount(accountID string) (*Account, error) {

	truelayerURL := client.baseURL + "data/v1/accounts/" + accountID

	var account struct {
		Results []*Account `json:"results"`
	}

	err := client.get(truelayerURL, &account)
	if err != nil {
		return nil, err
	}

	return account.Results[0], nil
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
