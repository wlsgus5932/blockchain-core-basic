package repository

import (
	"blockchain-core/config"
	"context"

	"github.com/inconshreveable/log15"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client

	wallet *mongo.Collection
	tx     *mongo.Collection
	block  *mongo.Collection
	log    log15.Logger
}

func NewRepository(config *config.Config) (*Repository, error) {
	r := &Repository{
		log: log15.New("module", "repository"),
	}

	var err error
	ctx := context.Background()

	if r.client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo.Uri)); err != nil {
		r.log.Error("failed to connect to mongo", "uri:", config.Mongo.Uri)
		return nil, err
	} else if err != r.client.Ping(ctx, nil) {
		r.log.Error("failed to ping to mongo", "uri:", config.Mongo.Uri)
		return nil, err
	} else {
		db := r.client.Database(config.Mongo.DB)

		r.wallet = db.Collection("wallet")
		r.tx = db.Collection("tx")
		r.block = db.Collection("block")

		r.log.Info("success to connect repository", "uri:", config.Mongo.Uri, "db", config.Mongo.DB)
		return r, nil
	}
}
