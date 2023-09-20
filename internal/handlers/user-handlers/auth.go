package user_handlers

import (
	"encoding/json"
	"eth_bsc_multichain/internal/entities"
	"eth_bsc_multichain/internal/service/db/postgres"
	"eth_bsc_multichain/pkg/auth"
	"eth_bsc_multichain/pkg/middlewares"
	"fmt"
	"net/http"
	"strconv"
)

func (u *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user entities.UserDto
	err := decoder.Decode(&user)
	if err != nil {
		u.logger.Error(fmt.Sprintf("[User_Signup] Error decoding post body: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf("Error decoding post body: %v", err)))
		return
	}
	if !user.IsValid() {
		u.logger.Error(fmt.Sprintf("[User_Signup] Invalid userDTO value: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf("Error validating post body: %v", err)))
		return
	}
	iEntity, err := user.ToEntity()
	if err != nil {
		u.logger.Error(fmt.Sprintf("[User_Signup] Error converting userDTO to entity: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("Error converting userDTO to entity: %v", err)))
		return
	}
	entity, _ := iEntity.(*entities.User)
	err = u.repo.Database().User().Create(entity)
	if err != nil {
		u.logger.Error(fmt.Sprintf("[User_Signup] Error inserting into database: %v", err))
		if postgres.IsDuplicateErr(err) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Already registered!")))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("Error inserting into database: %v", err)))
		return
	}
	//w.WriteHeader(http.StatusOK)
	//_, _ = w.Write([]byte(fmt.Sprintf("Data: %v", user)))
	middlewares.WriteResponse(&w, http.StatusOK, fmt.Sprintf("Data: %v", user))
}

func (u *UserHandler) Login(authSrv *auth.Authenticator) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user entities.UserDto
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			u.logger.Error(fmt.Sprintf("[User_Login] Error decoding post body: %v", err))
			middlewares.WriteResponse(&w, http.StatusBadRequest, fmt.Sprintf("[User-Login] Error decoding post body: %v", err))
			return
		}
		iUser, err := user.ToEntity()
		if err != nil {
			u.logger.Error(fmt.Sprintf("[User_Login] Error converting userDTO to entity: %v", err))
			middlewares.WriteResponse(&w, http.StatusInternalServerError, fmt.Sprintf("[User_Login] Error converting userDTO to entity: %v", err))
			return
		}
		entity, _ := iUser.(*entities.User)
		userExisted, err := u.repo.Database().User().FindByAddress(entity.Address)
		if err != nil {
			u.logger.Error(fmt.Sprintf("[User_Login] Error finding user: %v", err))
			middlewares.WriteResponse(&w, http.StatusInternalServerError, fmt.Sprintf("[User_Login] Error finding user: %v", err))
			return
		}
		if userExisted.ID == 0 {
			u.logger.Error(fmt.Sprintf("[User_Login] User not found"))
			middlewares.WriteResponse(&w, http.StatusNotFound, fmt.Sprintf("[User_Login] User not found"))
			return
		}
		err = u.repo.Cache().Set(r.Context(),
			fmt.Sprintf("user:%v", userExisted.ID),
			userExisted,
			auth.AccessTokenExpiry)
		if err != nil {
			u.logger.Error(fmt.Sprintf("[User_Login] Error setting user to redis cache: %v", err))
			middlewares.WriteResponse(&w, http.StatusInternalServerError, fmt.Sprintf("[User_Login] Error setting user to cache: %v", err))
			return
		}
		accessToken := (*authSrv).GenerateAccessToken(userExisted.Address, strconv.FormatUint(userExisted.ID, 10), r)
		//middlewares.WriteResponse(&w, http.StatusOK, `{ "access_token": "`+accessToken+`" }`)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{ "access_token": "` + accessToken + `" }`))
	}
}
