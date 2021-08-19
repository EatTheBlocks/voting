package db

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Options struct {
	URI string
}

type Client struct {
	client *mongo.Client
}

func NewClient(ctx context.Context, opts Options) (*Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(opts.URI))
	if err != nil {
		return nil, errors.Wrap(err, "mongodb connect")
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.Wrap(err, "mongodb ping")
	}

	return &Client{
		client: client,
	}, nil
}

func (c Client) Disconnect(ctx context.Context) error {
	return c.client.Disconnect(ctx)
}

func (c Client) Client() *mongo.Client {
	return c.client
}
