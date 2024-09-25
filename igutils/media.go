package igutils

import (
	"encoding/json"
	"log/slog"

	"github.com/cockroachdb/errors"
)

type Media struct {
	Caption   string `json:"caption,omitempty"`
	MediaType string `json:"media_type"`
	MediaURL  string `json:"media_url"`
	LikeCount int    `json:"like_count"`
	Permalink string `json:"permalink"`
	Timestamp string `json:"timestamp"`
	Username  string `json:"username"`
	ID        string `json:"id"`
	Children  struct {
		Data []struct {
			MediaURL  string `json:"media_url"`
			MediaType string `json:"media_type"`
			ID        string `json:"id"`
		} `json:"data"`
	} `json:"children,omitempty"`
}

type RecentMediaResponse struct {
	Name  string        `json:"name"`
	Media MediaResponse `json:"media"`
	ID    string        `json:"id"`
}

type MediaResponse struct {
	Data   []Media `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
		Next string `json:"next"`
	} `json:"paging"`
}

func (c *Client) GetRecentMediaList() (*RecentMediaResponse, error) {
	id, err := c.GetInstagramBusinessAccoutId()
	if err != nil {
		return nil, err
	}

	mediaList, err := c.GetWithToken("/"+id+"", map[string]string{
		"fields": "name,media{caption,media_type,children{media_url,media_type,thumbnail_url},thumbnail_url,media_url,like_count,permalink,timestamp,username}",
	})
	if err != nil {
		return nil, err
	}

	res := RecentMediaResponse{}
	if err := json.Unmarshal(mediaList, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

var mediaListRecursiveCount = 0

func (c *Client) getMediaListRecursive(url string, arr *[]*MediaResponse) error {
	mediaListRecursiveCount++
	resData, err := c.GetFromRawURL(url)
	if err != nil {
		return err
	}

	// fmt.Println(string(resData))

	res := MediaResponse{}
	if err := json.Unmarshal(resData, &res); err != nil {
		return err
	}
	*arr = append(*arr, &res)

	if res.Paging.Next == "" {
		return nil
	}
	return c.getMediaListRecursive(res.Paging.Next, arr)
}

func (c *Client) GetAllMediaList() ([]*MediaResponse, error) {
	ret := []*MediaResponse{}

	recent, err := c.GetRecentMediaList()
	if err != nil {
		return nil, err
	}
	ret = append(ret, &recent.Media)

	if err := c.getMediaListRecursive(recent.Media.Paging.Next, &ret); err != nil {
		return nil, errors.Wrap(err, "failed to get all media list")
	}

	slog.Debug("Getting all media list done", "count", mediaListRecursiveCount)
	return ret, nil
}
