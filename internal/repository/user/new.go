package user

import (
	"eth_bsc_multichain/internal/service/db"
	"logur.dev/logur"
)

type Repo interface {
	Create()

	Find()

	Delete()
}

type impl struct {
	logger logur.LoggerFacade
	db     *db.DB
}

func New(logger logur.LoggerFacade, db *db.DB) Repo {
	return &impl{
		logger: logger,
		db:     db,
	}
}

func (i impl) Create() {
	//TODO implement me
	panic("implement me")
}

func (i impl) Find() {
	//TODO implement me
	panic("implement me")
}

func (i impl) Delete() {
	//TODO implement me
	panic("implement me")
}
