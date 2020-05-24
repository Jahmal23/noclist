package httpMagic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

const maxDurationMS = 3000

type HttpRequest struct {
	UUID        	 string
	BaseUrl     	 string
	Function    	 string
	Username    	 string
	Password    	 string
	APIkey      	 string
	XRequestChecksum string
	ContentType 	 string
	M           	 map[string]string
	RequestBody 	 []byte
	VerifySSL   	 bool
}

//todo - this needs tests, was modified from our Json expecting library
func HttpGetRawString(req HttpRequest) (string, error) {
	fmt.Println("Making Get Request...")

	//We are not logging the httpRequest struct to avoid exposing credentials
	//Calling code can log as they feel necessary when constructing the HttpRequest struct passed in to this method.

	url := buildUrlWithParams(req.BaseUrl, req.Function, req.M)

	fmt.Println("Full Request", url)

	client := http.Client{}
	timeout := time.Duration(maxDurationMS * time.Millisecond)
	client.Timeout = timeout

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	if req.Username != "" && req.Password != "" {
		request.SetBasicAuth(req.Username, req.Password)
	}

	if req.APIkey != "" {
		request.Header.Add("x-api-key", req.APIkey)
	}

	if req.XRequestChecksum != ""  {
		request.Header.Add("X-Request-Checksum", req.XRequestChecksum)
	}

	resp, doErr := client.Do(request)

	if doErr != nil || resp == nil {
		fmt.Println("Error on dispatching request.", doErr)
		return "", doErr
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseErr := errors.New(resp.Status)
		return "", responseErr
	}

	response, err :=  ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(response), err
}

//todo - needs tests and consolidating.  Could just return the response to the caller
func HttpGetRawHeader(req HttpRequest, headerName string) (string, error) {
	fmt.Println("Making Get Request...")

	//We are not logging the httpRequest struct to avoid exposing credentials
	//Calling code can log as they feel necessary when constructing the HttpRequest struct passed in to this method.

	url := buildUrlWithParams(req.BaseUrl, req.Function, req.M)

	fmt.Println("Full Request", url)

	client := http.Client{}
	timeout := time.Duration(maxDurationMS * time.Millisecond)
	client.Timeout = timeout

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	if req.Username != "" && req.Password != "" {
		request.SetBasicAuth(req.Username, req.Password)
	}

	if req.APIkey != "" {
		request.Header.Add("x-api-key", req.APIkey)
	}

	if req.XRequestChecksum != ""  {
		request.Header.Add("X-Request-Checksum", req.XRequestChecksum)
	}

	resp, doErr := client.Do(request)

	if doErr != nil || resp == nil {
		fmt.Println("Error on dispatching request.", doErr)
		return "", doErr
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseErr := errors.New(resp.Status)
		return "", responseErr
	}

	return resp.Header.Get(headerName), nil
}

func CreateSOAPTestServer(intendedResp string, httpCodeToReturn int) (server *httptest.Server, err error) {
	intendedResponse := []byte(strings.TrimSpace(intendedResp))

	server = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(httpCodeToReturn)
		rw.Write(intendedResponse)

	}))

	return server, err
}

func CreateJSONTestServer(response interface{}, httpReturnCode int) (server *httptest.Server, err error) {
	server = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		js, e := json.Marshal(response)
		err = e

		rw.WriteHeader(httpReturnCode)
		rw.Write(js)
	}))

	return server, err
}

func buildUrlWithParams(BaseUrl, Function string, m map[string]string) (url string) {

	var sb strings.Builder

	sb.WriteString(BaseUrl)
	sb.WriteString("/")
	sb.WriteString(Function)
	sb.WriteString("?")

	for k, v := range m {
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(v)
		sb.WriteString("&")
	}

	//cuts off the final '&' as they only exist between parameters
	url = sb.String()
	url = url[:len(url)-1]
	return url

}
