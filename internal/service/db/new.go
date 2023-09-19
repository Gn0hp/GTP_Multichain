package db

import (
	"eth_bsc_multichain/internal"
	postgres2 "eth_bsc_multichain/internal/service/db/postgres"
	"eth_bsc_multichain/pkg"
	"fmt"
	"gorm.io/gorm"
	"logur.dev/logur"
)

type DB struct {
	gormDb *gorm.DB
}

func NewPostgresDb(logger logur.LoggerFacade, config postgres2.Config) *DB {
	logger.Info("Connecting to database...")
	db, err := postgres2.NewConnector(config)
	if err != nil {
		panic(fmt.Sprintf("db(new): connect database failed, error: %v", err))
	}
	logger.Info("Database connected!")
	_, _ = db.DB()

	internal.RegisterApp(string(pkg.MySqlConnectorAppName), db)
	return &DB{
		gormDb: db,
	}
}

func (db *DB) GormDB() *gorm.DB {
	return db.gormDb
}
func (db *DB) Close() {
	dbConnection, _ := db.gormDb.DB()
	if dbConnection != nil {
		dbConnection.Close()
	}
}
