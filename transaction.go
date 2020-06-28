package truelayer

import "time"

// Transaction ...
type Transaction struct {
	ID             string                 `json:"transaction_id"`
	Date           time.Time              `json:"timestamp"`
	Description    string                 `json:"description"`
	Type           string                 `json:"transaction_type"`
	Category       string                 `json:"transaction_category"`
	Classification []string               `json:"classification"`
	Merchant       string                 `json:"merchant_name"`
	Amount         float32                `json:"amount"`
	Currency       string                 `json:"currency"`
	Metadata       map[string]interface{} `json:"meta"`
	RunningBalance *TransactionBalance    `json:"running_balance"`
}

func (client *Client) getTransactions(productType, accountID string) ([]*Transaction, error) {

	truelayerURL := client.baseURL + "data/v1/" + productType + "/" + accountID + "/transactions"

	var transactions struct {
		Results []*Transaction `json:"results"`
	}

	err := client.get(truelayerURL, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions.Results, nil
}

func (client *Client) getPendingTransactions(productType, accountID string) ([]*Transaction, error) {

	truelayerURL := client.baseURL + "data/v1/" + productType + "/" + accountID + "/transactions/pending"

	var pendingTransactions struct {
		Results []*Transaction `json:"results"`
	}

	err := client.get(truelayerURL, &pendingTransactions)
	if err != nil {
		return nil, err
	}

	return pendingTransactions.Results, nil
}

// GetAccountTransactions ...
func (client *Client) GetAccountTransactions(accountID string, pending bool) ([]*Transaction, error) {

	if pending {
		return client.getPendingTransactions("accounts", accountID)
	}

	return client.getTransactions("accounts", accountID)
}

// GetCardTransactions ...
func (client *Client) GetCardTransactions(accountID string, pending bool) ([]*Transaction, error) {

	if pending {
		return client.getPendingTransactions("cards", accountID)
	}

	return client.getTransactions("cards", accountID)
}
