package bq

import (
	"context"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/google-cloud-go-testing/bigquery/bqiface"
)

type client struct{ bqiface.Client }

type event struct{ string }

func NewClient() (client, error) {
	ctx := context.Background()
	c, err := bigquery.NewClient(ctx, "my-project")
	if err != nil {
		return client{}, err
	}
	return client{bqiface.AdaptClient(c)}, nil
}

func (c client) publishEvent(ctx context.Context, e string) error {
	table := c.Client.Dataset("my-dataset").Table("my-table")
	return table.Uploader().Put(ctx, event{e})
}
