package blockchain

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockchainClient interface {
	CreateEthAccount() (AccountData, error)
}

// EthClient provides methods to interact with an Ethereum blockchain network.
// It uses the Ethereum client to query information and perform blockchain operations.
type EthClient struct {
	netUrl string
	Client *ethclient.Client
}

// NewClient creates and returns a new Client instance connected to the specified Ethereum network URL.
func NewClient(rawurl string) (*EthClient, error) {

	client, err := ethclient.Dial(rawurl)
	if err != nil {
		return nil, fmt.Errorf("error establishing connection: %e", err)
	}

	return &EthClient{netUrl: rawurl, Client: client}, nil
}
