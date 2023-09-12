package repository

import (
	"eth_bsc_multichain/internal/service/db"
	"logur.dev/logur"
)

type Registry interface {
	Database() DatabaseRepo
	Cache() CacheRepo
}

type impl struct {
	cacheRepo CacheRepo
	dbRepo    DatabaseRepo
}

func New(logger logur.LoggerFacade, db *db.DB) Registry {
	return &impl{
		dbRepo:    NewDbImpl(logger, db),
		cacheRepo: NewCacheImpl(),
	}
}
func (i impl) Database() DatabaseRepo {
	//TODO implement me
	panic("implement me")
}

func (i impl) Cache() CacheRepo {
	//TODO implement me
	panic("implement me")
}
