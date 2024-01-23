package entity

import (
	"errors"
	"time"
)

/*

	id				BIGINT	NOT NULL,
  	user_id			BIGINT NOT NULL,
	operation_id 	BIGINT NOT NULL,
	performed_at	timestamp,
	stock_product_id	BIGINT,
	quantity		BIGINT,
	label_id		BIGINT,
*/

var ErrPerformedAtIsRequired = errors.New("time of transaction is required")

type Transaction struct {
	Id           int `json:"id"`
	User         `json:"user"`
	Operation    `json:"operation"`
	StockProduct `json:"stock_product"`
	Label        `json:"label"`
	PerformedAt  time.Time `json:"performed_at"`
	Quantity     int       `json:"quantity"`
}

func (t *Transaction) Validate() error {
	if t.Id == 0 {
		return ErrIdIsRequired
	}
	if t.User.Id == 0 {
		return ErrUserIsRequired
	}
	if t.PerformedAt.IsZero() {
		return ErrPerformedAtIsRequired
	}
	if t.Quantity < 0 {
		return ErrInvalidQuantity
	}
	return nil
}

func NewTransaction(id int, user User, operation Operation, stockProduct StockProduct, label Label, performedAt time.Time, quantity int) (*Transaction, error) {
	t := &Transaction{
		Id:           id,
		User:         user,
		Operation:    operation,
		StockProduct: stockProduct,
		Label:        label,
		PerformedAt:  performedAt,
		Quantity:     quantity,
	}

	return t, t.Validate()
}
