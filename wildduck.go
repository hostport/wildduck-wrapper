package wildduck

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

var SecretKey string

type APIResponse struct {
	Header     http.Header
	RawJSON    []byte
	Status     string
	StatusCode int
}

func newAPIResponse(res *http.Response, resBody []byte) *APIResponse {
	return &APIResponse{
		Header:     res.Header,
		RawJSON:    resBody,
		Status:     res.Status,
		StatusCode: res.StatusCode,
	}
}

type Backend interface {
	Call(method, path string, params interface{}, v interface{}) error
	CallRaw(method, path, secretKey string, body []byte, v interface{}) error
}

type BackendImplementation struct {
	URL        string
	HTTPClient *http.Client
}

type BackendConfig struct {
	URL        string
	HTTPClient *http.Client
}

func (s *BackendImplementation) Call(method, path string, params interface{}, v interface{}) error {

	var body []byte
	var err error

	if !(len(SecretKey) > 0) {
		return errors.New("secretKey must have a value")
	}

	if params != nil {
		// This is a little unfortunate, but Go makes it impossible to compare
		// an interface value to nil without the use of the reflect package and
		// its true disciples insist that this is a feature and not a bug.
		//
		// Here we do invoke reflect because (1) we have to reflect anyway to
		// use encode with the form package, and (2) the corresponding removal
		// of boilerplate that this enables makes the small performance penalty
		// worth it.
		reflectValue := reflect.ValueOf(params)

		if reflectValue.Kind() == reflect.Struct && !reflectValue.IsNil() {
			body, err = json.Marshal(params)
			if err != nil {
				return err
			}
		}
	}

	return s.CallRaw(method, path, SecretKey, body, v)
}

func (s *BackendImplementation) CallRaw(method, path, secretKey string, body []byte, v interface{}) error {

	req, err := s.NewRequest(method, path, secretKey, "application/json")
	if err != nil {
		return err
	}

	if err = s.Do(req, body, v); err != nil {
		return err
	}

	return nil
}

func (s *BackendImplementation) NewRequest(method, path, secretKey, contentType string) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = s.URL + path

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Access-Token", secretKey)
	req.Header.Add("Content-Type", contentType)

	return req, nil
}

type nopReadCloser struct {
	io.Reader
}

func (nopReadCloser) Close() error { return nil }

func (s *BackendImplementation) Do(req *http.Request, body []byte, v interface{}) error {
	var res *http.Response
	var err error

	var resBody []byte

	if body != nil {

		reader := bytes.NewReader(body)

		req.Body = nopReadCloser{reader}

		req.GetBody = func() (io.ReadCloser, error) {
			reader := bytes.NewReader(body)
			return nopReadCloser{reader}, nil
		}
	}

	res, err = s.HTTPClient.Do(req)

	if err == nil {
		resBody, err = ioutil.ReadAll(res.Body)
		res.Body.Close()
	}

	if err != nil {
		fmt.Printf("Request failed with error: %v", err)
	}

	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, v)

	return err
}

func GetBackend() Backend {
	return newBackendImplementation(&BackendConfig{
		URL:        "http://10.0.1.20:8080/",
		HTTPClient: httpClient,
	})
}

func newBackendImplementation(config *BackendConfig) Backend {
	return &BackendImplementation{
		HTTPClient: config.HTTPClient,
		URL:        config.URL,
	}
}

var httpClient = &http.Client{
	Timeout: 80 * time.Second,
	Transport: &http.Transport{
		TLSNextProto: make(map[string]func(string, *tls.Conn) http.RoundTripper),
	},
}
