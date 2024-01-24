package database

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestStockFindById(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created, _ := entity.NewStock(6, "Stock Name")

	//act
	dao := NewStockDAO(db)
	_ = dao.Create(created)
	st, err := dao.FindById(6)

	fmt.Println(st)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.NotEqual(t, "", st.Description)
	assert.NotEmpty(t, st.CreatedAt)
	assert.NotEmpty(t, st.UpdatedAt)
}

func TestStockFindAll(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	st, _ := entity.NewStock(6, "Stock Name")

	//act
	dao := NewStockDAO(db)
	_ = dao.Create(st)

	ptList, err := dao.FindAll()

	for _, st := range ptList {
		fmt.Println(st)
	}

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(ptList), 1)
}

func TestStockCreate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	st, _ := entity.NewStock(1, "Stock Name")

	//act
	dao := NewStockDAO(db)
	_ = dao.Delete(st)
	err = dao.Create(st)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	ptFound, _ := dao.FindById(1)
	assert.Equal(t, ptFound.Description, st.Description)

}

func TestStockUpdate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	st, _ := entity.NewStock(6, "Stock Name")

	//act
	dao := NewStockDAO(db)
	_ = dao.Create(st)
	st.Description = "Material Special"
	err = dao.Update(st)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	ptFound, _ := dao.FindById(6)
	fmt.Println(ptFound)
	assert.Equal(t, "Material Special", ptFound.Description)
}

func TestStockDelete(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	st, _ := entity.NewStock(6, "Stock Name")

	//act
	dao := NewStockDAO(db)
	_ = dao.Create(st)
	err = dao.Delete(st)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	ptFound, _ := dao.FindById(6)
	assert.Equal(t, 0, ptFound.Id)
}
