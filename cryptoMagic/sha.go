package cryptoMagic

import (
	"crypto/sha256"
	"encoding/hex"
)

//todo - can't test the error condition, need to make a simple shaHandler interface and inject to make testable

func ToSha256(val string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(val))

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
