package entity

import "time"

/*
id				BIGINT	NOT NULL,
  	name			VARCHAR(200)	NOT NULL,
	created_at		timestamp,
	updated_at		timestamp,
*/

type Operation struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *Operation) Validate() error {
	if s.Id == 0 {
		return ErrIdIsRequired
	}
	if s.Name == "" {
		return ErrNameIsRequired
	}
	return nil
}

func NewOperation(id int, name string) (*Operation, error) {
	s := &Operation{
		Id:        id,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s, s.Validate()
}
