package repository

type CacheRepo interface {
	Ping() error
	Get() error
	Set() error
	Delete() error
	Exists() error
	HGet() error
	HSGet() error
}

type cacheImpl struct {
}

func NewCacheImpl() CacheRepo {
	return cacheImpl{}
}
func (c cacheImpl) Ping() error {
	//TODO implement me
	panic("implement me")
}

func (c cacheImpl) Get() error {
	//TODO implement me
	panic("implement me")
}

func (c cacheImpl) Set() error {
	//TODO implement me
	panic("implement me")
}

func (c cacheImpl) Delete() error {
	//TODO implement me
	panic("implement me")
}

func (c cacheImpl) Exists() error {
	//TODO implement me
	panic("implement me")
}

func (c cacheImpl) HGet() error {
	//TODO implement me
	panic("implement me")
}

func (c cacheImpl) HSGet() error {
	//TODO implement me
	panic("implement me")
}
