package truelayer

const (
	// AuthBaseURL ...
	AuthBaseURL = "https://auth.truelayer.com"

	// APIBaseURL ...
	APIBaseURL = "https://api.truelayer.com/"

	// StatusBaseURL ...
	StatusBaseURL = "https://status-api.truelayer.com"

	// DefaultTimeout ...
	DefaultTimeout = 60000
)

var scopes = []string{"accounts", "info", "transactions", "balance", "cards", "offline_access", "direct_debits", "standing_orders"}

func isValidScope(givenScope string) bool {
	for _, scope := range scopes {
		if scope == givenScope {
			return true
		}
	}

	return false
}
