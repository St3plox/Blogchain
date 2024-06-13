package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type AccountData struct {
	PrivateKey []byte
	PublicKey  []byte
	AddressHex string
}

// NOTE: needs to be changed in future

func CreateEthAccount(netUrl string) (AccountData, error) {

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

	available, err := isAvailable(address, netUrl)
	if err != nil {
		return AccountData{}, err
	}
	if !available {
		return AccountData{}, errors.New("Generated address is unavaiable")
	}

	return AccountData{
		PublicKey:  publicKeyBytes,
		PrivateKey: privateKeyBytes,
		AddressHex: address,
	}, nil
}

// isAvailable checks if an Ethereum address is "available" (no contract code, no transaction history)
func isAvailable(addressHex string, netUrl string) (bool, error) {

	client, err := ethclient.Dial(netUrl)
	if err != nil {
		return false, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	address := common.HexToAddress(addressHex)

	// Check if the address has associated code (indicating it may be a contract)
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is the latest block
	if err != nil {
		return false, fmt.Errorf("failed to get code: %v", err)
	}

	if len(bytecode) > 0 {
		return false, nil
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return false, fmt.Errorf("failed to get logs: %v", err)
	}

	if len(logs) > 0 {
		fmt.Printf("Address %s has been involved in transactions\n", address.Hex())
		return false, nil
	}

	return true, nil
}
