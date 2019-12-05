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
		str := query.stringify()
		if str == "" {
			continue
		}
		queries = append(queries, str)
	}
	var mutations []string
	for _, mutation := range payload.mutations {
		str := mutation.stringify()
		if str == "" {
			continue
		}
		mutations = append(mutations, str)
	}
	var query []string
	if len(queries) != 0 {
		query = append(query, fmt.Sprintf("{%s}", strings.Join(queries, "")))
	}
	if len(mutations) != 0 {
		query = append(query, fmt.Sprintf("mutation{%s}", strings.Join(mutations, "")))
	}
	req, err := http.NewRequest(http.MethodPost, baseURL, strings.NewReader(
		(url.Values{"query": query,}).Encode(),
	))
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

type Payload struct {
	queries   []Query
	mutations []Mutation
}
