package repository

import (
	"errors"
	"time"

	"github.com/NebilMuhe/store"
	"github.com/NebilMuhe/utils"
	"github.com/gofiber/fiber/v2/log"
)

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

type userImpl struct {
	myUsers []store.User
}

func NewUserRepository() UserRepository {
	return &userImpl{
		myUsers: []store.User{},
	}
}

// Create implements UserRepository.
func (u *userImpl) Create(user store.User) error {
	_, err := u.GetByEmail(user.Email)
	if err != nil {
		return errors.New("user already exists")
	}

	err = user.HashPassword()
	if err != nil {
		log.Error("Error hasing password", err)
		return errors.New("Something went wrong")
	}

	user.CreatedAt = time.Now().UnixMilli()
	u.myUsers = append(u.myUsers, user)
	return nil
}

// Delete implements UserRepository.
func (u *userImpl) Delete(id string) error {
	for index, user := range u.myUsers {
		if user.Id == id {
			u.myUsers = utils.RemoveIndex(u.myUsers, index)
			return nil
		}
	}
	return errors.New("can not remove user with this id")
}

// GetAll implements UserRepository.
func (u *userImpl) GetAll() []store.User {
	return u.myUsers
}

// GetOne implements UserRepository.
func (u *userImpl) GetOne(id string) (*store.User, error) {
	for _, user := range u.myUsers {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

// Update implements UserRepository.
func (u *userImpl) Update(params store.User) error {
	for index, user := range u.myUsers {
		if user.Id == params.Id {
			utils.UpdateStruct(u.myUsers[index], params)
			return nil
		}
	}

	return errors.New("user not found")
}

// GetByEmail implements UserRepository.
func (u *userImpl) GetByEmail(email string) (*store.User, error) {
	panic("unimplemented")
}

// AddAddress implements UserRepository.
func (u *userImpl) AddAddress(id string, address store.Address) error {
	panic("unimplemented")
}

// RemoveAddress implements UserRepository.
func (u *userImpl) RemoveAddress(id string, addrssId string) error {
	panic("unimplemented")
}

// UpdateAddress implements UserRepository.
func (u *userImpl) UpdateAddress(id string, address store.Address) error {
	panic("unimplemented")
}

// ValidateUserWithPassword implements UserRepository.
func (u *userImpl) ValidateUserWithPassword(username string, password string) (*store.User, error) {
	panic("unimplemented")
}
