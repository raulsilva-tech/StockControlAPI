package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStock(t *testing.T) {
	//arrange and act
	s, err := NewStock(1, "Stock")
	//assert
	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, 1, s.Id)
	assert.Equal(t, "Stock", s.Description)
	assert.NotEmpty(t, s.CreatedAt)
	assert.NotEmpty(t, s.UpdatedAt)
}

func TestWhenStockIdIsRequired(t *testing.T) {
	//arrange and act
	_, err := NewStock(0, "Stock")
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrIdIsRequired)
}

func TestWhenStockDescriptionIsRequired(t *testing.T) {
	//arrange and act
	_, err := NewStock(1, "")
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrDescriptionIsRequired)
}
