package transcodingapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *DefaultClient) getResource(ctx context.Context, result interface{}, path string) error {
	return c.reqWithMethodAndPayload(ctx, http.MethodGet, path, result, nil)
}

func (c *DefaultClient) postResource(ctx context.Context, resource interface{}, result interface{}, path string) error {
	return c.reqWithMethodAndPayload(ctx, http.MethodPost, path, result, resource)
}

func (c *DefaultClient) removeResource(ctx context.Context, result interface{}, path string) error {
	return c.reqWithMethodAndPayload(ctx, http.MethodDelete, path, result, nil)
}

func (c *DefaultClient) reqWithMethodAndPayload(ctx context.Context, method string, path string, result interface{}, reqBody interface{}) error {
	var req *http.Request
	var err error

	if reqBody != nil {
		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(reqBody)
		if err != nil {
			return err
		}
		req, err = http.NewRequestWithContext(ctx, method, c.BaseURL.String()+path, body)
	} else {
		req, err = http.NewRequestWithContext(ctx, method, c.BaseURL.String()+path, nil)
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
