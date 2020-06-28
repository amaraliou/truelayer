package truelayer

import "net/http"

// Client ...
type Client struct {
	baseURL string
	http    *http.Client
}
