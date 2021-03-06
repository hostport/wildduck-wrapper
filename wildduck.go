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
	"strings"
	"time"
)

var SecretKey string
var Endpoint string

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
	callRaw(method, path, secretKey string, body []byte, v interface{}) error
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
		body, err = json.Marshal(params)
		if err != nil {
			return fmt.Errorf("could not marshal: %v \n error: %v", params, err)
		}
	}

	return s.callRaw(method, path, SecretKey, body, v)
}

func (s *BackendImplementation) callRaw(method, path, secretKey string, body []byte, v interface{}) error {

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

	if err != nil {
		fmt.Printf("Request failed with error: %v", err)
		return err
	}

	resBody, err = ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		fmt.Printf("Request failed with error: %v", err)
		return err
	}

	err = json.Unmarshal(resBody, v)
	if err != nil {
		err = fmt.Errorf("could not unmarshal: %s, \n error: %v", resBody, err)
	}

	return err
}

func GetBackend() Backend {
	return newBackendImplementation(&BackendConfig{
		URL:        Endpoint,
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
