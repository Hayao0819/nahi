package igutils

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

func (c *Client) GetSemiLongToken() (string, error) {

	auth, err := c.GetAuth()
	if err != nil {
		return "", err
	}

	// gettting semi-long token
	semilong, err := c.Get("/oauth/access_token", map[string]string{
		"grant_type":        "fb_exchange_token",
		"client_id":         auth.appId,
		"client_secret":     auth.secret,
		"fb_exchange_token": c.token,
	})
	if err != nil {
		return "", err
	}
	var parsed_semilong struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(semilong, &parsed_semilong); err != nil {
		return "", err
	}
	return parsed_semilong.AccessToken, nil
}

func (c *Client) GetLongToken() (string, error) {
	semilomg, err := c.GetSemiLongToken()
	if err != nil {
		return "", err
	}
	c.token = semilomg
	slog.Info("get semi-long token", "token", c.token)

	// getting user id
	accounts, err := c.GetAccountsInfo()
	if err != nil {
		return "", err
	}

	slog.Info("get user id", "id", accounts.Id)

	long, err := c.GetWithToken("/"+accounts.Id+"/accounts", map[string]string{})
	if err != nil {
		return "", err
	}
	var parsed_long struct {
		Data []struct {
			AccessToken string `json:"access_token"`
		}
	}

	if err := json.Unmarshal(long, &parsed_long); err != nil {
		return "", err
	}

	if len(parsed_long.Data) == 0 {
		return "", fmt.Errorf("no page found")
	}

	return parsed_long.Data[0].AccessToken, nil
}
