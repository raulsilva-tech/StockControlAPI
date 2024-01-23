package entity

import "time"

type Stock struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

func (s *Stock) Validate() error {
	if s.Id == 0 {
		return ErrIdIsRequired
	}
	if s.Description == "" {
		return ErrDescriptionIsRequired
	}
	return nil
}

func NewStock(id int, description string) (*Stock, error) {
	s := &Stock{
		Id:          id,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s, s.Validate()
}
