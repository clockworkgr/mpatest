package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers mpa-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/mpa/positions/{id}", getPositionHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/mpa/positions", listPositionHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/mpa/positions", createPositionHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/mpa/positions/{id}", updatePositionHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/mpa/positions/{id}", deletePositionHandler(clientCtx)).Methods("POST")

}
