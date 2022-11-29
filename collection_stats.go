package opensea

import (
	"context"

	"github.com/casiphia/gopkg/restgo"
	"github.com/casiphia/opensea-go/model"
)

// CollectionStats Use this endpoint to fetch stats for a specific collection,
// including realtime floor price statistics
func (c *Client) CollectionStats(ctx context.Context, req *CollectionRequest) (*model.CollectionStats, error) {
	var rsp, err = c.get(ctx, "/api/v1/collection/:collection_slug/stats", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response CollectionStatsResponse
	err = ParseRsp(rsp, &response)
	if err != nil {
		return nil, err
	}
	return response.Stats, nil
}

type CollectionStatsResponse struct {
	Stats *model.CollectionStats `opensea:"stats"`
}
