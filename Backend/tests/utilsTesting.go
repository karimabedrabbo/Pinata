package tests

import (
	"encoding/json"
	"fmt"
	ginjwt "github.com/appleboy/gin-jwt/v2"
	jwt "github.com/dgrijalva/jwt-go"
	fuzz "github.com/google/gofuzz"
	"github.com/karimabedrabbo/eyo/api/apputils"
	"github.com/karimabedrabbo/eyo/api/managers"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestSuite struct {
	router *managers.Router
	db *managers.Database
	auth *managers.Auth
	sanitize *managers.Sanitize
	fuzzer *fuzz.Fuzzer
}


func CheckResponseCode(t *testing.T, s *TestSuite, code int, body string, acceptable ...int) {
	matched := false
	for _, acceptCode := range acceptable {
		matched = matched || code == acceptCode
	}

	if !matched {
		msg := fmt.Sprintf("bad response code:\nErrorMessage: %v", body)
		assert.Fail(t, msg)
	}
}

func CheckResponse(t *testing.T, s *TestSuite, expected interface{}, response []byte) {
	var err error
	if expected, err = s.sanitize.SanitizeJsonObj(expected); err != nil {
		msg := fmt.Sprintf("failed to sanitize expected interface: %v", err)
		assert.Fail(t, msg)
		return
	}

	var responseInterface interface{}
	if responseInterface, err = apputils.Unmarshal(response); err != nil {
		msg := fmt.Sprintf("failed to unmarshal response into interface: %v", err)
		assert.Fail(t, msg)
		return
	}

	assert.Equal(t, expected, responseInterface)
}

func CheckResponseToken(t *testing.T, s *TestSuite, expected interface{}, response []byte) {
	var err error
	var responseMap map[string]interface{}
	if responseMap, err = BytesToMap(response); err != nil {
		msg := fmt.Sprintf("error parsing response to map: %v", err)
		assert.Fail(t, msg)
		return
	}

	if !assert.Contains(t, responseMap, "expires_at") {
		return
	}

	if !assert.Contains(t, responseMap, "token") {
		return
	}

	//token string -> parse -> map -> bytes
	var token *jwt.Token
	if token, err = s.auth.GinJwtClient.ParseTokenString(responseMap["token"].(string)); err != nil {
		msg := fmt.Sprintf("couldn't parse token:\nErrorMessage: %v", err)
		assert.Fail(t, msg)
		return
	}

	var claimsMap map[string]interface{}
	claimsMap = ginjwt.ExtractClaimsFromToken(token)

	//we don't know these fields in advance (also not in expected)
	var ok bool
	if _, ok = claimsMap["orig_iat"]; ok {
		delete(claimsMap, "orig_iat")
	}

	if _, ok = claimsMap["exp"]; ok {
		delete(claimsMap, "exp")
	}

	var claimsBytes []byte
	if claimsBytes, err = apputils.Marshal(claimsMap); err != nil {
		msg := fmt.Sprintf("couldn't marshal claims:\nErrorMessage: %v", err)
		assert.Fail(t, msg)
		return
	}

	CheckResponse(t, s, expected, claimsBytes)
}

func BytesToMap(b []byte) (map[string]interface{}, error) {
	var err error
	var bMap map[string]interface{}
	if err = json.Unmarshal(b, &bMap); err != nil {
		return nil, err
	}
	return bMap, nil
}

func BytesToSlice(b []byte) ([]map[string]interface{}, error) {
	var err error
	var bSlice []map[string]interface{}
	if err = json.Unmarshal(b, &bSlice); err != nil {
		return nil, err
	}
	return bSlice, nil
}
