package transcodingapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *DefaultClient) getResource(result interface{}, path string) error {
	return c.reqWithMethodAndPayload(http.MethodGet, path, result, nil)
}

func (c *DefaultClient) postResource(resource interface{}, result interface{}, path string) error {
	return c.reqWithMethodAndPayload(http.MethodPost, path, result, resource)
}

func (c *DefaultClient) removeResource(result interface{}, path string) error {
	return c.reqWithMethodAndPayload(http.MethodDelete, path, result, nil)
}

func (c *DefaultClient) reqWithMethodAndPayload(method string, path string, result interface{}, reqBody interface{}) error {
	var req *http.Request
	var err error

	if reqBody != nil {
		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(reqBody)
		if err != nil {
			return err
		}
		req, err = http.NewRequest(method, c.BaseURL.String()+path, body)
	} else {
		req, err = http.NewRequest(method, c.BaseURL.String()+path, nil)
	}

	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("recieved a non 2xx status response, got a %s with body %q", resp.Status, string(b))
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}
