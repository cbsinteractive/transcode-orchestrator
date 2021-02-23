package transcoding

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func encode(v interface{}) (io.Reader, error) {
	if v == nil {
		return nil, nil
	}
	b := new(bytes.Buffer)
	return b, json.NewEncoder(b).Encode(v)
}

func (c *Client) do(ctx context.Context, method string, path string, in, out interface{}) error {
	c.ensure()
	body, err := encode(in)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, method, c.Base.String()+path, body)
	if err != nil {
		return err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	// TODO(as): a short read killed accurate delivery of error messages
	// need to handle this case delicately

	if c := resp.StatusCode; c < 200 || c > 299 {
		return fmt.Errorf("http status: %d: body %q", c, string(data))
	}

	return json.Unmarshal(data, out)
}
