package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/identity/accounts/{address}", getHandler(cliCtx)).Methods("GET")
	r.HandleFunc("/identity/accounts/{address}", setHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/identity/accounts/{address}/import", importHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/identity/accounts/{address}", deleteHandler(cliCtx)).Methods("DELETE")
}
