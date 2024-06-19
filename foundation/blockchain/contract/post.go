package contract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PostContract struct {
	client   *ethclient.Client
	Contract *Contract
}

func NewPostContract(client *ethclient.Client, contractAddress string) (*PostContract, error) {

	address := common.HexToAddress(contractAddress)
	instance, err := NewContract(address, client)
	if err != nil {
		return nil, err
	}

	return &PostContract{
		client:   client,
		Contract: instance,
	}, nil
}
