package igutils

import (
	"encoding/json"

	"github.com/samber/lo"
)

type storiesResponse struct {
	Data []struct {
		Id string `json:"id"`
	} `json:"data"`
}

func (c *Client) GetStories() (*[]string, error) {
	id, err := c.GetInstagramBusinessAccoutId()
	if err != nil {
		return nil, err
	}

	// /{ig-user-id}/stories
	stories, err := c.GetWithToken("/"+id+"/stories", nil)
	if err != nil {
		return nil, err
	}

	var storiesResponse storiesResponse
	if err := json.Unmarshal(stories, &storiesResponse); err != nil {
		return nil, err
	}

	ret := lo.Map(storiesResponse.Data, func(d struct {
		Id string `json:"id"`
	}, index int) string {
		return d.Id
	})

	// fmt.Println(string(stories))
	// utils.PrintAsJson(stories)

	return &ret, nil
}
