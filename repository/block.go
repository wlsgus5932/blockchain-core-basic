package repository

import (
	"context"

	"github.com/hacpy/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) GetLatestBlock() (*types.Block, error) {
	ctx := context.Background()

	var block types.Block

	opt := options.FindOne().SetSort(bson.M{"time": -1})

	if err := r.block.FindOne(ctx, bson.M{}, opt).Decode(&block); err != nil {
		return nil, err
	} else {
		return &block, nil
	}

}
