package entity

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

/*
	id				BIGINT	NOT NULL,
  	name		VARCHAR(200)	NOT NULL,
	email		VARCHAR(100) NOT NULL,
	password	VARCHAR(100) NOT NULL,
	created_at		timestamp,
	updated_at		timestamp,
*/

var (
	ErrEmailIsRequired    = errors.New("email is required")
	ErrNameIsRequired     = errors.New("name is required")
	ErrPasswordIsRequired = errors.New("password is required")
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Validate() error {
	if u.Id == 0 {
		return ErrIdIsRequired
	}
	if u.Name == "" {
		return ErrNameIsRequired
	}
	if u.Email == "" {
		return ErrEmailIsRequired
	}
	if u.Password == "" {
		return ErrPasswordIsRequired
	}

	return nil
}

func NewUser(id int, name, email, password string) (*User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &User{
		Id:        id,
		Name:      name,
		Email:     email,
		Password:  string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return u, u.Validate()
}
