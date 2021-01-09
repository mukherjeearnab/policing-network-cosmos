package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// GetChargesheetCount get the total number of chargesheet
func (k Keeper) GetChargesheetCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.ChargesheetCountPrefix)
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

// SetChargesheetCount set the total number of chargesheet
func (k Keeper) SetChargesheetCount(ctx sdk.Context, count int64)  {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.ChargesheetCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateChargesheet creates a chargesheet
func (k Keeper) CreateChargesheet(ctx sdk.Context, msg types.MsgCreateChargesheet) {
	// Create the chargesheet
	count := k.GetChargesheetCount(ctx)
    var chargesheet = types.Chargesheet{
        Creator: msg.Creator,
        ID:      strconv.FormatInt(count, 10),
        ID: msg.ID,
        OfficerIDs: msg.OfficerIDs,
        FirIDs: msg.FirIDs,
        InvestigationIDs: msg.InvestigationIDs,
        Content: msg.Content,
        Complete: msg.Complete,
    }

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ChargesheetPrefix + chargesheet.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(chargesheet)
	store.Set(key, value)

	// Update chargesheet count
    k.SetChargesheetCount(ctx, count+1)
}

// GetChargesheet returns the chargesheet information
func (k Keeper) GetChargesheet(ctx sdk.Context, key string) (types.Chargesheet, error) {
	store := ctx.KVStore(k.storeKey)
	var chargesheet types.Chargesheet
	byteKey := []byte(types.ChargesheetPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &chargesheet)
	if err != nil {
		return chargesheet, err
	}
	return chargesheet, nil
}

// SetChargesheet sets a chargesheet
func (k Keeper) SetChargesheet(ctx sdk.Context, chargesheet types.Chargesheet) {
	chargesheetKey := chargesheet.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(chargesheet)
	key := []byte(types.ChargesheetPrefix + chargesheetKey)
	store.Set(key, bz)
}

// DeleteChargesheet deletes a chargesheet
func (k Keeper) DeleteChargesheet(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.ChargesheetPrefix + key))
}

//
// Functions used by querier
//

func listChargesheet(ctx sdk.Context, k Keeper) ([]byte, error) {
	var chargesheetList []types.Chargesheet
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ChargesheetPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var chargesheet types.Chargesheet
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &chargesheet)
		chargesheetList = append(chargesheetList, chargesheet)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, chargesheetList)
	return res, nil
}

func getChargesheet(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	chargesheet, err := k.GetChargesheet(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, chargesheet)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetChargesheetOwner(ctx sdk.Context, key string) sdk.AccAddress {
	chargesheet, err := k.GetChargesheet(ctx, key)
	if err != nil {
		return nil
	}
	return chargesheet.Creator
}


// Check if the key exists in the store
func (k Keeper) ChargesheetExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.ChargesheetPrefix + key))
}
