package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type Task struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Title   string         `json:"title" yaml:"title"`
	ID      string         `json:"id" yaml:"id"`
}
