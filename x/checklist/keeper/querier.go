package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for checklist clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// case types.QueryParams:
		// 	return queryParams(ctx, k)
		// TODO: Put the modules query routes
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown checklist query endpoint")
		}
	}
}

// func queryParams(ctx sdk.Context, k Keeper) ([]byte, error) {
// 	params := k.GetParams(ctx)

// 	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
// 	if err != nil {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
// 	}

// 	return res, nil
// }

// TODO: Add the modules query functions
// They will be similar to the above one: queryParams()
