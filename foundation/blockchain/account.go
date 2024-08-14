package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

// AccountData contains the details of an Ethereum account
type AccountData struct {
	PrivateKey []byte
	PublicKey  []byte
	AddressHex string
}

// NOTE: needs to be changed in future

// CreateEthAccount generates a new Ethereum account with a private key, public key, and address.
// It also checks if the generated address is available (i.e., not used for contracts or transactions).
func (c *Client) CreateEthAccount() (AccountData, error) {

	// Generate a new private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return AccountData{}, err
	}

	// Convert the private key to bytes
	privateKeyBytes := crypto.FromECDSA(privateKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return AccountData{}, err
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	available, err := c.isAvailable(address)
	if err != nil {
		return AccountData{}, err
	}
	if !available {
		return AccountData{}, errors.New("Generated address is unavailable")
	}

	return AccountData{
		PublicKey:  publicKeyBytes,
		PrivateKey: privateKeyBytes,
		AddressHex: address,
	}, nil
}

// isAvailable checks if an Ethereum address is "available" (no contract code, no transaction history)
func (c *Client) isAvailable(addressHex string) (bool, error) {

	address := common.HexToAddress(addressHex)

	// Check if the address has associated code (indicating it may be a contract)
	bytecode, err := c.Client.CodeAt(context.Background(), address, nil) // nil is the latest block
	if err != nil {
		return false, fmt.Errorf("failed to get code: %v", err)
	}

	if len(bytecode) > 0 {
		return false, nil
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}
	logs, err := c.Client.FilterLogs(context.Background(), query)
	if err != nil {
		return false, fmt.Errorf("failed to get logs: %v", err)
	}

	if len(logs) > 0 {
		fmt.Printf("Address %s has been involved in transactions\n", address.Hex())
		return false, nil
	}

	return true, nil
}
