package database

import (
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func arrangeTransaction() *entity.Transaction {

	user, _ := entity.NewUser(1, "Transaction Name", "mail@mail.com", "1234")
	product, _ := entity.NewProduct(1, "", entity.ProductType{Id: 1})
	stockProduct, _ := entity.NewStockProduct(1, entity.Stock{Id: 1}, *product, 1, 0)
	label, _ := entity.NewLabel(3, "123", time.Now(), *product)
	operation, _ := entity.NewOperation(6, "Devolution")

	tr, _ := entity.NewTransaction(1, *user, *operation, *stockProduct, *label, time.Now(), 1)

	return tr
}

func TestTransactionFindById(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created := arrangeTransaction()

	//act
	dao := NewTransactionDAO(db)
	_ = dao.Create(created)
	found, err := dao.FindById(1)

	fmt.Println(found)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.Equal(t, created.Id, found.Id)
	assert.Equal(t, created.User.Id, found.User.Id)
	assert.Equal(t, created.Operation.Id, found.Operation.Id)
	assert.Equal(t, created.Label.Id, found.Label.Id)
	assert.Equal(t, created.Quantity, found.Quantity)
	assert.Equal(t, created.StockProduct.Id, found.StockProduct.Id)
	assert.NotEmpty(t, found.PerformedAt)

}

func TestTransactionFindAll(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created := arrangeTransaction()

	//act
	dao := NewTransactionDAO(db)
	_ = dao.Create(created)

	list, err := dao.FindAll()

	for _, user := range list {
		fmt.Println(user)
	}

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(list), 1)
}

func TestTransactionCreate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created := arrangeTransaction()

	//act
	dao := NewTransactionDAO(db)
	_ = dao.Delete(created)
	err = dao.Create(created)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(1)
	assert.Equal(t, found.Id, created.Id)

}

func TestTransactionUpdate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created := arrangeTransaction()

	//act
	dao := NewTransactionDAO(db)
	_ = dao.Create(created)
	created.Quantity = 10
	err = dao.Update(created)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(1)
	fmt.Println(found)
	assert.Equal(t, 10, found.Quantity)
}

func TestTransactionDelete(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created := arrangeTransaction()

	//act
	dao := NewTransactionDAO(db)
	_ = dao.Create(created)
	err = dao.Delete(created)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(1)
	assert.Equal(t, 0, found.Id)
}
