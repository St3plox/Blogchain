package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Roles          []Role    `json:"roles"`
	HashedPassword []byte    `json:"hashed_password"`
	AddressHex     string    `json:"address_hex"`
	PublicKey      []byte    `json:"public_key"`
	PrivateKey     []byte    `json:"private_key"`
	SecretKey      []byte    `json:"secret_key"`
	DateCreated    time.Time `json:"date_created"`
	DateUpdated    time.Time `json:"date_updated"`
}

type NewUser struct {
	Name         string
	Email        string
	PasswordHash []byte `json:"password_hash"`
}
