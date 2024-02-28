package service

import (
	"blockchain-core/types"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) CreateBlock(txs []*types.Transaction, prevHash []byte, height int64) *types.Block {
	var pHash []byte

	if latestBlock, err := s.repository.GetLatestBlock(); err != nil {
		if err == mongo.ErrNoDocuments {
			s.log.Info("Genesis Block Will Be Created")
			newBlock := createBlockInner(txs, pHash, height)

			return newBlock

		} else {
			s.log.Crit("Failed To Get Latest Block", "err", err)
			panic(err)
		}
	} else {
		pHash = latestBlock.Hash().Bytes()

		newBlock := createBlockInner(txs, pHash, height)

		return newBlock
	}

}

func createBlockInner(txs []*types.Transaction, prevHash []byte, height int64) *types.Block {
	return &types.Block{
		Time:         time.Now().Unix(),
		Hash:         []byte{},
		Transactions: txs,
		PrevHash:     prevHash,
		Nonce:        0,
		Height:       height,
	}
}
