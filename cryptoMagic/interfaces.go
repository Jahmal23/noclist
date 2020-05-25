package cryptoMagic

import "io"

type ShaHandler interface {
	io.Writer
	Sum(b []byte) []byte
}
