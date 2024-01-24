package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOperation(t *testing.T) {
	//arrange and act
	o, err := NewOperation(1, "Operation")
	//assert
	assert.Nil(t, err)
	assert.NotNil(t, o)
	assert.Equal(t, 1, o.Id)
	assert.Equal(t, "Operation", o.Name)
	assert.NotEmpty(t, o.CreatedAt)
	assert.NotEmpty(t, o.UpdatedAt)
}

func TestWhenOperationIdIsRequired(t *testing.T) {
	//arrange and act
	_, err := NewOperation(0, "Operation")
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrIdIsRequired)
}

func TestWhenOperationNameIsRequired(t *testing.T) {
	//arrange and act
	_, err := NewOperation(1, "")
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrNameIsRequired)
}
