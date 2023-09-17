package router

import (
	user_handlers "eth_bsc_multichain/internal/handlers/user-handlers"
	"eth_bsc_multichain/pkg/auth"
	"eth_bsc_multichain/pkg/middlewares"
	"github.com/gorilla/mux"
	"logur.dev/logur"
	"net/http"
)

func New(logger logur.LoggerFacade) *mux.Router {
	r := mux.NewRouter()

	r.Use(auth.Middleware(logger))
	r.Use(middlewares.Middleware(logger))

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
	userHandler := user_handlers.New(logger, nil)
	user := api.PathPrefix("/auth").Subrouter()

	user.HandleFunc("/register", userHandler.Signup).Methods("POST")

	//token := api.PathPrefix("/token").Subrouter()

	return r
}
