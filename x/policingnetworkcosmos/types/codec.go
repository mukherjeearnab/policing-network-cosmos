package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
  // this line is used by starport scaffolding # 1
		cdc.RegisterConcrete(MsgCreateJudgement{}, "policingnetworkcosmos/CreateJudgement", nil)
		cdc.RegisterConcrete(MsgSetJudgement{}, "policingnetworkcosmos/SetJudgement", nil)
		cdc.RegisterConcrete(MsgDeleteJudgement{}, "policingnetworkcosmos/DeleteJudgement", nil)
		cdc.RegisterConcrete(MsgCreateChargesheet{}, "policingnetworkcosmos/CreateChargesheet", nil)
		cdc.RegisterConcrete(MsgSetChargesheet{}, "policingnetworkcosmos/SetChargesheet", nil)
		cdc.RegisterConcrete(MsgDeleteChargesheet{}, "policingnetworkcosmos/DeleteChargesheet", nil)
		cdc.RegisterConcrete(MsgCreateEvidence{}, "policingnetworkcosmos/CreateEvidence", nil)
		cdc.RegisterConcrete(MsgSetEvidence{}, "policingnetworkcosmos/SetEvidence", nil)
		cdc.RegisterConcrete(MsgDeleteEvidence{}, "policingnetworkcosmos/DeleteEvidence", nil)
		cdc.RegisterConcrete(MsgCreateInvestigation{}, "policingnetworkcosmos/CreateInvestigation", nil)
		cdc.RegisterConcrete(MsgSetInvestigation{}, "policingnetworkcosmos/SetInvestigation", nil)
		cdc.RegisterConcrete(MsgDeleteInvestigation{}, "policingnetworkcosmos/DeleteInvestigation", nil)
		cdc.RegisterConcrete(MsgCreateFir{}, "policingnetworkcosmos/CreateFir", nil)
		cdc.RegisterConcrete(MsgSetFir{}, "policingnetworkcosmos/SetFir", nil)
		cdc.RegisterConcrete(MsgDeleteFir{}, "policingnetworkcosmos/DeleteFir", nil)
		cdc.RegisterConcrete(MsgCreateProfile{}, "policingnetworkcosmos/CreateProfile", nil)
		cdc.RegisterConcrete(MsgSetProfile{}, "policingnetworkcosmos/SetProfile", nil)
		cdc.RegisterConcrete(MsgDeleteProfile{}, "policingnetworkcosmos/DeleteProfile", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
