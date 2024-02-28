package repository

import (
	"blockchain-core/types"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateNewWallet(wallet *types.Wallet) error {
	ctx := context.Background()
	wallet.Time = time.Now().Unix()

	opt := options.Update().SetUpsert(true)
	filter := bson.M{"privateKey": wallet.PrivateKey}
	update := bson.M{"$set": wallet}

	fmt.Println("wallet::", wallet)
	fmt.Println("r:::", r)
	r.wallet.UpdateOne(ctx, filter, update, opt)

	return nil
}

func (r *Repository) GetWallet(pk string) (*types.Wallet, error) {
	ctx := context.Background()

	filter := bson.M{"privateKey": pk}

	var wallet types.Wallet

	if err := r.wallet.FindOne(ctx, filter, options.FindOne()).Decode(&wallet); err != nil {
		return nil, err
	} else {
		return &wallet, nil
	}
}
