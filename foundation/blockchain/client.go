package blockchain

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Client provides methods to interact with an Ethereum blockchain network.
// It uses the Ethereum client to query information and perform blockchain operations.
type Client struct {
	netUrl string
	Client *ethclient.Client
}

// NewClient creates and returns a new Client instance connected to the specified Ethereum network URL.
func NewClient(rawurl string) (*Client, error) {

	client, err := ethclient.Dial(rawurl)
	if err != nil {
		return nil, fmt.Errorf("error setablishing connection: %e", err)
	}

	return &Client{netUrl: rawurl, Client: client}, nil
}
