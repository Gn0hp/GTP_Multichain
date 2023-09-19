package user

import (
	"eth_bsc_multichain/internal/entities"
	"eth_bsc_multichain/internal/service/db"
	"logur.dev/logur"
)

type Repo interface {
	Create(entity *entities.User) error

	FindAll() []entities.User

	Delete()

	FindByAddress(address string) (entities.User, error)

	//FindByOptions(options map[string]interface{}) ([]entities.User, error)
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

func (i impl) Create(user *entities.User) error {
	return i.db.GormDB().Model(entities.User{}).Create(user).Error
}

func (i impl) FindAll() []entities.User {
	var users []entities.User
	i.db.GormDB().Model(entities.User{}).Find(&users)
	return users
}

func (i impl) FindByAddress(address string) (entities.User, error) {
	var user entities.User
	tx := i.db.GormDB().Model(entities.User{}).Where("address = ?", address).Find(&user)
	return user, tx.Error
}

//func (i impl) FindByOptions(options map[string]interface{}) ([]entities.User, error) {
//	//utils.StructToMap()
//}

func (i impl) Delete() {
	//TODO implement me
	panic("implement me")
}
