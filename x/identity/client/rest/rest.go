package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/accounts/{address}", storeName), getHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/accounts/{address}", storeName), setHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/accounts/{address}/import", storeName), importHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/accounts/{address}", storeName), deleteHandler(cliCtx)).Methods("DELETE")
}
