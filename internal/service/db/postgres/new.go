package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnector(config Config) (*gorm.DB, error) {
	database, err := gorm.Open(postgres.Open(GetDSN(config)), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("connect database failed, error: %v", err))
	}
	return database, err
}
