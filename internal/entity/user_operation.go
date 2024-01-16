package entity

import (
	"errors"
	"time"
)

/*
  	id				BIGINT	NOT NULL,
  	user_id			BIGINT NOT NULL,
	operation_id 	BIGINT NOT NULL,
	created_at		timestamp,
	updated_at		timestamp,
*/

var ErrOperationIsRequired = errors.New("operation is required")

type UserOperation struct {
	Id        int       `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	User      `json:"user"`
	Operation `json:"operation"`
}

func (uo *UserOperation) Validate() error {
	if uo.Id == 0 {
		return ErrIdIsRequired
	}
	if uo.User.Id == 0 {
		return ErrUserIsRequired
	}
	if uo.Operation.Id == 0 {
		return ErrOperationIsRequired
	}
	return nil
}

func NewUserOperation(id int, user User, operation Operation) (*UserOperation, error) {
	uo := &UserOperation{
		Id:        id,
		User:      user,
		Operation: operation,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return uo, uo.Validate()
}
