package database

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestStockProductFindById(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	s, _ := entity.NewStock(1, "")
	pt, _ := entity.NewProductType(1, "")
	p, _ := entity.NewProduct(1, "Dexaprazol", *pt)
	created, _ := entity.NewStockProduct(1, *s, *p, 1, 0)

	//act
	dao := NewStockProductDAO(db)
	_ = dao.Create(created)

	found, err := dao.FindById(1)

	fmt.Println(p)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.Equal(t, created.Id, found.Id)
	assert.Equal(t, created.Stock.Id, found.Stock.Id)
	assert.Equal(t, created.Product.Id, found.Product.Id)
	assert.NotEmpty(t, p.CreatedAt)
	assert.NotEmpty(t, p.UpdatedAt)
}

func TestStockProductFindAll(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	s, _ := entity.NewStock(1, "")
	pt, _ := entity.NewProductType(1, "")
	p, _ := entity.NewProduct(3, "Dexaprazol", *pt)
	created, _ := entity.NewStockProduct(1, *s, *p, 1, 0)

	//act
	dao := NewStockProductDAO(db)
	_ = dao.Create(created)

	pList, err := dao.FindAll()

	for _, p := range pList {
		fmt.Println(p)
	}

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(pList), 1)
}

func TestStockProductCreate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	s, _ := entity.NewStock(1, "")
	pt, _ := entity.NewProductType(1, "")
	p, _ := entity.NewProduct(1, "Dexaprazol", *pt)
	sp, _ := entity.NewStockProduct(1, *s, *p, 1, 0)

	//act
	dao := NewStockProductDAO(db)
	_ = dao.Delete(sp)
	err = dao.Create(sp)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(1)
	assert.Equal(t, sp.Id, found.Id)
	assert.Equal(t, sp.Stock.Id, found.Stock.Id)
	assert.Equal(t, sp.Product.Id, found.Product.Id)
	assert.Equal(t, sp.Quantity, found.Quantity)
	assert.Equal(t, sp.Factor, found.Factor)

}

func TestStockProductUpdate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	s, _ := entity.NewStock(1, "")
	pt, _ := entity.NewProductType(1, "")
	p, _ := entity.NewProduct(1, "Dexaprazol", *pt)
	sp, _ := entity.NewStockProduct(1, *s, *p, 1, 0)

	//act
	dao := NewStockProductDAO(db)
	_ = dao.Create(sp)
	sp.Factor = 2
	sp.Quantity = 1
	err = dao.Update(sp)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(1)
	fmt.Println(found)
	assert.Equal(t, 2, found.Factor)
	assert.Equal(t, 1, found.Quantity)
}

func TestStockProductDelete(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	s, _ := entity.NewStock(1, "")
	pt, _ := entity.NewProductType(1, "")
	p, _ := entity.NewProduct(1, "Dexaprazol", *pt)
	sp, _ := entity.NewStockProduct(1, *s, *p, 1, 0)

	//act
	dao := NewStockProductDAO(db)
	_ = dao.Create(sp)
	err = dao.Delete(sp)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(1)
	assert.Equal(t, 0, found.Id)
}


