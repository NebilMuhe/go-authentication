package store

import "golang.org/x/crypto/bcrypt"

type Address struct {
	Id      string `json:"id"`
	Tag     string `json:"tag"`
	Country string `json:"country"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
}

type Nominee struct {
	FullName string  `json:"full_name"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Address  Address `json:"address"`
	Relation string  `json:"relation"`
}

type User struct {
	Id             string    `json:"id"`
	FullName       string    `json:"full_name"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Address        []Address `json:"address"`
	Nominee        Nominee   `json:"nominee"`
	Roles          string    `json:"roles"`
	Status         string    `json:"status"`
	Deleted        bool      `json:"deleted"`
	CreatedAt      int64     `json:"created_at"`
	UpdatedAt      int64     `json:"updated_at"`
	AdminUpdatedAt int64     `json:"admin_updated_at"`
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

func (u *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
