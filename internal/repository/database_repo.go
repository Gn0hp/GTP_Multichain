package repository

import (
	"eth_bsc_multichain/internal/repository/transactions"
	"eth_bsc_multichain/internal/repository/user"
	"eth_bsc_multichain/internal/service/db"
	"logur.dev/logur"
)

type DatabaseRepo interface {
	User() user.Repo
	Transaction() transactions.Repo
}

type dbImpl struct {
	user        user.Repo
	transaction transactions.Repo
}

func NewDbImpl(logger logur.LoggerFacade, db *db.DB) DatabaseRepo {
	return &dbImpl{
		user:        user.New(logger, db),
		transaction: transactions.New(logger, db),
	}
}

func (d dbImpl) User() user.Repo {
	return d.user
}
func (d dbImpl) Transaction() transactions.Repo {
	return d.transaction
}
