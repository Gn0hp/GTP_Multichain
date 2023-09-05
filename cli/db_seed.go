package main

import (
	"encoding/json"
	"eth_bsc_multichain/internal"
	"eth_bsc_multichain/internal/entities"
	"eth_bsc_multichain/internal/service/db/postgres"
	"eth_bsc_multichain/pkg/logger"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"logur.dev/logur"
)

type MigrateService struct {
	internal.DefaultService
	logger logur.LoggerFacade
	gormDb *gorm.DB
}

func (s *MigrateService) Init() {
	s.DefaultService.Init()
	var (
		logCf = logger.Config{}
		dbCf  = postgres.Config{}
	)
	cfBytes, _ := json.Marshal(viper.GetStringMap("logger"))
	json.Unmarshal(cfBytes, &logCf)
	cfBytes, _ = json.Marshal(viper.GetStringMap("database"))
	json.Unmarshal(cfBytes, &dbCf)

	log := logger.NewLogger(logCf)
	if dbCf.Params == nil {
		dbCf.Params = make(map[string]string)
	}
	gormDb, err := postgres.NewConnector(dbCf)
	if err != nil {
		panic(err)
	}
	log.Info("Database connected")
	s.logger = log
	s.gormDb = gormDb
}

func main() {
	migrateService := MigrateService{}
	migrateService.Init()

	tables := []interface{}{
		entities.Token{},
	}

	err := migrateService.gormDb.AutoMigrate(tables...)
	if err != nil {
		migrateService.logger.Error(fmt.Sprintf("Seed failed, details: %v", err))
		return
	}

	migrateService.logger.Info("Seed completed")
}
