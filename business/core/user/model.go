package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name"`
	Email        string             `json:"email"`
	Roles        []Role             `json:"roles"`
	PasswordHash []byte             `json:"password_hash"`
	AddressHex   string             `json:"address_hex"`
	PublicKey    []byte             `json:"public_key"`
	PrivateKey   []byte             `json:"private_key"`
	SecretKey    []byte             `json:"secret_key"`
	DateCreated  time.Time          `json:"date_created"`
	DateUpdated  time.Time          `json:"date_updated"`
}

type NewUser struct {
	Name     string
	Email    string
	Password []byte `json:"password"`
}

type UserDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Roles       []Role    `json:"roles"`
	AddressHex  string    `json:"address_hex"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//==============================================================================

func Map(user User) UserDTO {
	return UserDTO{
		ID:          user.ID.String(),
		Name:        user.Name,
		Email:       user.Email,
		Roles:       user.Roles,
		AddressHex:  user.AddressHex,
		DateCreated: user.DateCreated,
		DateUpdated: user.DateUpdated,
	}
}
