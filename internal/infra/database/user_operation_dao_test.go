package database

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserOperationFindById(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	uo, _ := entity.NewUserOperation(1, entity.User{Id: 1}, entity.Operation{Id: 6})

	//act
	dao := NewUserOperationDAO(db)
	_ = dao.Create(uo)
	found, err := dao.FindById(1)

	fmt.Println(found)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.Equal(t, uo.Id, found.Id)
	assert.Equal(t, uo.Operation.Id, found.Operation.Id)
	assert.Equal(t, uo.User.Id, found.User.Id)
	assert.NotEmpty(t, found.CreatedAt)
	assert.NotEmpty(t, found.UpdatedAt)
}

func TestUserOperationFindAll(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	uo, _ := entity.NewUserOperation(1, entity.User{Id: 1}, entity.Operation{Id: 6})

	//act
	dao := NewUserOperationDAO(db)
	_ = dao.Create(uo)

	list, err := dao.FindAll()

	for _, op := range list {
		fmt.Println(op)
	}

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(list), 1)
}

func TestUserOperationCreate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	uo, _ := entity.NewUserOperation(1, entity.User{Id: 1}, entity.Operation{Id: 6})

	//act
	dao := NewUserOperationDAO(db)
	_ = dao.Delete(uo)
	err = dao.Create(uo)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(1)
	assert.Equal(t, found.Id, uo.Id)

}

func TestUserOperationUpdate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	uo, _ := entity.NewUserOperation(1, entity.User{Id: 1}, entity.Operation{Id: 6})

	//act
	dao := NewUserOperationDAO(db)
	_ = dao.Create(uo)
	uo.Operation.Id = 2
	err = dao.Update(uo)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(1)
	fmt.Println(found)
	assert.Equal(t, 2, found.Operation.Id)
}

func TestUserOperationDelete(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	uo, _ := entity.NewUserOperation(1, entity.User{Id: 1}, entity.Operation{Id: 6})

	//act
	dao := NewUserOperationDAO(db)
	_ = dao.Create(uo)
	err = dao.Delete(uo)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(1)
	assert.Equal(t, 0, found.Id)
}
