package igutils

import (
	"encoding/json"
	"log/slog"
)

// func (c *Client) GetUserId() (string, error) {
// 	user_id, err := c.GetWithToken("/me", map[string]string{"fields": "id"})
// 	if err != nil {
// 		return "", err
// 	}

// 	parsed_userId := struct {
// 		Id string `json:"id"`
// 	}{}
// 	if err := json.Unmarshal(user_id, &parsed_userId); err != nil {
// 		return "", err
// 	}
// 	return parsed_userId.Id, nil
// }

type AccountsInfo struct {
	InstagramBusinessAccount struct {
		Id string `json:"id"`
	} `json:"instagram_business_account"`
	ConnectedInstagramAccount struct {
		Id string `json:"id"`
	} `json:"connected_instagram_account"`
	InstagramAccounts struct {
		Data []struct {
			Id       string `json:"id"`
			Username string `json:"username"`
		} `json:"data"`
	} `json:"instagram_accounts"`
	Id string `json:"id"`
}

func (c *Client) GetAccountsInfo() (*AccountsInfo, error) {
	slog.Info("Getting account info")

	//me?fields=instagram_business_account,connected_instagram_account,instagram_accounts{id,username}
	res, err := c.GetWithToken("/me", map[string]string{
		"fields": "instagram_business_account,connected_instagram_account,instagram_accounts{id,username}",
	})
	if err != nil {
		return nil, err
	}
	info := AccountsInfo{}
	if err := json.Unmarshal(res, &info); err != nil {
		return nil, err
	}

	slog.Info("Account info", "id", info.Id)
	return &info, nil
}

func (c *Client) GetInstagramBusinessAccoutId() (string, error) {
	accounts, err := c.GetAccountsInfo()
	if err != nil {
		return "", err
	}
	return accounts.InstagramBusinessAccount.Id, nil
}
