package igutils

import "encoding/json"

func (c *Client) GetPageId() (string, error) {
	res, err := c.Get("/me", map[string]string{
		"fields": "id",
	})
	if err != nil {
		return "", err
	}

	parsed := struct {
		Id string `json:"id"`
	}{}
	if err := json.Unmarshal(res, &parsed); err != nil {
		return "", err
	}

	return parsed.Id, nil
}
