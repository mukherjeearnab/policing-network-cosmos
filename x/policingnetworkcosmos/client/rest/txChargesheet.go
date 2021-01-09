package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mukherjeearnab/policing-network-cosmos/x/policingnetworkcosmos/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createChargesheetRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	ID string `json:"ID"`
	OfficerIDs string `json:"OfficerIDs"`
	FirIDs string `json:"FirIDs"`
	InvestigationIDs string `json:"InvestigationIDs"`
	Content string `json:"Content"`
	Complete string `json:"Complete"`
	
}

func createChargesheetHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createChargesheetRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedID := req.ID
		
		parsedOfficerIDs := req.OfficerIDs
		
		parsedFirIDs := req.FirIDs
		
		parsedInvestigationIDs := req.InvestigationIDs
		
		parsedContent := req.Content
		
		parsedComplete := req.Complete
		

		msg := types.NewMsgCreateChargesheet(
			creator,
			parsedID,
			parsedOfficerIDs,
			parsedFirIDs,
			parsedInvestigationIDs,
			parsedContent,
			parsedComplete,
			
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type setChargesheetRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	ID 		string `json:"id"`
	Creator string `json:"creator"`
	ID string `json:"ID"`
	OfficerIDs string `json:"OfficerIDs"`
	FirIDs string `json:"FirIDs"`
	InvestigationIDs string `json:"InvestigationIDs"`
	Content string `json:"Content"`
	Complete string `json:"Complete"`
	
}

func setChargesheetHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setChargesheetRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedID := req.ID
		
		parsedOfficerIDs := req.OfficerIDs
		
		parsedFirIDs := req.FirIDs
		
		parsedInvestigationIDs := req.InvestigationIDs
		
		parsedContent := req.Content
		
		parsedComplete := req.Complete
		

		msg := types.NewMsgSetChargesheet(
			creator,
			req.ID,
			parsedID,
			parsedOfficerIDs,
			parsedFirIDs,
			parsedInvestigationIDs,
			parsedContent,
			parsedComplete,
			
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deleteChargesheetRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	ID 		string `json:"id"`
}

func deleteChargesheetHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deleteChargesheetRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgDeleteChargesheet(req.ID, creator)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
