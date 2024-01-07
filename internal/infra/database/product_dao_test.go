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

	//act
	dao := NewProductDAO(db)
	p, err := dao.FindById(1)

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

	//act
	dao := NewProductDAO(db)
	ptList, err := dao.FindAll()

	for _, p := range ptList {
		fmt.Println(p)
	}

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
}

func TestProductCreate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(1, "")
	p, _ := entity.NewProduct(3, "Dexaprazol", *pt)

	//act
	dao := NewProductDAO(db)
	err = dao.Create(p)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	ptFound, _ := dao.FindById(3)
	assert.Equal(t, ptFound.Description, p.Description)
	assert.Equal(t, ptFound.Type.Id, p.Type.Id)

}

func TestProductUpdate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(2, "")
	p, _ := entity.NewProduct(3, "Dexa", *pt)

	//act
	dao := NewProductDAO(db)
	err = dao.Update(p)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	ptFound, _ := dao.FindById(3)
	fmt.Println(ptFound)
	assert.Equal(t, ptFound.Description, p.Description)
	assert.Equal(t, ptFound.Type.Id, p.Type.Id)
}

func TestProductDelete(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(1, "")
	p, _ := entity.NewProduct(3, "Dexaprazol", *pt)

	//act
	dao := NewProductDAO(db)
	err = dao.Delete(p)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	ptFound, _ := dao.FindById(3)
	assert.Equal(t, 0, ptFound.Id)
}
