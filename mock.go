package bq

import (
	"context"
)

func NewMockClient() (client, error) {
	ctx := context.Background()
	return client{mockClient{}}, nil
}

// mockClient must implement bqiface.Client interface
type mockClient struct{}

func (m mockClient) Close() error {
	return nil
}

// ... Implment other methods of bqifacae.Client interface
