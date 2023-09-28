package router

import (
	system_handlers "eth_bsc_multichain/internal/handlers/system_handlers"
	"eth_bsc_multichain/internal/handlers/transaction_handlers"
	user_handlers "eth_bsc_multichain/internal/handlers/user_handlers"
	"eth_bsc_multichain/internal/repository"
	database "eth_bsc_multichain/internal/service/db"
	"eth_bsc_multichain/internal/service/db/redis"
	"eth_bsc_multichain/pkg/auth"
	"eth_bsc_multichain/pkg/build_info"
	"eth_bsc_multichain/pkg/middlewares"
	"github.com/gorilla/mux"
	"logur.dev/logur"
	"net/http"
)

func New(logger logur.LoggerFacade, buildInfo build_info.BuildInfo, db *database.DB, client *redis.Client) *mux.Router {
	r := mux.NewRouter()

	r.Use(auth.Middleware(logger))
	r.Use(middlewares.Middleware(logger))

	repo := repository.New(logger, db, client)
	authService := auth.New()
	// root path start with /api/v1
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("go in Health check")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
		return
	}).Methods("GET")
	// user path start with /user
	userHandler := user_handlers.New(logger, repo)
	systemHandler := system_handlers.New(buildInfo)
	transactionHandler := transaction_handlers.New(logger, repo)

	api.HandleFunc("/liveness", systemHandler.Liveness).Methods("GET")

	user := api.PathPrefix("/auth").Subrouter()
	user.HandleFunc("/register", userHandler.Signup).Methods("POST")
	user.HandleFunc("/login", userHandler.Login(&authService)).Methods("POST")

	token := api.PathPrefix("/token").Subrouter()
	token.HandleFunc("/balance", transactionHandler.GetBalance).Methods("GET")
	return r
}
