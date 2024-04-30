package repository

import "github.com/NebilMuhe/store"

type UserRepository interface {
	GetAll() []store.User
	GetOne(id string) (*store.User, error)
	GetByEmail(email string) (*store.User, error)
	Create(user store.User) error
	Delete(id string) error
	Update(params store.User) error
	ValidateUserWithPassword(username, password string) (*store.User, error)
	AddAddress(id string, address store.Address) error
	UpdateAddress(id string, address store.Address) error
	RemoveAddress(id, addrssId string) error
}
