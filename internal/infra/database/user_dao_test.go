package database

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserFindById(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created, _ := entity.NewUser(1, "User Name", "mail@mail.com", "1234")

	//act
	dao := NewUserDAO(db)
	_ = dao.Create(created)
	user, err := dao.FindById(1)

	fmt.Println(user)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.Equal(t, created.Id, user.Id)
	assert.Equal(t, created.Name, user.Name)
	assert.Equal(t, created.Email, user.Email)
	assert.Equal(t, created.Password, user.Password)
	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdatedAt)
}

func TestUserFindAll(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created, _ := entity.NewUser(1, "User Name", "mail@mail.com", "1234")

	//act
	dao := NewUserDAO(db)
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

func TestUserCreate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created, _ := entity.NewUser(2, "User Name", "mail@mail.com", "1234")

	//act
	dao := NewUserDAO(db)
	_ = dao.Delete(created)
	err = dao.Create(created)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(2)
	assert.Equal(t, found.Id, created.Id)

}

func TestUserUpdate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created, _ := entity.NewUser(1,"User Name","mail2@mail.com","1234")


	//act
	dao := NewUserDAO(db)
	_ = dao.Create(created)
	created.Name = "Mark"
	err = dao.Update(created)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(1)
	fmt.Println(found)
	assert.Equal(t, "Mark", found.Name)
}

func TestUserDelete(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created, _ := entity.NewUser(1,"User Name","mail@mail.com","1234")

	//act
	dao := NewUserDAO(db)
	_ = dao.Create(created)
	err = dao.Delete(created)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(1)
	assert.Equal(t, 0, found.Id)
}
