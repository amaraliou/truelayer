package truelayer

import "net/http"

// Client ...
type Client struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	AuthLink     string
	Code         string
	AccessToken  string
	http         *http.Client
}

// NewClient ...
func NewClient(clientID, clientSecret string) *Client {
	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		http:         &http.Client{},
	}
}
