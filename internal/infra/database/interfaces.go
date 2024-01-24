package database

import "github.com/raulsilva-tech/StockControlAPI/internal/entity"

type ProductTypeInterface interface {
	Create(pt *entity.ProductType) error
	Update(pt *entity.ProductType) error
	Delete(pt *entity.ProductType) error
	FindById(id int) (*entity.ProductType, error)
	FindAll() ([]*entity.ProductType, error)
}

type ProductInterface interface {
	Create(p *entity.Product) error
	Update(p *entity.Product) error
	Delete(p *entity.Product) error
	FindById(id int) (*entity.Product, error)
	FindAll() ([]*entity.Product, error)
}

type LabelInterface interface {
	Create(l *entity.Label) error
	Update(l *entity.Label) error
	Delete(l *entity.Label) error
	FindById(id int) (*entity.Label, error)
	FindAll() ([]*entity.Label, error)
}

type StockProductInterface interface {
	Create(sp *entity.StockProduct) error
	Update(sp *entity.StockProduct) error
	Delete(sp *entity.StockProduct) error
	FindById(id int) (*entity.StockProduct, error)
	FindAll() ([]*entity.StockProduct, error)
}

type UserInterface interface {
	Create(u *entity.User) error
	Update(u *entity.User) error
	Delete(u *entity.User) error
	FindById(id int) (*entity.User, error)
	FindAll() ([]*entity.User, error)
}

type OperationInterface interface {
	Create(op *entity.Operation) error
	Update(op *entity.Operation) error
	Delete(op *entity.Operation) error
	FindById(id int) (*entity.Operation, error)
	FindAll() ([]*entity.Operation, error)
}

type UserOperationInterface interface {
	Create(uo *entity.UserOperation) error
	Update(uo *entity.UserOperation) error
	Delete(uo *entity.UserOperation) error
	FindById(id int) (*entity.UserOperation, error)
	FindAll() ([]*entity.UserOperation, error)
}

type TransactionInterface interface {
	Create(tr *entity.Transaction) error
	Update(tr *entity.Transaction) error
	Delete(tr *entity.Transaction) error
	FindById(id int) (*entity.Transaction, error)
	FindAll() ([]*entity.Transaction, error)
}

type UserSessionInterface interface {
	Create(us *entity.UserSession) error
	Update(us *entity.UserSession) error
	Delete(us *entity.UserSession) error
	FindById(id int) (*entity.UserSession, error)
	FindAll() ([]*entity.UserSession, error)
}

type StockInterface interface {
	Create(st *entity.Stock) error
	Update(st *entity.Stock) error
	Delete(st *entity.Stock) error
	FindById(id int) (*entity.Stock, error)
	FindAll() ([]*entity.Stock, error)
}
