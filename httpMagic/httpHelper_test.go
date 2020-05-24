package httpMagic

import (
	"crypto/sha256"
	"fmt"
	"github.com/stretchr/testify/assert"
	"noclist/cryptoMagic"
	"testing"
)

func Test_GetAuthTokenEndpoint(t *testing.T) {
	//todo - use a test server so we are not dependent on the endpoint directly - well, it IS docker
	//server, err := CreateJSONTestServer(&struct{ Description string }{"Value"}, 200)
	//defer server.Close()

	//todo - that url should come from env variable
	req := HttpRequest{UUID: "FOO", BaseUrl: "http://localhost:8888", Function: "auth", M: make(map[string]string)}

	result, err := HttpGetRawHeader(req, "Badsec-Authentication-Token") //make a constant

	assert.Nil(t, err)
	assert.NotEmpty(t, result)
}

func Test_GetUsersEndpoint(t *testing.T) {
	//todo - use a test server so we are not dependent on the endpoint directly - well, it IS docker
	//server, err := CreateJSONTestServer(&struct{ Description string }{"Value"}, 200)
	//defer server.Close()

	//todo - that url should come from env variable
	authreq := HttpRequest{UUID: "FOO", BaseUrl: "http://localhost:8888", Function: "auth"}

	token, err := HttpGetRawHeader(authreq, "Badsec-Authentication-Token") //make a constant

	shaHelper := &cryptoMagic.ShaHelper{Handler: sha256.New()}

	checkSum, _ := shaHelper.ToSha256(fmt.Sprintf("%s/users", token))

	userreq := HttpRequest{UUID: "FOO", BaseUrl: "http://localhost:8888", Function: "users", XRequestChecksum: checkSum}

	result, err := HttpGetRawString(userreq)

	assert.Nil(t, err)
	assert.NotEmpty(t, result)
}
