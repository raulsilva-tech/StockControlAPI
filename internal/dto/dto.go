package dto

import "time"

type CreateProductTypeInput struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}
type CreateLabelInput struct {
	Id        int       `json:"id"`
	Code      string    `json:"code"`
	ValidDate time.Time `json:"valid_date"`
	ProductId int       `json:"product_id"`
}
type CreateOperationInput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type CreateProductInput struct {
	Id            int       `json:"id"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ProductTypeId int       `json:"product_type_id"`
}
type CreateStockProductInput struct {
	Id        int `json:"id"`
	Factor    int `json:"factor"`
	Quantity  int `json:"quantity"`
	StockId   int `json:"stock_id"`
	ProductId int `json:"product_id"`
}
type CreateStockInput struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}
type CreateUserInput struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type CreateTransactionInput struct {
	Id             int       `json:"id"`
	UserId         int       `json:"user_idd"`
	OperationId    int       `json:"operation_id"`
	StockProductId int       `json:"stock_product_id"`
	LabelId        int       `json:"label_id"`
	PerformedAt    time.Time `json:"performed_at"`
	Quantity       int       `json:"quantity"`
}
type CreateUserOperationInput struct {
	Id          int `json:"id"`
	UserId      int `json:"user_id"`
	OperationId int `json:"operation_id"`
}
type CreateUserSessionInput struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
}
