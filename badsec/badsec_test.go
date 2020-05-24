package badsec

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"noclist/cryptoMagic"
	"noclist/httpMagic"
	"testing"
)

func Test_GetAuthTokenEndpoint(t *testing.T) {
	//todo - use a test server so we are not dependent on the endpoint directly - well, it IS docker
	//server, err := CreateJSONTestServer(&struct{ Description string }{"Value"}, 200)
	//defer server.Close()

	//todo - that url should come from env variable
	req := httpMagic.HttpRequest{UUID: "FOO", BaseUrl: "http://localhost:8888", Function: "auth", M: make(map[string]string)}

	result, err := httpMagic.HttpGetRawHeader(req, "Badsec-Authentication-Token") //make a constant

	assert.Nil(t, err)
	assert.NotEmpty(t, result)
}

func Test_GetUsersEndpoint(t *testing.T) {
	//todo - use a test server so we are not dependent on the endpoint directly - well, it IS docker
	//server, err := CreateJSONTestServer(&struct{ Description string }{"Value"}, 200)
	//defer server.Close()

	//todo - that url should come from env variable
	authreq := httpMagic.HttpRequest{UUID: "FOO", BaseUrl: "http://localhost:8888", Function: "auth"}

	token, err := httpMagic.HttpGetRawHeader(authreq, "Badsec-Authentication-Token") //make a constant

	checkSum, _ :=  cryptoMagic.ToSha256(fmt.Sprintf("%s/users", token))

	userreq := httpMagic.HttpRequest{UUID: "FOO", BaseUrl: "http://localhost:8888", Function: "users", XRequestChecksum: checkSum}

	result, err := httpMagic.HttpGetRawString(userreq)

	assert.Nil(t, err)
	assert.NotEmpty(t, result)
}