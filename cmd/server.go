package cmd

import (
	"context"
	"eth_bsc_multichain/cmd/router"
	db_service "eth_bsc_multichain/internal/service/db"
	"eth_bsc_multichain/internal/service/db/redis"
	"eth_bsc_multichain/pkg"
	"eth_bsc_multichain/pkg/build_info"
	"eth_bsc_multichain/pkg/config"
	logger_pkg "eth_bsc_multichain/pkg/logger"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	version    string
	commitHash string
	buildDate  string
)

func server() error {
	ctx := context.TODO()
	v, f := viper.New(), pflag.NewFlagSet(string(pkg.APIAppName), pflag.ExitOnError)
	cfg := config.New(v, f)

	// Create logger (first thing after configuration loading)
	logger := logger_pkg.NewLogger(cfg.Log)
	// Override the global standard library logger to make sure everything uses our logger
	logger_pkg.SetStandartLogger(logger)
	logger.Info("") //TODO:  build INFO
	addr := fmt.Sprintf(":%s", cfg.App.HttpAddress)

	// Cors
	//allowHeaders := []string{"X-Requested-With", "Content-Type", "Authorization"}
	//allowOrigins := []string{os.Getenv("ORIGIN")}
	allowHeaders := []string{"*"}
	allowOrigins := []string{"*"} // Both of these are for development only. DO NOT use in production
	allowMethods := []string{"GET", "POST", "PUT", "PATCH", "HEAD", "OPTIONS", "DELETE"}
	c := cors.New(cors.Options{
		AllowedOrigins:   allowOrigins,
		AllowedMethods:   allowMethods,
		AllowedHeaders:   allowHeaders,
		AllowCredentials: true,
	})
	// r Setup
	buildInfo := build_info.New(version, commitHash, buildDate)
	db := db_service.NewPostgresDb(logger, cfg.Database)
	redisClient, err := redis.New(ctx,
		fmt.Sprintf("%v:%v",
			viper.GetString("redis.host"),
			viper.GetString("redis.port")),
		logger)
	if err != nil {
		return err
	}
	defer db.Close()
	srv := &http.Server{
		Addr:         addr,
		Handler:      c.Handler(router.New(logger, buildInfo, db, redisClient)),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	// Create a channel to receive server error
	serverErr := make(chan error, 1)
	go func() {
		logger.Info(fmt.Sprintf("Starting server at %s", addr))
		serverErr <- srv.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	select {
	case <-shutdown:
		logger.Info("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logger.Error("Server shutdown fail")
			return srv.Close()
		}
	case err := <-serverErr:
		return errors.WithStack(err)
	}
	return nil
}
func Start() {
	_ = server()
}
