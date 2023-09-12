package auth

import (
	"crypto"
	"crypto/subtle"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var secretKey string

func init() {
	if len(secretKey) == 0 {
		sha256 := crypto.SHA256.New()
		sha256.Write([]byte(fmt.Sprintf("%v_&_%v", time.Now().Unix(), time.Now().UnixNano())))
		secretKey = string(sha256.Sum(nil))
	}
}

type JwtService struct {
	secretKey     string
	secretKeyByte []byte
}

func newJwtService() *JwtService {
	return &JwtService{
		secretKey:     secretKey,
		secretKeyByte: []byte(secretKey),
	}
}

func (j *JwtService) generateAccessToken(address, id string, r *http.Request, expiredTime time.Time) string {
	aud := parseClaimString(r)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:   address,
		Audience: aud,
		ExpiresAt: &jwt.NumericDate{
			Time: expiredTime,
		},
		ID: id,
	})
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *JwtService) validateToken(tokenString string, r *http.Request) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.secretKeyByte, nil
	})
	if err != nil {
		return
	}

	if !token.Valid {
		err = ErrorTokenExpired
		return
	}

	aud := parseClaimString(r)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrorMapClaimNotFound
	}
	for _, str := range aud {
		if !verifyAudience(claims, str, true) {
			err = ErrorTokenAudience
			return
		}
	}
	return token, nil

}
func parseClaimString(r *http.Request) jwt.ClaimStrings {
	var (
		host      = r.RemoteAddr
		userAgent = r.UserAgent()
	)
	return jwt.ClaimStrings{
		host,
		userAgent,
	}
}

// --- helper

// verifyAudience convert jwt.MapClaims into []string for comparison
func verifyAudience(mapClaims jwt.MapClaims, cmp string, req bool) bool {
	var aud []string
	switch v := mapClaims["aud"].(type) {
	case string:
		aud = append(aud, v)
	case []string:
		aud = v
	case []interface{}:
		for _, a := range v {
			vs, ok := a.(string)
			if !ok {
				return false
			}
			aud = append(aud, vs)
		}
	}
	return verifyAud(aud, cmp, req)
}

// verifyAud verifies the aud claim against the cmp string
func verifyAud(aud []string, cmp string, required bool) bool {
	if len(aud) == 0 {
		return !required
	}
	// use a var here to keep constant time compare when looping over a number of claims
	result := false

	var stringClaims string
	for _, a := range aud {
		if subtle.ConstantTimeCompare([]byte(a), []byte(cmp)) != 0 {
			result = true
		}
		stringClaims = stringClaims + a
	}

	// case where "" is sent in one or many aud claims
	if len(stringClaims) == 0 {
		return !required
	}

	return result
}
