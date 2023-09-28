package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
	"logur.dev/logur"
	"net/http"
	"strings"
	"time"
)

type Authenticator interface {
	GenerateAccessToken(address, id string, r *http.Request) string
}
type AuthenticatorImpl struct {
	JwtService
}

func (i AuthenticatorImpl) GenerateAccessToken(address, id string, r *http.Request) string {
	return i.JwtService.generateAccessToken(address, id, r, time.Now().Add(AccessTokenExpiry))
}

func New() Authenticator {
	return AuthenticatorImpl{
		JwtService: *newJwtService(),
	}
}

func Middleware(logger logur.LoggerFacade) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !skipTokenCheck(r.URL.Path) {
				accessToken := extractAccessToken(r)
				if accessToken == "" {
					logger.Info("[Auth] no access token, proceed to extract refresh token")
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				// if access token present
				t, err := newJwtService().validateToken(accessToken, r)
				if err != nil {
					logger.Error("[Auth] error when verify access token, :", ErrorToLogFields("details: ", err))
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				userDetail, err := extractTokenMetadata(t)
				if err != nil {
					logger.Error("[Auth] error when extract token metadata, :", ErrorToLogFields("details: ", err))
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				SetUserContext(r, *userDetail)
			}
			next.ServeHTTP(w, r)
		})
	}
}
func extractAccessToken(r *http.Request) string {
	// Get token from Authorization header
	bearerToken := r.Header.Get("Authorization")
	// Split the token
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 && len(strArr[1]) > 0 {
		return strArr[1]
	}
	return ""
}

func extractTokenMetadata(token *jwt.Token) (*UserDetail, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrorMapClaimNotFound
	}
	userId, ok := claims[IssuerIdClaimKey] // RegisterClaim.ID
	if !ok {
		return nil, ErrorUserIDNotFound
	}
	address, ok := claims[IssuerAddressClaimKey] // RegisterClaim.Issuer
	if !ok {
		return nil, ErrorUserAddressNotFound
	}
	return &UserDetail{
		UserID:  cast.ToInt64(userId),
		Address: cast.ToString(address),
	}, nil
}

func skipTokenCheck(url string) bool {
	//var skipPath []string
	skipPath := []string{
		"/api/v1/auth/login",
		"/api/v1/auth/register",
		"/api/v1/health-check",
		"/api/v1/liveness",
	}
	for _, str := range skipPath {
		if strings.Contains(url, str) {
			return true
		}
	}
	return false
}
