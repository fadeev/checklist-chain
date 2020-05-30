package keeper

import (
	"github.com/fadeev/checklist-chain/x/checklist/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for checklist clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryListTasks:
			return listTasks(ctx, k)
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

func listTasks(ctx sdk.Context, k Keeper) ([]byte, error) {
	var taskList []types.Task
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.TaskPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var task types.Task
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &task)
		taskList = append(taskList, task)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, taskList)
	return res, nil
}
