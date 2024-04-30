package service

import (
	"github.com/NebilMuhe/store"
	"github.com/gofiber/fiber/v2"
)

type UserUseCase interface {
	//Admin Routes
	GetAll(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	AdminUpdateAddress(c *fiber.Ctx) error
	AdminAddAddress(c *fiber.Ctx) error
	AdminRemoveAddress(c *fiber.Ctx) error
	AdminChangePassword(c *fiber.Ctx) error

	//user routes [open to all]
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	ForgotPassword(c *fiber.Ctx) error
	ChangePasswordWithToken(c *fiber.Ctx) error

	//user protected route [offer login]
	EditProfile(c *fiber.Ctx) error
	AddAddress(c *fiber.Ctx) error
	UpdateAddress(c *fiber.Ctx) error
	RemoveAddress(c *fiber.Ctx) error
	DeleteAccount(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
}

type CreateUserRequest struct {
	FullName string        `json:"full_name"`
	Email    string        `json:"email"`
	Phone    string        `json:"phone"`
	Password string        `json:"password"`
	Nominee  store.Nominee `json:"nominee"`
	Roles    []string      `json:"roles"`
	Status   string        `json:"status"`
}

type UpdateUserRequest struct {
	FullName string        `json:"full_name"`
	Email    string        `json:"email"`
	Phone    string        `json:"phone"`
	Password string        `json:"password"`
	Nominee  store.Nominee `json:"nominee"`
	Roles    []string      `json:"roles"`
	Status   string        `json:"status"`
	Deleted  bool          `json:"deleted"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRespose struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

type RegisterRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type ForgotPassword struct {
	Email string `json:"email"`
}

type ChangePasswordWithTokenRequest struct {
	Email       string `json:"email"`
	Token       string `json:"token"`
	NewPassword string `json:"password"`
}

type EditProfileRequest struct {
	FullName string        `json:"full_name"`
	Phone    string        `json:"phone"`
	Nominee  store.Nominee `json:"nominee"`
}
