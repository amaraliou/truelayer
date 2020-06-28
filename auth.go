package truelayer

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// Token ...
type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	Type         string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

// GetAuthURLOptions ...
type GetAuthURLOptions struct {
	RedirectURI  string
	Scope        []string
	State        string
	Providers    []string
	ResponseType string
}

// GetAuthURL ...
func (client *Client) GetAuthURL(opts GetAuthURLOptions) (string, error) {

	if opts.RedirectURI == "" {
		return "", errors.Errorf("Truelayer: RedirectURI is required")
	}

	if len(opts.Scope) == 0 {
		return "", errors.Errorf("Truelayer: Scope is required")
	}

	for _, givenScope := range opts.Scope {
		if !isValidScope(givenScope) {
			return "", errors.Errorf("truelayer: %s is not a valid scope option", givenScope)
		}
	}

	params := url.Values{}
	params.Add("client_id", client.ClientID)
	params.Add("redirect_uri", opts.RedirectURI)
	params.Add("scope", strings.Join(opts.Scope, "%20"))
	params.Add("response_type", "code")
	params.Add("providers", "uk-ob-all")

	authURL := fmt.Sprintf("%s/?%s", AuthBaseURL, params.Encode())
	authURL, err := url.PathUnescape(authURL)
	if err != nil {
		return "", err
	}

	fmt.Print(authURL)
	client.AuthLink = authURL

	return client.AuthLink, nil
}
