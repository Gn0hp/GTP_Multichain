package user_handlers

import (
	"eth_bsc_multichain/internal/repository"
	"logur.dev/logur"
	"net/http"
)

type UserHandler struct {
	logger logur.LoggerFacade
	repo   repository.Registry
}

func New(logger logur.LoggerFacade, registry repository.Registry) UserHandler {
	return UserHandler{
		logger: logger,
		repo:   registry,
	}
}

func (u *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
func (u *UserHandler) Login() {

}
