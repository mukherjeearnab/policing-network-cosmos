package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
)

// GetJudgementCount get the total number of judgement
func (k Keeper) GetJudgementCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.JudgementCountPrefix)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetJudgementCount set the total number of judgement
func (k Keeper) SetJudgementCount(ctx sdk.Context, count int64) {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.JudgementCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateJudgement creates a judgement
func (k Keeper) CreateJudgement(ctx sdk.Context, msg types.MsgCreateJudgement) {
	// Create the judgement
	count := k.GetJudgementCount(ctx)
	var judgement = types.Judgement{
		Creator:       msg.Creator,
		ID:            strconv.FormatInt(count, 10),
		ChargeSheetID: msg.ChargeSheetID,
		CourtID:       msg.CourtID,
		Content:       msg.Content,
		Complete:      false,
	}

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.JudgementPrefix + judgement.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(judgement)
	store.Set(key, value)

	// Update judgement count
	k.SetJudgementCount(ctx, count+1)
}

// GetJudgement returns the judgement information
func (k Keeper) GetJudgement(ctx sdk.Context, key string) (types.Judgement, error) {
	store := ctx.KVStore(k.storeKey)
	var judgement types.Judgement
	byteKey := []byte(types.JudgementPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &judgement)
	if err != nil {
		return judgement, err
	}
	return judgement, nil
}

// SetJudgement sets a judgement
func (k Keeper) SetJudgement(ctx sdk.Context, judgement types.Judgement) {
	judgementKey := judgement.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(judgement)
	key := []byte(types.JudgementPrefix + judgementKey)
	store.Set(key, bz)
}

// DeleteJudgement deletes a judgement
func (k Keeper) DeleteJudgement(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.JudgementPrefix + key))
}

//
// Functions used by querier
//

func listJudgement(ctx sdk.Context, k Keeper) ([]byte, error) {
	var judgementList []types.Judgement
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.JudgementPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var judgement types.Judgement
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &judgement)
		judgementList = append(judgementList, judgement)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, judgementList)
	return res, nil
}

func getJudgement(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	judgement, err := k.GetJudgement(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, judgement)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetJudgementOwner(ctx sdk.Context, key string) sdk.AccAddress {
	judgement, err := k.GetJudgement(ctx, key)
	if err != nil {
		return nil
	}
	return judgement.Creator
}

// JudgementExists check if the key exists in the store
func (k Keeper) JudgementExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.JudgementPrefix + key))
}
