package igutils

type StoryInsight struct {
	Impressions       int
	Shares            int
	Replies           int
	TotalInteractions int
	Navigation        int
	Follows           int
	ProfileVisits     int
	ProfileActivity   int
	Reach             int
}

func (c *Client) GetStoryInsight(mediaID string) (*StoryInsight, error) {

	res, err := c.getInsight(mediaID, "impressions", "shares", "replies", "total_interactions", "navigation", "follows", "profile_visits", "profile_activity", "reach")
	if err != nil {
		return nil, err
	}

	ret := StoryInsight{}
	for _, d := range res.Data {
		switch d.Name {
		case "impressions":
			ret.Impressions = d.Values[0].Value
		case "shares":
			ret.Shares = d.Values[0].Value
		case "replies":
			ret.Replies = d.Values[0].Value
		case "total_interactions":
			ret.TotalInteractions = d.Values[0].Value
		case "navigation":
			ret.Navigation = d.Values[0].Value
		case "follows":
			ret.Follows = d.Values[0].Value
		case "profile_visits":
			ret.ProfileVisits = d.Values[0].Value
		case "profile_activity":
			ret.ProfileActivity = d.Values[0].Value
		case "reach":
			ret.Reach = d.Values[0].Value
		default:
		}
	}

	return &ret, nil
}
