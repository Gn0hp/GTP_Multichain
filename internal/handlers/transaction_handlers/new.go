package transaction_handlers

import (
	"eth_bsc_multichain/internal/repository"
	"logur.dev/logur"
)

type BscCaller struct {
}

type TransactionHandler struct {
	logger logur.LoggerFacade
	repo   repository.Registry
}

func New(logger logur.LoggerFacade, registry repository.Registry) TransactionHandler {
	return TransactionHandler{
		logger: logger,
		repo:   registry,
	}
}
