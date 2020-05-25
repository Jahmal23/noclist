package main

import (
	"crypto/sha256"
	"fmt"
	"noclist/badsec"
	"noclist/httpMagic"
	"os"
)

func main() {
	bs := &badsec.BadSec{ShaHandler: sha256.New(),
		BackOffHandler: &httpMagic.ExpBackOff{},
		Endpoint:       "http://localhost:8888"} //todo - endpoint from ENV

	//todo - depending on the use case, this could be a go routine, but best to keep things simple first
	list, err := bs.GetJsonUsers()

	if err != nil {
		fmt.Println(fmt.Sprintf("Could not get users due to %s", err.Error()))
		os.Exit(1)
	}

	fmt.Println(list)
}
