package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	Name           string
	Email          string
	Roles          []Role
	HashedPassword []byte
	AddressHex     string
	PublicKey      []byte
	PrivateKey     []byte
	DateCreated    time.Time
	DateUpdated    time.Time
}

type NewUser struct {
	Name         string
	Email        string
	PasswordHash []byte `json:"password_hash"`
}
