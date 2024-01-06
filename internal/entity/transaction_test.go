package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func arrange() (*User, *Operation, *StockProduct, *Label) {

	u, _ := NewUser(1, "Peter", "peter@mail.com", "123")
	op, _ := NewOperation(1, "Operation")
	stock, _ := NewStock(1, "Stock")
	product, _ := NewProduct(1, "Test", ProductType{1, "Type", time.Now(), time.Now()})
	sp, _ := NewStockProduct(1, *stock, *product, 2, 0)
	l, _ := NewLabel(1, "1234", time.Now().Add(time.Hour*24), *product)

	return u, op, sp, l
}

func TestNewTransaction(t *testing.T) {
	//arrange
	u, op, sp, l := arrange()
	timeNow := time.Now()
	//act
	tr, err := NewTransaction(1, *u, *op, *sp, *l, timeNow, 1)

	//assure
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, 1, tr.Quantity)
	assert.Equal(t, timeNow, tr.PerformedAt)
	assert.Equal(t, tr.User.Id, u.Id)
	assert.Equal(t, tr.Operation.Id, op.Id)
	assert.Equal(t, tr.StockProduct.Id, sp.Id)
	assert.Equal(t, tr.Label.Id, l.Id)
}

func TestTransactionWhenIdIsRequired(t *testing.T) {
	//arrange
	u, op, sp, l := arrange()
	timeNow := time.Now()
	//act
	_, err := NewTransaction(0, *u, *op, *sp, *l, timeNow, 1)
	//assure
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrIdIsRequired)
}
func TestTransactionWhenQuantityIsInvalid(t *testing.T) {
	//arrange
	u, op, sp, l := arrange()
	timeNow := time.Now()
	//act
	_, err := NewTransaction(1, *u, *op, *sp, *l, timeNow, -1)
	//assure
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrInvalidQuantity)
}
func TestTransactionWhenTimeIsInvalid(t *testing.T) {
	//arrange
	u, op, sp, l := arrange()
	var timeNow time.Time
	//act
	_, err := NewTransaction(1, *u, *op, *sp, *l, timeNow, 1)
	//assure
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrPerformedAtIsRequired)
}

func TestTransactionWhenUserIsRequired(t *testing.T) {
	//arrange
	u, op, sp, l := arrange()
	u.Id = 0
	timeNow := time.Now()
	//act
	_, err := NewTransaction(1, *u, *op, *sp, *l, timeNow, 1)
	//assure
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrUserIsRequired)
}
