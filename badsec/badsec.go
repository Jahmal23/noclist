package badsec

import (
	"encoding/json"
	"errors"
	"fmt"
	"noclist/cryptoMagic"
	"noclist/httpMagic"
	"strings"
)

//todo - comes from env file
const TokenHeaderName = "Badsec-Authentication-Token"

type BadSec struct {
	cryptoMagic.ShaHandler
}

func (bs *BadSec) GetAuthToken() (string, error) {
	//todo - url comes from env variable file
	//todo - pass in the timeout as a variable
	req := httpMagic.HttpRequest{UUID: "FOO", BaseUrl: "http://localhost:8888", Function: "auth"}

	return httpMagic.HttpGetRawHeader(req, TokenHeaderName)
}

func (bs *BadSec) GetUsers(authToken string) (string, error) {

	//todo - might want to improve this validation
	if len(authToken) ==  0 {
		return "", errors.New("invalid auth token")
	}

	shaHelper := &cryptoMagic.ShaHelper{Handler: bs.ShaHandler}

	//todo - users endpoint comes from an env
	checkSum, err := shaHelper.ToSha256(fmt.Sprintf("%s/users", authToken))

	if err != nil {
		return "", err
	}

	userreq := httpMagic.HttpRequest{UUID: "FOO", BaseUrl: "http://localhost:8888", Function: "users", XRequestChecksum: checkSum}

	return httpMagic.HttpGetRawString(userreq)
}

//todo - needs basic tests
func (bs *BadSec) GetJsonUsers() (string, error) {

	token, err := bs.GetAuthToken()

	if err != nil {
		return "", nil
	}

	users, err := bs.GetUsers(token)

	if err != nil {
		return "", nil
	}

	return bs.resultsToJson(users)
}

//todo - needs negative tests
func (bs *BadSec) resultsToJson(results string) (string, error) {

	if len(results) ==  0 {
		return "", errors.New("invalid result list")
	}

	jsoned, err := json.Marshal(strings.Split(results, "\n"))

	if err != nil {
		return "", nil
	}

	return string(jsoned), nil
}