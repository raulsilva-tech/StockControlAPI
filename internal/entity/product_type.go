package entity

import (
	"time"
)

type ProductType struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (pt *ProductType) Validate() error {
	if pt.Id == 0 {
		return ErrIdIsRequired
	}
	if pt.Description == "" {
		return ErrDescriptionIsRequired
	}
	return nil
}

func NewProductType(id int, description string) (*ProductType, error) {
	pt := &ProductType{
		Id:          id,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := pt.Validate()

	return pt, err
}
