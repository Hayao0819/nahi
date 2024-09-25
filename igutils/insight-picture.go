package igutils

import (
	"log/slog"
)

type PictureInsight struct {
	Impressions       int // 閲覧された合計回数
	Reach             int // 閲覧したユニークInstagramアカウントの合計数
	Saved             int // 保存したユニークInstagramアカウントの合計数
	VideoViews        int // 動画IGメディアが閲覧された合計回数
	ProfileVisits     int
	Comments          int
	Follows           int
	Likes             int
	Shares            int
	TotalInteractions int
}

func (c *Client) GetPicAndVideoInsight(mediaID string) (*PictureInsight, error) {
	slog.Info("Getting pic and video insight", "id", mediaID)

	res, err := c.getInsight(mediaID,
		"impressions",
		"reach",
		"saved",
		"video_views",
		"profile_visits",
		"comments",
		"follows",
		"likes",
		"profile_activity",
		"profile_visits",
		"shares",
		"total_interactions",
	)

	if err != nil {
		return nil, err
	}

	ret := PictureInsight{}

	for _, d := range res.Data {
		switch d.Name {
		case "impressions":
			ret.Impressions = d.Values[0].Value
		case "reach":
			ret.Reach = d.Values[0].Value
		case "saved":
			ret.Saved = d.Values[0].Value
		case "video_views":
			ret.VideoViews = d.Values[0].Value
		case "profile_visits":
			ret.ProfileVisits = d.Values[0].Value
		case "comments":
			ret.Comments = d.Values[0].Value
		case "follows":
			ret.Follows = d.Values[0].Value
		case "likes":
			ret.Likes = d.Values[0].Value
		case "profile_activity":
			ret.ProfileVisits = d.Values[0].Value
		case "shares":
			ret.Shares = d.Values[0].Value
		case "total_interactions":
			ret.TotalInteractions = d.Values[0].Value
		default:
			slog.Warn("Unknown insight", "name", d.Name)
		}

	}

	return &ret, nil
}
