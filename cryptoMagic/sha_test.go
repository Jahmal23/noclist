package cryptoMagic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ToSha256(t *testing.T) {
	sha, err := ToSha256("12345/users")

	assert.Nil(t, err)
	assert.Equal(t, "c20acb14a3d3339b9e92daebb173e41379f9f2fad4aa6a6326a696bd90c67419", sha)
}