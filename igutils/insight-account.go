package igutils

import (
	"encoding/json"
	"log/slog"
)

type AccountInsight struct {
	Impressions []struct {
		Value   int
		EndTime string
	}
	Reach []struct {
		Value   int
		EndTime string
	}
	ProfileViews []struct {
		Value   int
		EndTime string
	}
}

var accountInsight *AccountInsight = nil

func (c *Client) GetAccountInsight() (*AccountInsight, error) {
	slog.Info("Getting account insight")
	if accountInsight != nil {
		return accountInsight, nil
	}

	info, err := c.GetAccountsInfo()
	if err != nil {
		return nil, err
	}
	insight, err := c.GetWithToken("/"+info.InstagramBusinessAccount.Id+"/insights", map[string]string{
		"metric": "impressions,reach,profile_views",
		"period": "day",
	})
	if err != nil {
		return nil, err
	}
	// utils.PrintAsJson(string(insight))

	res := insightResponse{}
	if err := json.Unmarshal(insight, &res); err != nil {
		return nil, err
	}

	ret := AccountInsight{}
	for _, d := range res.Data {
		if d.Name == "impressions" {
			for _, v := range d.Values {
				ret.Impressions = append(ret.Impressions, struct {
					Value   int
					EndTime string
				}{v.Value, v.EndTime})
			}
		}
		if d.Name == "reach" {
			for _, v := range d.Values {
				ret.Reach = append(ret.Reach, struct {
					Value   int
					EndTime string
				}{v.Value, v.EndTime})
			}
		}
		if d.Name == "profile_views" {
			for _, v := range d.Values {
				ret.ProfileViews = append(ret.ProfileViews, struct {
					Value   int
					EndTime string
				}{v.Value, v.EndTime})
			}
		}
	}

	accountInsight = &ret

	return &ret, nil
}
