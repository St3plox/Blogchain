package auth_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/St3plox/Blogchain/foundation/blockchain/auth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAuthIntegration(t *testing.T) {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Provide a valid private key for testing
	privateKeyHex := "0x4c0883a6910395b8e7e627a899ede44db0fb6186760b6d4a3e4d3a3d9c7e10e3"

	// Create an Admin instance
	admin, err := auth.NewAdmin(privateKeyHex, client)
	assert.NoError(t, err)

	// Generate the auth object
	authOpts, err := admin.GenerateAuth(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, authOpts)

	// Verify the auth object has the correct settings
	assert.Equal(t, big.NewInt(1000000000), authOpts.GasPrice)
	assert.Equal(t, uint64(3000000), authOpts.GasLimit)
	assert.NotNil(t, authOpts.From) // Ensure the From address is set
}
