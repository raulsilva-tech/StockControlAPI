package database

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestProductFindById(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(1, "")
	created, _ := entity.NewProduct(3, "Dexaprazol", *pt)

	//act
	dao := NewProductDAO(db)
	_ = dao.Create(created)

	p, err := dao.FindById(3)

	fmt.Println(p)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.NotEqual(t, "", p.Description)
	assert.NotEmpty(t, p.CreatedAt)
	assert.NotEmpty(t, p.UpdatedAt)
}

func TestProductFindAll(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(1, "")
	p, _ := entity.NewProduct(3, "Dexaprazol", *pt)

	//act
	dao := NewProductDAO(db)
	_ = dao.Create(p)

	pList, err := dao.FindAll()

	for _, p := range pList {
		fmt.Println(p)
	}

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(pList), 1)
}

func TestProductCreate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(1, "")
	p, _ := entity.NewProduct(3, "Dexaprazol", *pt)

	//act
	dao := NewProductDAO(db)
	_ = dao.Delete(p)
	err = dao.Create(p)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	ptFound, _ := dao.FindById(3)
	assert.Equal(t, ptFound.Description, p.Description)
	assert.Equal(t, ptFound.ProductType.Id, p.ProductType.Id)

}

func TestProductUpdate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(2, "")
	p, _ := entity.NewProduct(3, "Dexa", *pt)

	//act
	dao := NewProductDAO(db)
	_ = dao.Create(p)
	p.Description = "Octa"
	err = dao.Update(p)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(3)
	fmt.Println(found)
	assert.Equal(t, "Octa", found.Description)
}

func TestProductDelete(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(1, "")
	p, _ := entity.NewProduct(3, "Dexaprazol", *pt)

	//act
	dao := NewProductDAO(db)
	_ = dao.Create(p)
	err = dao.Delete(p)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	ptFound, _ := dao.FindById(3)
	assert.Equal(t, 0, ptFound.Id)
}
