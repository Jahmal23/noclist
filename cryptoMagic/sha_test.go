package cryptoMagic

import (
	"crypto/sha256"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ToSha256(t *testing.T) {

	shaHelper := &ShaHelper{Handler: sha256.New()}

	sha, err := shaHelper.ToSha256("12345/users")

	assert.Nil(t, err)
	assert.Equal(t, "c20acb14a3d3339b9e92daebb173e41379f9f2fad4aa6a6326a696bd90c67419", sha)
}

type ErrorMockShaHandler struct { //todo - use a mock framework -- GoMock -- mockgen
}

func (msh *ErrorMockShaHandler) Write(p []byte) (n int, err error) {
	return -1, errors.New("sha failure")
}
func (msh *ErrorMockShaHandler) Sum(b []byte) []byte {
	return []byte("")
}

func Test_ToSha256_Fail(t *testing.T) {
	shaHelper := &ShaHelper{Handler: &ErrorMockShaHandler{}}

	sha, err := shaHelper.ToSha256("12345/users")

	assert.Equal(t, "", sha)
	assert.Equal(t, errors.New("sha failure"), err)
}
