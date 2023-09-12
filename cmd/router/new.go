package router

import (
	user_handlers "eth_bsc_multichain/internal/handlers/user-handlers"
	"eth_bsc_multichain/pkg/auth"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"logur.dev/logur"
	"os"
)

func New(isDevEnv bool, logger logur.LoggerFacade) *mux.Router {
	r := mux.NewRouter()

	allowHeaders := []string{"X-Requested-With", "Content-Type", "Authorization"}
	allowMethods := []string{"GET", "POST", "PUT", "PATCH", "HEAD", "OPTIONS", "DELETE"}
	allowOrigins := []string{os.Getenv("ORIGIN")}
	if isDevEnv {
		allowOrigins = []string{"*"}
	}
	config := cors.Options{
		AllowedOrigins:   allowOrigins,
		AllowedMethods:   allowMethods,
		AllowedHeaders:   allowHeaders,
		AllowCredentials: true,
	}
	r.Use(cors.New(config).Handler)
	r.Use(auth.Middleware(logger))

	userHandler := user_handlers.New(logger, nil)
	// root path start with /api/v1
	api := r.PathPrefix("/api/v1").Subrouter()

	// user path start with /user
	user := api.PathPrefix("/user").Subrouter()
	user.HandleFunc("/signup", userHandler.Signup).Methods("POST")
	//token := api.PathPrefix("/token").Subrouter()

	return r
}
