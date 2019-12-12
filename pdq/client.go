package pdq

import (
	"github.com/di-wu/monday"
)

type SimpleClient struct {
	monday.Client
}

func NewSimpleClient(secret string) *SimpleClient {
	return &SimpleClient{*monday.NewClient(secret, nil)}
}
