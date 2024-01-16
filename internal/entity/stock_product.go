package entity

import (
	"errors"
	"time"
)

/*
  	id				BIGINT	NOT NULL,
	stock_id		BIGINT	NOT NULL,
 	product_id		BIGINT	NOT NULL,
	quantity		BIGINT,
	factor			INTEGER,
	created_at		timestamp,
	updated_at		timestamp,
	PRIMARY KEY ( id ));
*/

var (
	ErrStockIsRequired = errors.New("stock is required")
	ErrInvalidQuantity = errors.New("invalid quantity")
	ErrInvalidFactor   = errors.New("invalid factor")
)

type StockProduct struct {
	Id        int       `json:"id"`
	Factor    int       `json:"factor"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Stock     `json:"stock"`
	Product   `json:"product"`
}

func (sp *StockProduct) Validate() error {
	if sp.Id == 0 {
		return ErrIdIsRequired
	}
	if sp.Product.Id == 0 {
		return ErrProductIsRequired
	}
	if sp.Stock.Id == 0 {
		return ErrStockIsRequired
	}
	if sp.Factor < 0 {
		return ErrInvalidFactor
	}
	if sp.Factor == 0 {
		sp.Factor = 1
	}
	if sp.Quantity < 0 {
		return ErrInvalidQuantity
	}
	return nil
}

func NewStockProduct(id int, stock Stock, product Product, factor, quantity int) (*StockProduct, error) {

	sp := &StockProduct{
		Id:        id,
		Stock:     stock,
		Product:   product,
		Factor:    factor,
		Quantity:  quantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := sp.Validate()

	return sp, err
}
