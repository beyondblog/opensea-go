package opensea

import (
	"context"

	"github.com/casiphia/gopkg/restgo"
	"github.com/casiphia/opensea-go/model"
)

// Collections Use this endpoint to fetch collections on OpenSea
func (c *Client) Collections(ctx context.Context, req *CollectionsRequest) ([]*model.Collection, error) {
	var rsp, err = c.get(ctx, "/api/v1/collections", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response CollectionsResponse
	err = ParseRsp(rsp, &response)
	if err != nil {
		return nil, err
	}
	return response.Collections, nil
}

type CollectionsRequest struct {
	// A wallet address. If specified, will return collections where the owner owns at least one asset belonging to smart contracts in the collection.
	// The number of assets the account owns is shown as owned_asset_count for each collection.
	AssetOwner string `query:"asset_owner"`
	// Offset
	Offset int32 `query:"offset,required"`
	// Limit
	Limit int32 `query:"limit,required"`
}

type CollectionsResponse struct {
	Collections []*model.Collection `opensea:"collections"`
}
