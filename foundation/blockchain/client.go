package blockchain

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	netUrl string
	Client *ethclient.Client
}

func NewClient(rawurl string) (*Client, error) {

	client, err := ethclient.Dial(rawurl)
	if err != nil {
		return nil, fmt.Errorf("error setablishing connection: %e", err)
	}

	return &Client{netUrl: rawurl, Client: client}, nil
}
