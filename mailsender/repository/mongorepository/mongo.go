package mongorepository

import (
	"context"
	"mailsender/common"
	"mailsender/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient is a Service alias for mongo.Client.
// It implements Repository interface.
type MongoClient struct {
	*mongo.Client
}

// New establishes connection to mongo
// via config.MongoConnection and returns
// new mongo client or error if any.
func New(ctx context.Context) (repository.Repository, error) {
	mOpts := options.Client().ApplyURI(common.ServiceConfig.MongoConnection)
	mClient, err := mongo.Connect(ctx, mOpts)
	if err != nil {
		return nil, err
	}

	client := MongoClient{mClient}
	return &client, nil
}
