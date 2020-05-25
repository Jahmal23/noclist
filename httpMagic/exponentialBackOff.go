package httpMagic

import "time"

const WaitTimeMS = 1000
const MaxTries = 3

//todo - perhaps we could create a family of backoff handlers?
type ExpBackOff struct {
}

func (ebo *ExpBackOff) Execute(apiRequest func(req HttpRequest) (string, error), rec HttpRequest) (string, error) {

	//todo -- make flexible loop backoff N times...
	result, err := apiRequest(rec)

	if err == nil {
		return result, err
	}

	time.Sleep(WaitTimeMS * 2 * time.Millisecond)

	result, err = apiRequest(rec)

	if err == nil { //success!
		return result, err
	}

	time.Sleep(WaitTimeMS * 4 * time.Millisecond)

	//last try, whatever happens happens!
	return apiRequest(rec)

}
