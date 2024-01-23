package entity

import (
	"errors"
	"time"
)

var (
	ErrCodeIsRequired      = errors.New("code is required")
	ErrValidDateIsRequired = errors.New("valid date is required")
	ErrProductIsRequired   = errors.New("product is required")
)

type Label struct {
	Id        int       `json:"id"`
	Code      string    `json:"code"`
	ValidDate time.Time `json:"valid_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Product   `json:"product"`
}

func (l *Label) Validate() error {
	if l.Id == 0 {
		return ErrIdIsRequired
	}
	if l.Code == "" {
		return ErrCodeIsRequired
	}
	if l.ValidDate.IsZero() {
		return ErrValidDateIsRequired
	}
	if l.Product.Id == 0 {
		return ErrProductIsRequired
	}
	return nil
}

func NewLabel(id int, code string, validDate time.Time, product Product) (*Label, error) {
	l := &Label{
		Id:        id,
		Code:      code,
		ValidDate: validDate,
		Product:   product,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := l.Validate()

	return l, err
}
