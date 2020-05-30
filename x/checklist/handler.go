package checklist

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/fadeev/checklist-chain/x/checklist/types"
	"github.com/google/uuid"
)

// NewHandler creates an sdk.Handler for all the checklist type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// TODO: Define your msg cases
		//
		//Example:
		// case Msg<Action>:
		// 	return handleMsg<Action>(ctx, k, msg)
		case MsgCreateTask:
			return handleMsgCreateTask(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgCreateTask(ctx sdk.Context, k Keeper, msg MsgCreateTask) (*sdk.Result, error) {
	var task = types.Task{
		Creator: msg.Creator,
		ID:      uuid.New().String(),
		Title:   msg.Title,
	}
	k.CreateTask(ctx, task)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
