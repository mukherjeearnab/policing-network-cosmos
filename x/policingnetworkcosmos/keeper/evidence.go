package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// GetEvidenceCount get the total number of evidence
func (k Keeper) GetEvidenceCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.EvidenceCountPrefix)
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

// SetEvidenceCount set the total number of evidence
func (k Keeper) SetEvidenceCount(ctx sdk.Context, count int64)  {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.EvidenceCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateEvidence creates a evidence
func (k Keeper) CreateEvidence(ctx sdk.Context, msg types.MsgCreateEvidence) {
	// Create the evidence
	count := k.GetEvidenceCount(ctx)
    var evidence = types.Evidence{
        Creator: msg.Creator,
        ID:      strconv.FormatInt(count, 10),
        ID: msg.ID,
        FileExt: msg.FileExt,
        InvestigationID: msg.InvestigationID,
    }

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.EvidencePrefix + evidence.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(evidence)
	store.Set(key, value)

	// Update evidence count
    k.SetEvidenceCount(ctx, count+1)
}

// GetEvidence returns the evidence information
func (k Keeper) GetEvidence(ctx sdk.Context, key string) (types.Evidence, error) {
	store := ctx.KVStore(k.storeKey)
	var evidence types.Evidence
	byteKey := []byte(types.EvidencePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &evidence)
	if err != nil {
		return evidence, err
	}
	return evidence, nil
}

// SetEvidence sets a evidence
func (k Keeper) SetEvidence(ctx sdk.Context, evidence types.Evidence) {
	evidenceKey := evidence.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(evidence)
	key := []byte(types.EvidencePrefix + evidenceKey)
	store.Set(key, bz)
}

// DeleteEvidence deletes a evidence
func (k Keeper) DeleteEvidence(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.EvidencePrefix + key))
}

//
// Functions used by querier
//

func listEvidence(ctx sdk.Context, k Keeper) ([]byte, error) {
	var evidenceList []types.Evidence
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.EvidencePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var evidence types.Evidence
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &evidence)
		evidenceList = append(evidenceList, evidence)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, evidenceList)
	return res, nil
}

func getEvidence(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	evidence, err := k.GetEvidence(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, evidence)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetEvidenceOwner(ctx sdk.Context, key string) sdk.AccAddress {
	evidence, err := k.GetEvidence(ctx, key)
	if err != nil {
		return nil
	}
	return evidence.Creator
}


// Check if the key exists in the store
func (k Keeper) EvidenceExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.EvidencePrefix + key))
}
