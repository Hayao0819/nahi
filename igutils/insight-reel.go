package igutils

import (
	"log/slog"
)

type ReelInsight struct {
	Shares                    int
	Comments                  int
	Plays                     int
	Likes                     int
	Saved                     int
	VideoViews                int
	TotalInteractions         int
	Reach                     int
	IgReelsVideoViewTotalTime int
	IgReelsAvgWatchTime       int
	ClipsReplaysCount         int
}

func (c *Client) GetReelInsight(id string) (*ReelInsight, error) {
	slog.Info("Getting reel insight", "id", id)
	// insight, err := c.GetWithToken("/"+id+"/insights", map[string]string{
	// 	"metric": "shares,comments,plays,likes,saved,video_views,total_interactions,reach,ig_reels_video_view_total_time,ig_reels_avg_watch_time,clips_replays_count,",
	// })

	// if err != nil {
	// 	return nil, err
	// }

	// res := insightResponse{}
	// if err := json.Unmarshal(insight, &res); err != nil {
	// 	return nil, err
	// }

	res, err := c.getInsight(id, "shares", "comments", "plays", "likes", "saved", "video_views", "total_interactions", "reach", "ig_reels_video_view_total_time", "ig_reels_avg_watch_time", "clips_replays_count")
	if err != nil {
		return nil, err
	}

	slog.Debug("Reel Insight", "insight", res)

	ret := ReelInsight{}
	for _, d := range res.Data {
		switch d.Name {
		case "shares":
			ret.Shares = d.Values[0].Value
		case "comments":
			ret.Comments = d.Values[0].Value
		case "plays":
			ret.Plays = d.Values[0].Value
		case "likes":
			ret.Likes = d.Values[0].Value
		case "saved":
			ret.Saved = d.Values[0].Value
		case "video_views":
			ret.VideoViews = d.Values[0].Value
		case "total_interactions":
			ret.TotalInteractions = d.Values[0].Value
		case "reach":
			ret.Reach = d.Values[0].Value
		case "ig_reels_video_view_total_time":
			ret.IgReelsVideoViewTotalTime = d.Values[0].Value
		case "ig_reels_avg_watch_time":
			ret.IgReelsAvgWatchTime = d.Values[0].Value
		case "clips_replays_count":
			ret.ClipsReplaysCount = d.Values[0].Value

		default:
			slog.Warn("Unknown insight", "name", d.Name)
		}
	}

	return &ret, nil
}
