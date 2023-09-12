package repository

import (
	"eth_bsc_multichain/internal/repository/user"
	"eth_bsc_multichain/internal/service/db"
	"logur.dev/logur"
)

type DatabaseRepo interface {
	User() user.Repo
}

type dbImpl struct {
	user user.Repo
}

func (d dbImpl) User() user.Repo {
	return d.user
}

func NewDbImpl(logger logur.LoggerFacade, db *db.DB) DatabaseRepo {
	return &dbImpl{
		user: user.New(logger, db),
	}
}
