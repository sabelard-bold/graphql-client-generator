package graphql

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	http.Client
}

type Request struct {
	URL       *url.URL
	Header    http.Header
	Query     string
	Variables interface{}
}

func (c *Client) Query(ctx context.Context, request Request) ([]byte, error) {
	body, err := json.Marshal(map[string]interface{}{
		"query":     request.Query,
		"variables": request.Variables,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, request.URL.String(), bytes.NewReader(body))
	req.Header = request.Header

	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
