package database

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestOperationFindById(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created, _ := entity.NewOperation(6, "Devolution")

	//act
	dao := NewOperationDAO(db)
	_ = dao.Create(created)
	op, err := dao.FindById(6)

	fmt.Println(op)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.NotEqual(t, "", op.Name)
	assert.NotEmpty(t, op.CreatedAt)
	assert.NotEmpty(t, op.UpdatedAt)
}

func TestOperationFindAll(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	op, _ := entity.NewOperation(6, "Devolution")

	//act
	dao := NewOperationDAO(db)
	_ = dao.Create(op)

	opList, err := dao.FindAll()

	for _, op := range opList {
		fmt.Println(op)
	}

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(opList), 1)
}

func TestOperationCreate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	op, _ := entity.NewOperation(2, "Withdrawal")

	//act
	dao := NewOperationDAO(db)
	_ = dao.Delete(op)
	err = dao.Create(op)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	opFound, _ := dao.FindById(2)
	assert.Equal(t, opFound.Name, op.Name)

}

func TestOperationUpdate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	op, _ := entity.NewOperation(6, "Devolution")

	//act
	dao := NewOperationDAO(db)
	_ = dao.Create(op)
	op.Name = "Supply"
	err = dao.Update(op)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	opFound, _ := dao.FindById(6)
	fmt.Println(opFound)
	assert.Equal(t, "Supply", opFound.Name)
}

func TestOperationDelete(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	op, _ := entity.NewOperation(6, "Devolution")

	//act
	dao := NewOperationDAO(db)
	_ = dao.Create(op)
	err = dao.Delete(op)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	opFound, _ := dao.FindById(6)
	assert.Equal(t, 0, opFound.Id)
}
