package igutils

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
)

func (c *Client) GetWithToken(path string, params map[string]string) ([]byte, error) {
	if c.token == "" {
		return nil, errors.New("token is empty")
	}

	if params == nil {
		params = make(map[string]string)
	}
	params["access_token"] = c.token
	return c.Get(path, params)
}

func (c *Client) GetFromRawURL(url string) ([]byte, error) {
	slog.Debug("GET", "url", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	requestCount++
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	slog.Debug("GET response", "body", string(respBody))
	return respBody, nil
}

func (c *Client) Get(path string, params map[string]string) ([]byte, error) {
	url := c.base + path

	// Add query parameters
	if len(params) > 0 {
		url += "?"
		for key, value := range params {
			url += key + "=" + value + "&"
		}
	}

	// Send GET request
	return c.GetFromRawURL(url)
}
