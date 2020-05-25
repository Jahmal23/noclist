package badsec

import (
	"crypto/sha256"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetAuthToken(t *testing.T) {
	//todo - a DI framework would automatically initialize the shaHandler
	bs := &BadSec{ShaHandler: sha256.New()}

	token, err := bs.GetAuthToken()

	assert.Nil(t, err)
	assert.NotEmpty(t, token) //todo - assert is actually a valid token, mock out the http call!
	//todo - also by injecting http handler, we could mock errors
}

func Test_GetUsers(t *testing.T) {
	bs := &BadSec{ShaHandler: sha256.New()}

	token, _ := bs.GetAuthToken()

	users, err := bs.GetUsers(token)

	assert.Nil(t, err)
	assert.NotEmpty(t, users) //todo - assert is actually a valid user list, mock out the http call!
}

func Test_GetUsers_Bad_Token(t *testing.T) {
	bs := &BadSec{ShaHandler: sha256.New()}

	badToken := ""

	users, err := bs.GetUsers(badToken)

	assert.Empty(t, users)
	assert.Equal(t, errors.New("invalid auth token"), err)
}

func Test_resultsToJson(t *testing.T) {
	bs := &BadSec{ShaHandler: sha256.New()}

	token, _ := bs.GetAuthToken()

	users, _ := bs.GetUsers(token)

	jsonList, err := bs.resultsToJson(users)

	assert.Nil(t, err)
	assert.NotEmpty(t, jsonList)
}