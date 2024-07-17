package contract

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type PostContract struct {
	Client   *ethclient.Client
	Contract *Contract
}

func NewPostContract(client *ethclient.Client, instance *Contract) (*PostContract, error) {

	return &PostContract{
		Client:   client,
		Contract: instance,
	}, nil

}
