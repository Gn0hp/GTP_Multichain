package user_handlers

import (
	"encoding/json"
	"eth_bsc_multichain/internal/repository"
	"fmt"
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
	decoder := json.NewDecoder(r.Body)
	var data interface{}
	err := decoder.Decode(&data)
	u.logger.Info("Body: ", map[string]interface{}{
		"body": data,
	})
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf("Error decoding post body: %v", err)))
	}
	_, _ = w.Write([]byte(fmt.Sprintf("Data: %v", data)))

}
func (u *UserHandler) Login() {

}
