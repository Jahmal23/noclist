package main

import (
	"crypto/sha256"
	"fmt"
	"noclist/badsec"
)

func main() {
	bs := &badsec.BadSec{ShaHandler: sha256.New()}
	fmt.Println(bs.GetJsonUsers())
}

