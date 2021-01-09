package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
)

// GetProfileCount get the total number of profile
func (k Keeper) GetProfileCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.ProfileCountPrefix)
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

// SetProfileCount set the total number of profile
func (k Keeper) SetProfileCount(ctx sdk.Context, count int64) {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.ProfileCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateProfile creates a profile
func (k Keeper) CreateProfile(ctx sdk.Context, msg types.MsgCreateProfile) {
	// Create the profile
	count := k.GetProfileCount(ctx)

	// Define Empty Slice
	var emptySlice []string

	var profile = types.Profile{
		Creator:     msg.Creator,
		ProfileType: msg.ProfileType,
		ID:          msg.ID,
		Name:        msg.Name,
		Role:        msg.Role,
		FirList:     emptySlice,
	}

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ProfilePrefix + profile.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(profile)
	store.Set(key, value)

	// Update profile count
	k.SetProfileCount(ctx, count+1)
}

// GetProfile returns the profile information
func (k Keeper) GetProfile(ctx sdk.Context, key string) (types.Profile, error) {
	store := ctx.KVStore(k.storeKey)
	var profile types.Profile
	byteKey := []byte(types.ProfilePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &profile)
	if err != nil {
		return profile, err
	}
	return profile, nil
}

// SetProfile sets a profile
func (k Keeper) SetProfile(ctx sdk.Context, profile types.Profile) {
	profileKey := profile.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(profile)
	key := []byte(types.ProfilePrefix + profileKey)
	store.Set(key, bz)
}

// DeleteProfile deletes a profile
func (k Keeper) DeleteProfile(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.ProfilePrefix + key))
}

//
// Functions used by querier
//

func listProfile(ctx sdk.Context, k Keeper) ([]byte, error) {
	var profileList []types.Profile
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ProfilePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var profile types.Profile
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &profile)
		profileList = append(profileList, profile)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, profileList)
	return res, nil
}

func getProfile(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	profile, err := k.GetProfile(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, profile)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetProfileOwner(ctx sdk.Context, key string) sdk.AccAddress {
	profile, err := k.GetProfile(ctx, key)
	if err != nil {
		return nil
	}
	return profile.Creator
}

// ProfileExists check if the key exists in the store
func (k Keeper) ProfileExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.ProfilePrefix + key))
}
