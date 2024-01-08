package database

import (
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserSessionFindById(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created, _ := entity.NewUserSession(1, entity.User{Id: 2, Name: "User Name", Email: "mail@mail.com", Password: "1234"}, time.Now(), time.Now())

	//act
	dao := NewUserSessionDAO(db)
	_ = dao.Create(created)
	us, err := dao.FindById(1)

	fmt.Println(us)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	assert.Equal(t, created.Id, us.Id)
	assert.Equal(t, created.User.Id, us.User.Id)
	assert.NotEmpty(t, us.StartedAt)

}

func TestUserSessionFindAll(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created, _ := entity.NewUserSession(1, entity.User{Id: 1, Name: "User Name", Email: "mail@mail.com", Password: "1234"}, time.Now(), time.Now())

	//act
	dao := NewUserSessionDAO(db)
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

func TestUserSessionCreate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created, _ := entity.NewUserSession(1, entity.User{Id: 1, Name: "User Name", Email: "mail@mail.com", Password: "1234"}, time.Now(), time.Now())

	//act
	dao := NewUserSessionDAO(db)
	_ = dao.Delete(created)
	err = dao.Create(created)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(1)
	assert.Equal(t, found.Id, created.Id)

}

func TestUserSessionUpdate(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created, _ := entity.NewUserSession(1, entity.User{Id: 1, Name: "User Name", Email: "mail@mail.com", Password: "1234"}, time.Now(), time.Now())

	//act
	dao := NewUserSessionDAO(db)
	_ = dao.Create(created)
	created.User.Id = 2
	err = dao.Update(created)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(1)
	fmt.Println(found)
	assert.Equal(t, 2, found.User.Id)
}

func TestUserSessionDelete(t *testing.T) {

	//arrange
	db, err := arrangeDBConnection()
	assert.Nil(t, err)

	created, _ := entity.NewUserSession(1, entity.User{Id: 1, Name: "User Name", Email: "mail@mail.com", Password: "1234"}, time.Now(), time.Now())

	//act
	dao := NewUserSessionDAO(db)
	_ = dao.Create(created)
	err = dao.Delete(created)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, _ := dao.FindById(1)
	assert.Equal(t, 0, found.Id)
}
