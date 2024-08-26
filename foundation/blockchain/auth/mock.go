package auth

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/mock"
)

// MockContractSigner is a mock implementation of the ContractSigner interface.
type MockContractSigner struct {
	mock.Mock
}

func (m *MockContractSigner) GenerateAuth(ctx context.Context) (*bind.TransactOpts, error) {
	args := m.Called(ctx)
	return args.Get(0).(*bind.TransactOpts), args.Error(1)
}
