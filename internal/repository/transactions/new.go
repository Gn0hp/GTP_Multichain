package transactions

import (
	"eth_bsc_multichain/internal/service/db"
	"logur.dev/logur"
)

type Repo interface {
	Find()
}

type impl struct {
	logger logur.LoggerFacade
	db     *db.DB
}

func (i impl) Find() {
	//TODO implement me
	panic("implement me")
}

func New(logger logur.LoggerFacade, db *db.DB) Repo {
	return &impl{
		logger: logger,
		db:     db,
	}
}
