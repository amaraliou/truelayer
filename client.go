package truelayer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	// DateLayout ...
	DateLayout = "2006-01-02"
	// TimestampLayout ...
	TimestampLayout             = "2006-01-02T15:04:05Z"
	defaultRetryDuration        = time.Second * 5
	rateLimitExceededStatusCode = 429
)

// Client ...
type Client struct {
	baseURL   string
	http      *http.Client
	AutoRetry bool
}

// Error ...
type Error struct {
	Type        string `json:"error"`
	Description string `json:"error_description"`
	Status      int    `json:"status"`
}

// Error ...
func (e Error) Error() string {
	return e.Type
}

func retryDuration(resp *http.Response) time.Duration {
	raw := resp.Header.Get("Retry-After")
	if raw == "" {
		return defaultRetryDuration
	}
	seconds, err := strconv.ParseInt(raw, 10, 32)
	if err != nil {
		return defaultRetryDuration
	}
	return time.Duration(seconds) * time.Second
}

func (c *Client) get(url string, result interface{}) error {
	for {
		resp, err := c.http.Get(url)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		if resp.StatusCode == rateLimitExceededStatusCode && c.AutoRetry {
			time.Sleep(retryDuration(resp))
			continue
		}
		if resp.StatusCode == http.StatusNoContent {
			return nil
		}
		if resp.StatusCode != http.StatusOK {
			return c.decodeError(resp)
		}

		err = json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			return err
		}

		break
	}

	return nil
}

func (c *Client) decodeError(resp *http.Response) error {
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(responseBody) == 0 {
		return fmt.Errorf("truelayer: HTTP %d: %s (body empty)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	buf := bytes.NewBuffer(responseBody)

	var e struct {
		E Error `json:"error"`
	}
	err = json.NewDecoder(buf).Decode(&e)
	if err != nil {
		return fmt.Errorf("truelayer: couldn't decode error: (%d) [%s]", len(responseBody), responseBody)
	}

	if e.E.Type == "" {
		e.E.Type = fmt.Sprintf("truelayer: unexpected HTTP %d: %s (empty error)",
			resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return e.E
}
