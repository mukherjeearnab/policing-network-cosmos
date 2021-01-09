package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// GetInvestigationCount get the total number of investigation
func (k Keeper) GetInvestigationCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.InvestigationCountPrefix)
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

// SetInvestigationCount set the total number of investigation
func (k Keeper) SetInvestigationCount(ctx sdk.Context, count int64)  {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.InvestigationCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateInvestigation creates a investigation
func (k Keeper) CreateInvestigation(ctx sdk.Context, msg types.MsgCreateInvestigation) {
	// Create the investigation
	count := k.GetInvestigationCount(ctx)
    var investigation = types.Investigation{
        Creator: msg.Creator,
        ID:      strconv.FormatInt(count, 10),
        ID: msg.ID,
        FirID: msg.FirID,
        OfficerID: msg.OfficerID,
        Content: msg.Content,
        Evidence: msg.Evidence,
        Complete: msg.Complete,
    }

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.InvestigationPrefix + investigation.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(investigation)
	store.Set(key, value)

	// Update investigation count
    k.SetInvestigationCount(ctx, count+1)
}

// GetInvestigation returns the investigation information
func (k Keeper) GetInvestigation(ctx sdk.Context, key string) (types.Investigation, error) {
	store := ctx.KVStore(k.storeKey)
	var investigation types.Investigation
	byteKey := []byte(types.InvestigationPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &investigation)
	if err != nil {
		return investigation, err
	}
	return investigation, nil
}

// SetInvestigation sets a investigation
func (k Keeper) SetInvestigation(ctx sdk.Context, investigation types.Investigation) {
	investigationKey := investigation.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(investigation)
	key := []byte(types.InvestigationPrefix + investigationKey)
	store.Set(key, bz)
}

// DeleteInvestigation deletes a investigation
func (k Keeper) DeleteInvestigation(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.InvestigationPrefix + key))
}

//
// Functions used by querier
//

func listInvestigation(ctx sdk.Context, k Keeper) ([]byte, error) {
	var investigationList []types.Investigation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.InvestigationPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var investigation types.Investigation
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &investigation)
		investigationList = append(investigationList, investigation)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, investigationList)
	return res, nil
}

func getInvestigation(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	investigation, err := k.GetInvestigation(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, investigation)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetInvestigationOwner(ctx sdk.Context, key string) sdk.AccAddress {
	investigation, err := k.GetInvestigation(ctx, key)
	if err != nil {
		return nil
	}
	return investigation.Creator
}


// Check if the key exists in the store
func (k Keeper) InvestigationExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.InvestigationPrefix + key))
}
