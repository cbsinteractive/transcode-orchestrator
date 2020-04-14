package rest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type RestClient struct {
	Client        *http.Client
	Configuration *Configuration
}

func NewRestClient(configuration *Configuration) *RestClient {
	return &RestClient{
		Configuration: configuration,
		Client:        &http.Client{},
	}
}

func (rc *RestClient) Patch(url string, jsonValue []byte, headers http.Header) (*http.Response, error) {
  return rc.request(http.MethodPatch, url, bytes.NewBuffer(jsonValue), headers)
}

func (rc *RestClient) Post(url string, jsonValue []byte, headers http.Header) (*http.Response, error) {
	return rc.request(http.MethodPost, url, bytes.NewBuffer(jsonValue), headers)
}

func (rc *RestClient) Put(url string, jsonValue []byte, headers http.Header) (*http.Response, error) {
	return rc.request(http.MethodPut, url, bytes.NewBuffer(jsonValue), headers)
}

func (rc *RestClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return rc.request(http.MethodDelete, url, nil, headers)
}

func (rc *RestClient) Get(url string, headers http.Header) (*http.Response, error) {
	return rc.request(http.MethodGet, url, nil, headers)
}

func (rc *RestClient) request(method string, url string, body io.Reader, headers http.Header) (*http.Response, error) {

	req, err := http.NewRequest(method, url, body)
	req.Header = headers

	if err != nil {
		return nil, err
	}

	resp, err := rc.Client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
