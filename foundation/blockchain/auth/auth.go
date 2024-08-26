// Package auth provides ability to generate bind.TransactOpts using admin account
package auth

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractSigner interface{
	GenerateAuth(ctx context.Context) (*bind.TransactOpts, error) 
}

type Admin struct {
	privateKey []byte
	client     *ethclient.Client
}

func NewAdmin(privateKeyHex string, client *ethclient.Client) (*Admin, error) {
	// Convert hex string to bytes
	privateKeyBytes, err := hex.DecodeString(privateKeyHex[2:]) // Remove "0x" prefix
	if err != nil {
		return nil, fmt.Errorf("invalid private key format: %w", err)
	}

	return &Admin{
		privateKey: privateKeyBytes,
		client:     client,
	}, nil
}

func (a *Admin) GenerateAuth(ctx context.Context) (*bind.TransactOpts, error) {
	return GenerateAuth(ctx, a.client, a.privateKey)
}

// GenerateAuth generates a new keyed transactor with the given private key bytes and context
func GenerateAuth(ctx context.Context, client *ethclient.Client, privateKeyBytes []byte) (*bind.TransactOpts, error) {
	// Load private key from bytes
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %w", err)
	}

	// Fetch the chain ID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Create keyed transactor with chain ID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create keyed transactor: %w", err)
	}

	// Set additional parameters
	auth.GasPrice = big.NewInt(1000000000) // 1 Gwei
	auth.GasLimit = uint64(3000000)        // 3,000,000 gas
	auth.Nonce = nil                       // Use default nonce
	auth.Context = ctx

	return auth, nil
}
