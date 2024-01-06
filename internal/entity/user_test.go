package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {

	//arrange act
	u, err := NewUser(1, "Peter", "peter@mail.com", "123")
	//assert
	assert.NotNil(t, u)
	assert.Nil(t, err)
	assert.Equal(t, 1, u.Id)
	assert.Equal(t, "Peter", u.Name)
	assert.Equal(t, "peter@mail.com", u.Email)
	assert.Equal(t, "123", u.Password)
	assert.NotEmpty(t, u.CreatedAt)
	assert.NotEmpty(t, u.UpdatedAt)
}

func TestUserWhenIdIsRequired(t *testing.T) {
	//arrange act
	_, err := NewUser(0, "Peter", "peter@mail.com", "123")
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrIdIsRequired)
}

func TestUserWhenNameIsRequired(t *testing.T) {
	//arrange act
	_, err := NewUser(1, "", "peter@mail.com", "123")
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrNameIsRequired)
}

func TestUserWhenEmailIsRequired(t *testing.T) {
	//arrange act
	_, err := NewUser(1, "Peter", "", "123")
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrEmailIsRequired)
}

func TestUserWhenPasswordIsRequired(t *testing.T) {
	//arrange act
	_, err := NewUser(1, "Peter", "peter@mail.com", "")
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrPasswordIsRequired)
}