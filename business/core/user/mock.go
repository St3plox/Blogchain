package user

import (
	"context"
	"net/mail"

	"github.com/stretchr/testify/mock"
)

// MockStorer mocks the Storer interface for testing.
type MockStorer struct {
	mock.Mock
}

func (m *MockStorer) Create(ctx context.Context, usr User) (User, error) {
	args := m.Called(ctx, usr)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockStorer) Delete(ctx context.Context, usr User) error {
	args := m.Called(ctx, usr)
	return args.Error(0)
}

func (m *MockStorer) Query(ctx context.Context, filter QueryFilter, orderBy string, pageNumber int, rowsPerPage int) ([]User, error) {
	args := m.Called(ctx, filter, orderBy, pageNumber, rowsPerPage)
	return args.Get(0).([]User), args.Error(1)
}

func (m *MockStorer) Count(ctx context.Context, filter QueryFilter) (int, error) {
	args := m.Called(ctx, filter)
	return args.Int(0), args.Error(1)
}

func (m *MockStorer) QueryByID(ctx context.Context, userID string) (User, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockStorer) QueryByIDs(ctx context.Context, userIDs []string) ([]User, error) {
	args := m.Called(ctx, userIDs)
	return args.Get(0).([]User), args.Error(1)
}

func (m *MockStorer) QueryByEmail(ctx context.Context, email mail.Address) (User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(User), args.Error(1)
}