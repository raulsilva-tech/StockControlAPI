package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	//arrange
	pt, _ := NewProductType(1, "Medicamento")
	//act
	p, err := NewProduct(2, "Novalgina", *pt)
	//assert
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, 2, p.Id)
	assert.Equal(t, "Novalgina", p.Description)
	assert.Equal(t, 1, p.Type.Id)
	assert.NotEmpty(t, p.CreatedAt)
	assert.NotEmpty(t, p.UpdatedAt)
}

func TestWhenProductDescriptionIsRequired(t *testing.T) {
	//arrange
	pt, _ := NewProductType(1, "Medicamento")
	//act
	_, err := NewProduct(1, "", *pt)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrDescriptionIsRequired)
}
func TestWhenProductIdIsRequired(t *testing.T) {
	//arrange
	pt, _ := NewProductType(1, "Medicamento")
	//act
	_, err := NewProduct(0, "Teste", *pt)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrIdIsRequired)
}
func TestWhenProductTypeIdFromProductIsRequired(t *testing.T) {
	//arrange
	pt, _ := NewProductType(0, "Medicamento")
	//act
	_, err := NewProduct(1, "Teste", *pt)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrProductTypeIsRequired)
}
