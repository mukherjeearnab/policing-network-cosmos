package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers policingnetworkcosmos-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/policingnetworkcosmos/judgement", createJudgementHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/policingnetworkcosmos/judgement", listJudgementHandler(cliCtx, "policingnetworkcosmos")).Methods("GET")
		r.HandleFunc("/policingnetworkcosmos/judgement/{key}", getJudgementHandler(cliCtx, "policingnetworkcosmos")).Methods("GET")
		r.HandleFunc("/policingnetworkcosmos/judgement", setJudgementHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/policingnetworkcosmos/judgement", deleteJudgementHandler(cliCtx)).Methods("DELETE")

		
		r.HandleFunc("/policingnetworkcosmos/chargesheet", createChargesheetHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/policingnetworkcosmos/chargesheet", listChargesheetHandler(cliCtx, "policingnetworkcosmos")).Methods("GET")
		r.HandleFunc("/policingnetworkcosmos/chargesheet/{key}", getChargesheetHandler(cliCtx, "policingnetworkcosmos")).Methods("GET")
		r.HandleFunc("/policingnetworkcosmos/chargesheet", setChargesheetHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/policingnetworkcosmos/chargesheet", deleteChargesheetHandler(cliCtx)).Methods("DELETE")

		
		r.HandleFunc("/policingnetworkcosmos/evidence", createEvidenceHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/policingnetworkcosmos/evidence", listEvidenceHandler(cliCtx, "policingnetworkcosmos")).Methods("GET")
		r.HandleFunc("/policingnetworkcosmos/evidence/{key}", getEvidenceHandler(cliCtx, "policingnetworkcosmos")).Methods("GET")
		r.HandleFunc("/policingnetworkcosmos/evidence", setEvidenceHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/policingnetworkcosmos/evidence", deleteEvidenceHandler(cliCtx)).Methods("DELETE")

		
		r.HandleFunc("/policingnetworkcosmos/investigation", createInvestigationHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/policingnetworkcosmos/investigation", listInvestigationHandler(cliCtx, "policingnetworkcosmos")).Methods("GET")
		r.HandleFunc("/policingnetworkcosmos/investigation/{key}", getInvestigationHandler(cliCtx, "policingnetworkcosmos")).Methods("GET")
		r.HandleFunc("/policingnetworkcosmos/investigation", setInvestigationHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/policingnetworkcosmos/investigation", deleteInvestigationHandler(cliCtx)).Methods("DELETE")

		
		r.HandleFunc("/policingnetworkcosmos/fir", createFirHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/policingnetworkcosmos/fir", listFirHandler(cliCtx, "policingnetworkcosmos")).Methods("GET")
		r.HandleFunc("/policingnetworkcosmos/fir/{key}", getFirHandler(cliCtx, "policingnetworkcosmos")).Methods("GET")
		r.HandleFunc("/policingnetworkcosmos/fir", setFirHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/policingnetworkcosmos/fir", deleteFirHandler(cliCtx)).Methods("DELETE")

		
		r.HandleFunc("/policingnetworkcosmos/profile", createProfileHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/policingnetworkcosmos/profile", listProfileHandler(cliCtx, "policingnetworkcosmos")).Methods("GET")
		r.HandleFunc("/policingnetworkcosmos/profile/{key}", getProfileHandler(cliCtx, "policingnetworkcosmos")).Methods("GET")
		r.HandleFunc("/policingnetworkcosmos/profile", setProfileHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/policingnetworkcosmos/profile", deleteProfileHandler(cliCtx)).Methods("DELETE")

		
}
