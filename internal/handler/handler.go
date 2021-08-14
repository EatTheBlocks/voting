package handler

import (
	"voting-hub/internal/db"
	"voting-hub/internal/ipfs"
)

type handler struct {
	db   *db.Client
	ipfs ipfs.Service

	rpcApi           string
	multicallAddress string
	tokenAddress     string
}

func NewHandler(db *db.Client, ipfs ipfs.Service, rpcApi, multicallAddress, tokenAddress string) *handler {
	return &handler{
		db:               db,
		ipfs:             ipfs,
		rpcApi:           rpcApi,
		multicallAddress: multicallAddress,
		tokenAddress:     tokenAddress,
	}
}
