package user

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/St3plox/Blogchain/business/data/order"
	"github.com/St3plox/Blogchain/foundation/blockchain"
	"github.com/St3plox/Blogchain/foundation/cachestore"
	"github.com/St3plox/Blogchain/foundation/keystore"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound              = errors.New("user not found")
	ErrUniqueEmail           = errors.New("email is not unique")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

type Storer interface {
	Create(ctx context.Context, usr User) (User, error)
	Update(ctx context.Context, usr User) error
	Delete(ctx context.Context, usr User) error
	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]User, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)
	QueryByID(ctx context.Context, userID string) (User, error)
	QueryByIDs(ctx context.Context, userID []string) ([]User, error)
	QueryByEmail(ctx context.Context, email mail.Address) (User, error)
}

// Core manages the set of APIs for user access.
type Core struct {
	storer      Storer
	ethClient   *blockchain.Client
	cacheStorer cachestore.CacheStorer
}

// NewCore constructs a core for user api access.
func NewCore(storer Storer, ethClient *blockchain.Client, cacheStorer cachestore.CacheStorer) (*Core, error) {

	return &Core{
		storer:      storer,
		ethClient:   ethClient,
		cacheStorer: cacheStorer,
	}, nil
}

// Create inserts a new user into the database.
func (c *Core) Create(ctx context.Context, nu NewUser) (User, error) {

	account, err := c.ethClient.CreateEthAccount()
	if err != nil {
		return User{}, fmt.Errorf("create : %w", err)
	}

	roles := []Role{RoleUser, RoleAdmin}
	now := time.Now()

	secretKey, err := keystore.GenerateSecretKey()
	if err != nil {
		return User{}, fmt.Errorf("create : %w", err)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("create : %w", err)
	}

	// TODO: Private key encryption
	user := User{
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: passwordHash,
		Roles:        roles,
		DateCreated:  now,
		DateUpdated:  now,
		PublicKey:    account.PublicKey,
		PrivateKey:   account.PrivateKey,
		SecretKey:    secretKey,
		AddressHex:   account.AddressHex,
	}

	usr, err := c.storer.Create(ctx, user)
	if err != nil {
		return User{}, fmt.Errorf("create: %w", err)
	}

	if err = c.cacheStorer.Set(ctx, user); err != nil {
		return User{}, fmt.Errorf("cache set: %w", err)
	}

	return usr, nil
}

// Delete removes a user from the database.
func (c *Core) Delete(ctx context.Context, usr User) error {

	if err := c.storer.Delete(ctx, usr); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	if err := c.cacheStorer.Delete(ctx, usr.CacheKey()); err != nil {
		return fmt.Errorf("cache delete: %w", err)
	}

	return nil
}

// Query retrieves a list of existing users from the database.
func (c *Core) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]User, error) {
	users, err := c.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return users, nil
}

// Count returns the total number of users in the store.
func (c *Core) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return c.storer.Count(ctx, filter)
}

// QueryByID gets the specified user from the database.
func (c *Core) QueryByID(ctx context.Context, userID string) (User, error) {

	var user User
	err := c.cacheStorer.Get(ctx, IdToCacheKey(userID), &user)
	if err == nil {
		return user, nil
	} else if err != redis.Nil {
		return User{}, fmt.Errorf("cache get: %w", err)
	}

	user, err = c.storer.QueryByID(ctx, userID)
	if err != nil {
		return User{}, fmt.Errorf("query: userID[%s]: %w", userID, err)
	}

	return user, nil
}

// QueryByIDs gets the specified user from the database.
func (c *Core) QueryByIDs(ctx context.Context, userIDs []string) ([]User, error) {
	users, err := c.storer.QueryByIDs(ctx, userIDs)
	if err != nil {
		return nil, fmt.Errorf("query: userIDs[%s]: %w", userIDs, err)
	}

	return users, nil
}

// QueryByEmail gets the specified user from the database by email.
func (c *Core) QueryByEmail(ctx context.Context, email mail.Address) (User, error) {
	user, err := c.storer.QueryByEmail(ctx, email)
	if err != nil {
		return User{}, fmt.Errorf("query: email[%s]: %w", email, err)
	}

	return user, nil
}

// =============================================================================

// Authenticate finds a user by their email and verifies their password. On
// success it returns a Claims User representing this user. The claims can be
// used to generate a token for future authentication.
func (c *Core) Authenticate(ctx context.Context, email mail.Address, password string) (User, error) {
	usr, err := c.QueryByEmail(ctx, email)
	if err != nil {
		return User{}, fmt.Errorf("query: email[%s]: %w", email, err)
	}

	if err := bcrypt.CompareHashAndPassword(usr.PasswordHash, []byte(password)); err != nil {
		return User{}, fmt.Errorf("comparehashandpassword: %w", ErrAuthenticationFailure)
	}

	return usr, nil
}
