package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewStockProduct(t *testing.T) {
	//arrange
	stock, _ := NewStock(1, "Stock")
	product, _ := NewProduct(1, "Test", ProductType{1, "Type", time.Now(), time.Now()})

	//act
	sp, err := NewStockProduct(1, *stock, *product, 2, 0)

	//assert
	assert.Nil(t, err)
	assert.NotNil(t, sp)
	assert.Equal(t, 1, sp.Id)
	assert.Equal(t, 2, sp.Factor)
	assert.NotEmpty(t, sp.CreatedAt)
	assert.NotEmpty(t, sp.UpdatedAt)
}

func TestWhenStockProductIdIsRequired(t *testing.T) {
	//arrange
	stock, _ := NewStock(1, "Stock")
	product, _ := NewProduct(1, "Test", ProductType{1, "Type", time.Now(), time.Now()})
	//act
	_, err := NewStockProduct(0, *stock, *product, 2, 0)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrIdIsRequired)
}

func TestWhenProductFromStockProductIsRequired(t *testing.T) {
	//arrange
	stock, _ := NewStock(1, "Stock")
	product, _ := NewProduct(0, "Test", ProductType{1, "Type", time.Now(), time.Now()})
	//act
	_, err := NewStockProduct(1, *stock, *product, 2, 0)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrProductIsRequired)
}

func TestWhenStockFromStockProductIsRequired(t *testing.T) {
	//arrange
	stock, _ := NewStock(0, "Stock")
	product, _ := NewProduct(1, "Test", ProductType{1, "Type", time.Now(), time.Now()})
	//act
	_, err := NewStockProduct(1, *stock, *product, 2, 0)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrStockIsRequired)
}

func TestWhenFactorIsInvalid(t *testing.T) {
	//arrange
	stock, _ := NewStock(1, "Stock")
	product, _ := NewProduct(1, "Test", ProductType{1, "Type", time.Now(), time.Now()})
	//act
	_, err := NewStockProduct(1, *stock, *product, -1, 0)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrInvalidFactor)
}

func TestWhenQuantityIsInvalid(t *testing.T) {
	//arrange
	stock, _ := NewStock(1, "Stock")
	product, _ := NewProduct(1, "Test", ProductType{1, "Type", time.Now(), time.Now()})
	//act
	_, err := NewStockProduct(1, *stock, *product, 0, -1)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrInvalidQuantity)
}

func TestWhenFactorIsZero(t *testing.T) {
	//arrange
	stock, _ := NewStock(1, "Stock")
	product, _ := NewProduct(1, "Test", ProductType{1, "Type", time.Now(), time.Now()})
	//act
	sp, err := NewStockProduct(1, *stock, *product, 0, 0)
	//assert
	assert.Nil(t, err)
	assert.Equal(t, 1, sp.Factor)
}
