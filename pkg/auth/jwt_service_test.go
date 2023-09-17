package auth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
	"time"
)

var id = "1"
var walletAddress = "0xFC7C98fF48Aa50D75b77A3CA3E7f528817b88255"
var urlRequest, _ = url.Parse(fmt.Sprintf("localhost:8080/v1/auth?wallet_address=%v", walletAddress))

var muxContext = http.Request{
	Method: http.MethodPost,
	URL:    urlRequest,
	Header: map[string][]string{
		"User-Agent": {"PostmanRuntime/7.29.0"},
	},
	RemoteAddr: "[::1]:60844",
}

func GetSampleJwtService() *JwtService {
	srv := newJwtService()
	srv.secretKey = "12345678"
	srv.secretKeyByte = []byte(srv.secretKey)
	return srv
}

func defaultExpired() time.Time {
	timestamp := time.Now().Unix() + time.Hour.Milliseconds()
	return time.Unix(timestamp, 0)
}

func expectToken() string {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiIweEZDN0M5OGZGNDhBYTUwRDc1Yjc3QTNDQTNFN2Y1Mjg4MTdiODgyNTUiLCJhdWQiOlsiWzo6MV06NjA4NDQiLCJQb3N0bWFuUnVudGltZS83LjI5LjAiXSwiZXhwIjoxNjUwNTM3Mzk4LCJqdGkiOiIxIn0.PAJYsb33ejQpZPWJMBGYskt9OFXIX06BKDHgeWLh1s8"
}

func TestJwtService_GenerateClaims(t *testing.T) {
	srv := GetSampleJwtService()
	claims := parseClaimString(&muxContext)
	token := srv.generateAccessToken(walletAddress, id, &muxContext, defaultExpired())
	assert.Equal(t, claims[0], muxContext.RemoteAddr)
	assert.Equal(t, claims[1], muxContext.UserAgent())
	assert.Equal(t, expectToken(), token)
}
func TestJwtService_ValidateToken(t *testing.T) {
	srv := GetSampleJwtService()
	token := srv.generateAccessToken(walletAddress, id, &muxContext, defaultExpired())
	tk, err := srv.validateToken(token, &muxContext)
	assert.NoError(t, err)
	assert.NotNil(t, tk)
}

func TestJwtService_ValidateToken_ExpectOk(t *testing.T) {
	srv := GetSampleJwtService()
	token := srv.generateAccessToken(walletAddress, id, &muxContext, defaultExpired())
	req := muxContext
	req.Header.Set("Authorization", fmt.Sprintf("Bearer1 %v", token))
	tk, err := srv.validateToken(token, &muxContext)
	assert.NoError(t, err)
	assert.NotNil(t, tk)
	assert.Equal(t, "invalid character 'i' looking for beginning of value", err.Error())
}
