package main

import (
	"eth_bsc_multichain/internal"
	db2 "eth_bsc_multichain/internal/service/db"
	"eth_bsc_multichain/pkg"
	"eth_bsc_multichain/pkg/config"
	"eth_bsc_multichain/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	version    string
	commitHash string
	buildDate  string
)

func main() {
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	//ctx := context.TODO()
	v, f := viper.New(), pflag.NewFlagSet(string(pkg.APIAppName), pflag.ExitOnError)
	cfg := config.New(v, f)

	customLogger := logger.NewLogger(cfg.Log)
	logger.SetStandartLogger(customLogger)

	srv := internal.DefaultService{}
	srv.Init()

	db := db2.New(customLogger, cfg.Database)
	defer db.Close()
	customLogger.Info("Database connected")
	return nil
}
