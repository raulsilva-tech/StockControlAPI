package entity

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewLabel(t *testing.T) {

	//arrange
	p, _ := NewProduct(1, "Test", ProductType{1, "Type", time.Now(), time.Now()})
	validDate := time.Now().Add(time.Hour * 24)

	//act
	l, err := NewLabel(1, "123", validDate, *p)
	//assert
	assert.Nil(t, err)
	assert.NotNil(t, l)
	assert.Equal(t, 1, l.Id)
	assert.Equal(t, "123", l.Code)
	assert.Equal(t, validDate, l.ValidDate)
	assert.NotEmpty(t, l.CreatedAt)
	assert.NotEmpty(t, l.UpdatedAt)

	myJson, _ := json.Marshal(l)
	fmt.Println(string(myJson))
}

func TestWhenLabelIdIsRequired(t *testing.T) {
	//arrange
	p, _ := NewProduct(1, "Test", ProductType{1, "Type", time.Now(), time.Now()})
	validDate := time.Now().Add(time.Hour * 24)
	//act
	_, err := NewLabel(0, "123", validDate, *p)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrIdIsRequired)
}

func TestWhenLabelCodeIsRequired(t *testing.T) {
	//arrange
	p, _ := NewProduct(1, "Test", ProductType{1, "Type", time.Now(), time.Now()})
	validDate := time.Now().Add(time.Hour * 24)
	//act
	_, err := NewLabel(1, "", validDate, *p)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrCodeIsRequired)
}

func TestWhenProductIsRequired(t *testing.T) {
	//arrange
	p, _ := NewProduct(0, "Test", ProductType{1, "Type", time.Now(), time.Now()})
	validDate := time.Now().Add(time.Hour * 24)
	//act
	_, err := NewLabel(1, "123", validDate, *p)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrProductIsRequired)
}

func TestWhenValidDateIsRequired(t *testing.T) {
	//arrange
	p, _ := NewProduct(1, "Test", ProductType{1, "Type", time.Now(), time.Now()})
	var validDate time.Time
	//act
	_, err := NewLabel(1, "123", validDate, *p)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrValidDateIsRequired)
}
