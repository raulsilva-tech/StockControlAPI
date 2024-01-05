package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProductType(t *testing.T) {
	//act
	pt, err := NewProductType(1, "Medicamento")
	//assert
	assert.Nil(t, err)
	assert.NotNil(t, pt)
	assert.Equal(t, 1, pt.Id)
	assert.Equal(t, "Medicamento", pt.Description)
	assert.NotEmpty(t, pt.CreatedAt)
	assert.NotEmpty(t, pt.UpdatedAt)
}
func TestWhenProductTypeDescriptionIsRequired(t *testing.T) {
	//act
	_, err := NewProductType(1, "")
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrDescriptionIsRequired)
}
func TestWhenProductTypeIdIsRequired(t *testing.T) {
	//act
	_, err := NewProductType(0, "Medicamento")
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrIdIsRequired)
}
