package entity

import (
	"errors"
	"time"
)

var (
	ErrIdIsRequired          = errors.New("id is required")
	ErrDescriptionIsRequired = errors.New("description is required")
	ErrProductTypeIsRequired = errors.New("product type is required")
)

type Product struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ProductType `json:"product_type"`
}

func (p *Product) Validate() error {
	if p.Id == 0 {
		return ErrIdIsRequired
	}
	if p.Description == "" {
		return ErrDescriptionIsRequired
	}
	if p.ProductType.Id == 0 {
		return ErrProductTypeIsRequired
	}
	return nil
}

func NewProduct(id int, description string, pt ProductType) (*Product, error) {
	p := &Product{
		Id:          id,
		Description: description,
		ProductType: pt,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := p.Validate()

	return p, err
}
