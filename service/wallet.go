package service

import (
	"blockchain-core/types"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"

	"github.com/hacpy/go-ethereum/common/hexutil"
	"github.com/hacpy/go-ethereum/crypto"
)

func newKeyPair() (string, string, error) {
	// y^2 = x^3 + ax + b
	// a = 0, b = 7
	p256 := elliptic.P256()

	if private, err := ecdsa.GenerateKey(p256, rand.Reader); err != nil {
		return "", "", err
	} else if private == nil {
		return "", "", errors.New("pk is nil")
	} else {
		privateKeyBytes := crypto.FromECDSA(private)
		privateKey := hexutil.Encode(privateKeyBytes)

		againPrivateKey, err := crypto.HexToECDSA(privateKey[2:])
		if err != nil {
			return "", "", err
		}

		cPublicKey := againPrivateKey.Public()
		publicKeyECDSA, ok := cPublicKey.(*ecdsa.PublicKey)

		if !ok {
			return "", "", errors.New("error casting public key type")
		}

		publicKey := crypto.PubkeyToAddress(*publicKeyECDSA)

		return privateKey, hexutil.Encode(publicKey[:]), nil
	}

}

func (s *Service) MakeWallet() *types.Wallet {
	var wallet types.Wallet
	var err error

	if wallet.PrivateKey, wallet.PublicKey, err = newKeyPair(); err != nil {
		return nil
	} else if err = s.repository.CreateNewWallet(&wallet); err != nil {
		return nil
	} else {
		return &wallet
	}
}

func (s *Service) GetWallet(pk string) (*types.Wallet, error) {
	if wallet, err := s.repository.GetWallet(pk); err != nil {
		return nil, err
	} else {
		return wallet, nil
	}
}
