package database

import "github.com/raulsilva-tech/StockControlAPI/internal/entity"

type ProductTypeInterface interface {
	Create(pt *entity.ProductType) error
	Update(pt *entity.ProductType) error
	Delete(pt *entity.ProductType) error
	FindById(id int) (*entity.ProductType, error)
	FindAll() ([]*entity.ProductType, error)
}
