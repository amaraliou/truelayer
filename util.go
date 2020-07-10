package truelayer

import "strings"

// ExtractMerchantFromDescription -> When merchant is nil
func (transaction *Transaction) ExtractMerchantFromDescription() {

	description := transaction.Description
	if transaction.Merchant == "" {
		merchant := strings.Split(description, " -")[0]
		transaction.Merchant = merchant
	}
}
