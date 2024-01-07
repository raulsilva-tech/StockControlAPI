package database

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func arrangeDBConnection() (*sql.DB, error) {

	//starting database connection
	DataSourceName := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", "localhost", "5432", "postgres", "root", "stockcontrol")
	fmt.Println(DataSourceName)
	db, err := sql.Open("postgres", DataSourceName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil

}

func TestProductTypeFindById(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created, _ := entity.NewProductType(6, "Special Material")

	//act
	dao := NewProductTypeDAO(db)
	_ = dao.Create(created)
	pt, err := dao.FindById(6)

	fmt.Println(pt)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.NotEqual(t, "", pt.Description)
	assert.NotEmpty(t, pt.CreatedAt)
	assert.NotEmpty(t, pt.UpdatedAt)
}

func TestProductTypeFindAll(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(6, "Special Material")

	//act
	dao := NewProductTypeDAO(db)
	_ = dao.Create(pt)

	ptList, err := dao.FindAll()

	for _, pt := range ptList {
		fmt.Println(pt)
	}

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(ptList), 1)
}

func TestProductTypeCreate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(6, "High Cost Material")

	//act
	dao := NewProductTypeDAO(db)
	_ = dao.Delete(pt)
	err = dao.Create(pt)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	ptFound, _ := dao.FindById(6)
	assert.Equal(t, ptFound.Description, pt.Description)

}

func TestProductTypeUpdate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(6, "Special Material")

	//act
	dao := NewProductTypeDAO(db)
	_ = dao.Create(pt)
	pt.Description = "Material Special"
	err = dao.Update(pt)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	ptFound, _ := dao.FindById(6)
	fmt.Println(ptFound)
	assert.Equal(t, "Material Special", ptFound.Description)
}

func TestProductTypeDelete(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(6, "Special Material")

	//act
	dao := NewProductTypeDAO(db)
	_ = dao.Create(pt)
	err = dao.Delete(pt)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	ptFound, _ := dao.FindById(6)
	assert.Equal(t, 0, ptFound.Id)
}
