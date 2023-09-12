package cmd

import (
	"eth_bsc_multichain/cmd/router"
	"eth_bsc_multichain/pkg"
	"eth_bsc_multichain/pkg/config"
	logger_pkg "eth_bsc_multichain/pkg/logger"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func server() {
	v, f := viper.New(), pflag.NewFlagSet(string(pkg.APIAppName), pflag.ExitOnError)
	cfg := config.New(v, f)

	// Create logger (first thing after configuration loading)
	logger := logger_pkg.NewLogger(cfg.Log)
	// Override the global standard library logger to make sure everything uses our logger
	logger_pkg.SetStandartLogger(logger)
	logger.Info("") //TODO:  build INFO

	srv := &http.Server{
		Addr:         "",
		Handler:      router.New(true, logger),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	// Create a channel to receive server error
	serverErr := make(chan error, 1)
	go func() {
		logger.Info("Running server at: ", map[string]interface{}{
			"address": "",
		})
		serverErr <- srv.ListenAndServe()
	}()

}
