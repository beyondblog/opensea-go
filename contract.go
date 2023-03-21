package opensea

import (
	"context"

	"github.com/beyondblog/opensea-go/model"
	"github.com/casiphia/gopkg/restgo"
)

// Contract Used to fetch more in-depth information about an contract asset
func (c *Client) Contract(ctx context.Context, req *ContractRequest) (*model.Contract, error) {
	var rsp, err = c.get(ctx, "/api/v1/asset_contract/:asset_contract_address", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response model.Contract
	err = ParseRsp(rsp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type ContractRequest struct {
	// Address of the contract
	AssetContractAddress string `path:"asset_contract_address,required"`
}
