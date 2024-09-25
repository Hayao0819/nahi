package igutils

import (
	"encoding/json"
	"log/slog"
	"strings"

	"github.com/cockroachdb/errors"
)

type insightError struct {
	Message string `json:"message"`
}

type insightValue struct {
	Value   int    `json:"value"`
	EndTime string `json:"end_time"`
}

type insightData struct {
	Name        string         `json:"name"`
	Period      string         `json:"period"`
	Values      []insightValue `json:"values"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Id          string         `json:"id"`
}

type insightResponse struct {
	Data  []insightData `json:"data"`
	Error insightError  `json:"error"`
}

func (c *Client) getInsight(id string, metric ...string) (*insightResponse, error) {
	insight, err := c.GetWithToken("/"+id+"/insights", map[string]string{
		"metric": strings.Join(metric, ","),
	})
	if err != nil {
		return nil, err
	}

	slog.Debug("Insight", "id", id, "metric", metric, "insight", string(insight))

	res := insightResponse{}
	if err := json.Unmarshal(insight, &res); err != nil {
		return nil, err
	}

	// utils.PrintAsJson(res)

	if res.Error.Message != "" {
		return nil, errors.Newf("Insight error: %s", res.Error.Message)
	}

	return &res, nil
}
