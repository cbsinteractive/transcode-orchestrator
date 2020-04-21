package transcoding

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

type testReqBody struct {
	SomeReqProp propType `json:"some_req_prop"`
}

type testResp struct {
	SomeProp propType `json:"some_prop"`
}

type propType struct {
	Name string `json:"name"`
}

type RespAssertion func(resp testResp) error

func Test_reqWithMethodAndPayload(t *testing.T) {
	assert := func(fns ...RespAssertion) []RespAssertion { return fns }

	respContentsAreExactly := func(want testResp) RespAssertion {
		return func(got testResp) error {
			if !reflect.DeepEqual(got, want) {
				return fmt.Errorf("got %v, expected %v", got, want)
			}
			return nil
		}
	}

	respIsSuccess := func() RespAssertion {
		return func(got testResp) error {
			if got.SomeProp.Name != "success" {
				return fmt.Errorf("expected %v to be a successful response", got)
			}
			return nil
		}
	}

	tests := []struct {
		title      string
		backend    http.HandlerFunc
		method     string
		reqBody    testReqBody
		path       string
		assertions []RespAssertion
	}{
		{
			title: "client marshalls responses correctly",
			backend: func(w http.ResponseWriter, r *http.Request) {
				writeProp(w, "test_name")
			},
			method: http.MethodGet,
			assertions: assert(respContentsAreExactly(testResp{
				SomeProp: propType{
					Name: "test_name",
				},
			})),
		},
		{
			title: "client sends requests to the correct path",
			backend: func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/test_path" {
					writeSuccess(w)
				}
			},
			path:       "/test_path",
			method:     http.MethodGet,
			assertions: assert(respIsSuccess()),
		},
		{
			title: "client sends request body correctly",
			backend: func(w http.ResponseWriter, r *http.Request) {
				reqBody := testReqBody{}
				err := json.NewDecoder(r.Body).Decode(&reqBody)
				if err == nil && reqBody.SomeReqProp.Name == "req_body" {
					writeSuccess(w)
				}
			},
			reqBody: testReqBody{
				SomeReqProp: propType{
					Name: "req_body",
				},
			},
			assertions: assert(respIsSuccess()),
		},
		{
			title: "client uses the correct HTTP method",
			backend: func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodPost {
					writeSuccess(w)
				}
			},
			method:     http.MethodPost,
			assertions: assert(respIsSuccess()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			backend := httptest.NewServer(http.HandlerFunc(tt.backend))
			defer backend.Close()
			backendURL, err := url.Parse(backend.URL)
			if err != nil {
				t.Error(err)
			}

			client := DefaultClient{BaseURL: backendURL}
			client.ensure()

			respObj := testResp{}
			err = client.reqWithMethodAndPayload(context.Background(), tt.method, tt.path, &respObj, tt.reqBody)
			if err != nil {
				t.Error(err)
			}

			for _, asrt := range tt.assertions {
				if err := asrt(respObj); err != nil {
					t.Error(err)
				}
			}
		})
	}
}

func writeSuccess(w http.ResponseWriter) {
	_, _ = w.Write([]byte(`{ "some_prop": { "name": "success"} }`))
}

func writeProp(w http.ResponseWriter, prop string) {
	_, _ = w.Write([]byte(`{ "some_prop": { "name": "` + prop + `"} }`))
}
