package handler

import (
	"voting-hub/internal/db"
	"voting-hub/internal/ipfs"
)

type handler struct {
	db   *db.Client
	ipfs ipfs.Service
}

func NewHandler(db *db.Client, ipfs ipfs.Service) *handler {
	return &handler{
		db:   db,
		ipfs: ipfs,
	}
}
