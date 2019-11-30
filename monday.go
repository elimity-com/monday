package monday

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const baseURL = "https://api.monday.com/v2/"

type Client struct {
	client *http.Client
	token  string
}

func NewClient(accessToken string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	return &Client{
		token:  accessToken,
		client: client,
	}
}

func (c *Client) Exec(ctx context.Context, payload Payload) (*http.Response, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var queries []string
	for _, query := range payload.queries {
		if query.str == "" {
			continue
		}
		queries = append(queries, query.str)
	}
	queryString := fmt.Sprintf("{%s}", strings.Join(queries, ""))
	req, err := http.NewRequest(http.MethodPost, baseURL, strings.NewReader((url.Values{"query": []string{queryString}}).Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
