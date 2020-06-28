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

// Provider ...
type Provider struct {
	ID          string `json:"provider_id"`
	DisplayName string `json:"display_name"`
	LogoURI     string `json:"logo_uri"`
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
