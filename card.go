package truelayer

import "time"

// Card ..
type Card struct {
	AccountID     string    `json:"account_id"`
	Network       string    `json:"card_network"`
	Type          string    `json:"card_type"`
	Currency      string    `json:"currency"`
	DisplayName   string    `json:"display_name"`
	PartialNumber string    `json:"partial_card_number"`
	NameOnCard    string    `json:"name_on_card"`
	ValidFrom     string    `json:"valid_from"`
	ValidTo       string    `json:"valid_to"`
	UpdatedAt     time.Time `json:"update_timestamp"`
	Provider      *Provider `json:"provider"`
}

// GetCards ...
func (client *Client) GetCards() ([]*Card, error) {

	truelayerURL := client.baseURL + "data/v1/cards"

	var cards struct {
		Results []*Card `json:"results"`
	}

	err := client.get(truelayerURL, &cards)
	if err != nil {
		return nil, err
	}

	return cards.Results, nil
}

// GetCard ...
func (client *Client) GetCard(accountID string) (*Card, error) {

	truelayerURL := client.baseURL + "data/v1/cards/" + accountID

	var card struct {
		Results []*Card `json:"results"`
	}

	err := client.get(truelayerURL, &card)
	if err != nil {
		return nil, err
	}

	return card.Results[0], nil
}
