package repository

import (
	"eth_bsc_multichain/internal/service/db"
	"eth_bsc_multichain/internal/service/db/redis"
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

func New(logger logur.LoggerFacade, db *db.DB, redisClient *redis.Client) Registry {
	return &impl{
		dbRepo:    NewDbImpl(logger, db),
		cacheRepo: NewCacheRepoImpl(redisClient),
	}
}
func (i impl) Database() DatabaseRepo {
	return i.dbRepo
}

func (i impl) Cache() CacheRepo {
	return i.cacheRepo
}
