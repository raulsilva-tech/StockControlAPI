package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUserSession(t *testing.T) {

	//arrange
	u, _ := NewUser(1, "Peter", "peter@mail.com", "123")
	timeNow := time.Now()
	//act
	us, err := NewUserSession(1, *u, timeNow, timeNow.Add(time.Minute*10))
	//assert
	assert.NotNil(t, us)
	assert.Nil(t, err)
	assert.Equal(t, 1, us.Id)
	assert.Equal(t, u.Id, us.User.Id)
	assert.Equal(t, timeNow, us.StartedAt)
	assert.Equal(t, timeNow.Add(time.Minute*10), us.FinishedAt)

}

func TestUserSessionWhenIdIsRequired(t *testing.T) {
	//arrange act
	u, _ := NewUser(1, "Peter", "peter@mail.com", "123")
	timeNow := time.Now()
	//act
	_, err := NewUserSession(0, *u, timeNow, timeNow.Add(time.Minute*10))
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrIdIsRequired)
}

func TestUserSessionWhenUserIsRequired(t *testing.T) {
	//arrange act
	u, _ := NewUser(0, "Peter", "peter@mail.com", "123")
	timeNow := time.Now()
	//act
	_, err := NewUserSession(1, *u, timeNow, timeNow.Add(time.Minute*10))
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrUserIsRequired)
}

func TestUserSessionWhenFinishedTimeIsInvalid(t *testing.T) {
	//arrange act
	u, _ := NewUser(1, "Peter", "peter@mail.com", "123")
	timeNow := time.Now()
	//act
	_, err := NewUserSession(1, *u, timeNow, timeNow.Add(time.Minute*10*-1))
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrInvalidFinishedTime)
}

func TestUserSessionWhenStartSessionIsRequired(t *testing.T) {
	//arrange act
	u, _ := NewUser(1, "Peter", "peter@mail.com", "123")
	var timeNow time.Time
	//act
	_, err := NewUserSession(1, *u, timeNow, timeNow.Add(time.Minute*10*-1))
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrStartSessionIsRequired)
}
