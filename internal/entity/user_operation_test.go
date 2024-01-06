package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserOperation(t *testing.T) {
	//arrange
	op, _ := NewOperation(1, "Withdrawal")
	u, _ := NewUser(1, "Peter", "peter@mail.com", "123")
	//act
	uo, err := NewUserOperation(1, *u, *op)
	//assert
	assert.Nil(t, err)
	assert.NotNil(t, uo)
	assert.Equal(t, 1, uo.Id)
	assert.Equal(t, uo.User.Id, u.Id)
	assert.Equal(t, uo.Operation.Id, op.Id)
	assert.NotEmpty(t, uo.CreatedAt)
	assert.NotEmpty(t, uo.UpdatedAt)
}

func TestWhenUserOperationIdIsRequired(t *testing.T) {
	//arrange
	op, _ := NewOperation(1, "Withdrawal")
	u, _ := NewUser(1, "Peter", "peter@mail.com", "123")
	//act
	_, err := NewUserOperation(0, *u, *op)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrIdIsRequired)
}

func TestUserOperationWhenOperationIsRequired(t *testing.T) {
	//arrange
	op, _ := NewOperation(0, "Withdrawal")
	u, _ := NewUser(1, "Peter", "peter@mail.com", "123")
	//act
	_, err := NewUserOperation(1, *u, *op)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrOperationIsRequired)
}

func TestUserOperationWhenUserIsRequired(t *testing.T) {
	//arrange
	op, _ := NewOperation(1, "Withdrawal")
	u, _ := NewUser(0, "Peter", "peter@mail.com", "123")
	//act
	_, err := NewUserOperation(1, *u, *op)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrUserIsRequired)
}
