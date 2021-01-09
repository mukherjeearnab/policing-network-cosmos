package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// GetFirCount get the total number of fir
func (k Keeper) GetFirCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.FirCountPrefix)
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

// SetFirCount set the total number of fir
func (k Keeper) SetFirCount(ctx sdk.Context, count int64)  {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.FirCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateFir creates a fir
func (k Keeper) CreateFir(ctx sdk.Context, msg types.MsgCreateFir) {
	// Create the fir
	count := k.GetFirCount(ctx)
    var fir = types.Fir{
        Creator: msg.Creator,
        ID:      strconv.FormatInt(count, 10),
        ID: msg.ID,
        CitizenID: msg.CitizenID,
        Content: msg.Content,
        InvestigationID: msg.InvestigationID,
    }

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.FirPrefix + fir.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(fir)
	store.Set(key, value)

	// Update fir count
    k.SetFirCount(ctx, count+1)
}

// GetFir returns the fir information
func (k Keeper) GetFir(ctx sdk.Context, key string) (types.Fir, error) {
	store := ctx.KVStore(k.storeKey)
	var fir types.Fir
	byteKey := []byte(types.FirPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &fir)
	if err != nil {
		return fir, err
	}
	return fir, nil
}

// SetFir sets a fir
func (k Keeper) SetFir(ctx sdk.Context, fir types.Fir) {
	firKey := fir.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(fir)
	key := []byte(types.FirPrefix + firKey)
	store.Set(key, bz)
}

// DeleteFir deletes a fir
func (k Keeper) DeleteFir(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.FirPrefix + key))
}

//
// Functions used by querier
//

func listFir(ctx sdk.Context, k Keeper) ([]byte, error) {
	var firList []types.Fir
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.FirPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var fir types.Fir
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &fir)
		firList = append(firList, fir)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, firList)
	return res, nil
}

func getFir(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	fir, err := k.GetFir(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, fir)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetFirOwner(ctx sdk.Context, key string) sdk.AccAddress {
	fir, err := k.GetFir(ctx, key)
	if err != nil {
		return nil
	}
	return fir.Creator
}


// Check if the key exists in the store
func (k Keeper) FirExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.FirPrefix + key))
}
