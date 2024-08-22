package blockchain

import "github.com/stretchr/testify/mock"


// MockClient is a mock implementation of the blockchain Client interface.
type MockClient struct {
	mock.Mock
}

// CreateEthAccount mocks the creation of an Ethereum account.
func (m *MockClient) CreateEthAccount() (AccountData, error) {
	args := m.Called()
	return args.Get(0).(AccountData), args.Error(1)
}

