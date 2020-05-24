package cryptoMagic

import (
	"encoding/hex"
)

type ShaHelper struct {
	Handler ShaHandler //todo - use a DI framework -- wire??
}

func (sh *ShaHelper) ToSha256(val string) (string, error) {

	_, err := sh.Handler.Write([]byte(val))

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(sh.Handler.Sum(nil)), nil
}
