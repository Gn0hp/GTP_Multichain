package auth

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type UserDetail struct {
	UserID  int64
	Address string
}

var (
	ErrUserNotFoundInCtx = errors.New("user not found in context")
)

const AccessTokenExpiry = time.Minute // for development
//const AccessTokenExpiry = time.Hour

const (
	AccessTokenKey        = "access_token"
	IssuerAddressClaimKey = "iss"
	IssuerIdClaimKey      = "jti"
)

// SetUserContext sets user details in context
func SetUserContext(r *http.Request, userDetails UserDetail) {
	ctx := context.WithValue(r.Context(), "user", userDetails)
	r.WithContext(ctx)
}
func GetUserFromContext(r *http.Request) (UserDetail, error) {
	userDetail, ok := r.Context().Value("user").(UserDetail)
	if !ok {
		return UserDetail{}, errors.WithStack(ErrUserNotFoundInCtx)
	}
	return userDetail, nil
}
