package httpMagic

import (
	"crypto/sha256"
	"fmt"
	"github.com/stretchr/testify/assert"
	"noclist/cryptoMagic"
	"testing"
)

//todo - needs negative tests
func Test_HttpGetRawHeader(t *testing.T) {
	server, _ := CreateJSONTestServer(&struct{ Description string }{"OK"}, 200) //todo - mock the header to make this test real
	defer server.Close()

	//todo - that url should come from env variable, so should the token header
	req := HttpRequest{UUID: "FOO", BaseUrl: server.URL, Function: "auth", TokenHeaderName: "Badsec-Authentication-Token"}

	_, err := HttpGetRawHeader(req) //make a constant

	assert.Nil(t, err)
}

//todo - needs negative tests
func Test_GetUsersEndpoint(t *testing.T) {
	server, err := CreateJSONTestServer(&struct{ Description string }{"Value"}, 200)
	defer server.Close()

	//todo - that url should come from env variable, so should the token header
	authreq := HttpRequest{UUID: "FOO", BaseUrl: server.URL, Function: "auth", TokenHeaderName: "Badsec-Authentication-Token"}

	token, _ := HttpGetRawHeader(authreq)

	shaHelper := &cryptoMagic.ShaHelper{Handler: sha256.New()}

	checkSum, _ := shaHelper.ToSha256(fmt.Sprintf("%s/users", token))

	userreq := HttpRequest{UUID: "FOO", BaseUrl: server.URL, Function: "users", XRequestChecksum: checkSum}

	result, err := HttpGetRawString(userreq)

	assert.Nil(t, err)
	assert.NotEmpty(t, result)
}
