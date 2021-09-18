package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/osmosis-sdk/client/rest"

	"github.com/cosmos/osmosis-sdk/client"
)

// RegisterRoutes registers REST routes for the upgrade module under the path specified by routeName.
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router) {
	r := rest.WithHTTPDeprecationHeaders(rtr)
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}
