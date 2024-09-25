package igutils

import "errors"

const FacebookEndPoint = "https://graph.facebook.com/v20.0"

type Client struct {
	base  string
	token string
	auth  *auth
}

type auth struct {
	appId  string
	secret string
}

func NewClient(base, token string) *Client {
	client := Client{
		base:  base,
		token: token,
	}
	return &client
}

func NewClientWithoutToken(base string) *Client {
	return NewClient(base, "")
}
func (c *Client) SetAuth(appid, secret string) {
	c.auth = &auth{
		appId:  appid,
		secret: secret,
	}
}

func (c *Client) GetAuth() (*auth, error) {
	if c.auth == nil {
		return nil, errors.New("auth is empty")
	}
	return c.auth, nil
}
