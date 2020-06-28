package truelayer

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

// Authenticator ...
type Authenticator struct {
	config  *oauth2.Config
	context context.Context
}

// NewAuthenticator ...
func NewAuthenticator(redirectURL, clientID, clientSecret string, scopes ...string) Authenticator {
	cfg := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthBaseURL,
			TokenURL: fmt.Sprintf("%s/connect/token", AuthBaseURL),
		},
	}

	tr := &http.Transport{
		TLSNextProto: map[string]func(authority string, c *tls.Conn) http.RoundTripper{},
	}
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{Transport: tr})
	return Authenticator{
		config:  cfg,
		context: ctx,
	}
}

// SetAuthInfo ...
func (a *Authenticator) SetAuthInfo(clientID, secretKey string) {
	a.config.ClientID = clientID
	a.config.ClientSecret = secretKey
}

// AuthURL ...
func (a Authenticator) AuthURL(state string) string {
	return a.config.AuthCodeURL(state)
}

// AuthURLWithDialog ...
func (a Authenticator) AuthURLWithDialog(state string) string {
	return a.config.AuthCodeURL(state, oauth2.SetAuthURLParam("show_dialog", "true"))
}

// Token ...
func (a Authenticator) Token(state string, r *http.Request) (*oauth2.Token, error) {
	values := r.URL.Query()
	if e := values.Get("error"); e != "" {
		return nil, errors.New("truelayer: auth failed - " + e)
	}
	code := values.Get("code")
	if code == "" {
		return nil, errors.New("truelayer: didn't get access code")
	}
	actualState := values.Get("state")
	if actualState != state {
		return nil, errors.New("truelayer: redirect state parameter doesn't match")
	}
	return a.config.Exchange(a.context, code)
}

// Exchange ...
func (a Authenticator) Exchange(code string) (*oauth2.Token, error) {
	return a.config.Exchange(a.context, code)
}

// NewClient ...
func (a Authenticator) NewClient(token *oauth2.Token) Client {
	client := a.config.Client(a.context, token)
	return Client{
		http:    client,
		baseURL: APIBaseURL,
	}
}

// Token gets the client's current token.
func (c *Client) Token() (*oauth2.Token, error) {
	transport, ok := c.http.Transport.(*oauth2.Transport)
	if !ok {
		return nil, errors.New("truelayer: oauth2 transport type not correct")
	}
	t, err := transport.Source.Token()
	if err != nil {
		return nil, err
	}

	return t, nil
}
