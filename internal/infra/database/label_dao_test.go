package database

import (
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestLabelFindById(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(1, "")
	p, _ := entity.NewProduct(1, "", *pt)
	created, _ := entity.NewLabel(3, "123", time.Now(), *p)

	//act
	dao := NewLabelDAO(db)
	_ = dao.Create(created)
	l, err := dao.FindById(3)

	fmt.Println(l)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.NotEqual(t, "", l.Code)
	assert.NotEqual(t, "", l.ValidDate)
	assert.NotEmpty(t, l.CreatedAt)
	assert.NotEmpty(t, l.UpdatedAt)
}

func TestLabelFindAll(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(1, "")
	p, _ := entity.NewProduct(1, "", *pt)
	l, _ := entity.NewLabel(3, "123", time.Now(), *p)

	//act
	dao := NewLabelDAO(db)
	_ = dao.Create(l)
	lList, err := dao.FindAll()

	for _, l := range lList {
		fmt.Println(l)
	}

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(lList), 1)
}

func TestLabelCreate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(1, "")
	p, _ := entity.NewProduct(1, "", *pt)
	l, _ := entity.NewLabel(3, "123", time.Now(), *p)

	//act
	dao := NewLabelDAO(db)
	_ = dao.Delete(l)
	err = dao.Create(l)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	labelFound, _ := dao.FindById(3)
	assert.Equal(t, labelFound.Code, l.Code)
	assert.Equal(t, labelFound.Product.Id, l.Product.Id)

}

func TestLabelUpdate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(1, "")
	p, _ := entity.NewProduct(1, "", *pt)
	l, _ := entity.NewLabel(3, "123", time.Now(), *p)

	//act
	dao := NewLabelDAO(db)
	_ = dao.Create(l)
	l.Code = "888"
	l.UpdatedAt = time.Now()
	err = dao.Update(l)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	labelFound, _ := dao.FindById(3)
	fmt.Println(labelFound)
	assert.Equal(t, "888", labelFound.Code)

}

func TestLabelDelete(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	pt, _ := entity.NewProductType(1, "")
	p, _ := entity.NewProduct(1, "", *pt)
	l, _ := entity.NewLabel(3, "123", time.Now(), *p)

	//act
	dao := NewLabelDAO(db)
	_ = dao.Create(l)
	err = dao.Delete(l)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	labelFound, _ := dao.FindById(3)
	assert.Equal(t, 0, labelFound.Id)
}
