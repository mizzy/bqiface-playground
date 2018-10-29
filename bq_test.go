package bq

import (
	"context"
	"testing"
)

func TestPublishEvent(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	err = c.publishEvent(ctx, "event")
	if err != nil {
		t.Fatal(err)
	}
}

func TestPublishEventWithMock(t *testing.T) {
	c, err := NewMockClient()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	err = c.publishEvent(ctx, "event")
	if err != nil {
		t.Fatal(err)
	}
}
