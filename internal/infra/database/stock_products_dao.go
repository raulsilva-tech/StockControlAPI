package database

import (
	"database/sql"

	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
)

type StockProductDAO struct {
	Db *sql.DB
}

func NewStockProductDAO(db *sql.DB) *StockProductDAO {
	return &StockProductDAO{Db: db}
}

func (dao *StockProductDAO) Create(p *entity.StockProduct) error {

	stmt, err := dao.Db.Prepare("insert into stock_products(id,stock_id,product_id,quantity,factor,updated_at,created_at) values($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Id, p.Stock.Id, p.Product.Id, p.Quantity, p.Factor, p.UpdatedAt, p.CreatedAt)

	return err
}

func (dao *StockProductDAO) Update(p *entity.StockProduct) error {

	stmt, err := dao.Db.Prepare("update stock_products set stock_id=$1, product_id=$2,quantity=$3,factor=$4 ,updated_at=$5 where id=$6")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Stock.Id, p.Product.Id, p.Quantity, p.Factor, p.UpdatedAt, p.Id)

	return err
}

func (dao *StockProductDAO) Delete(p *entity.StockProduct) error {

	stmt, err := dao.Db.Prepare("delete from stock_products where id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Id)

	return err
}

func (dao *StockProductDAO) FindById(id int) (*entity.StockProduct, error) {

	stmt, err := dao.Db.Prepare("select id,stock_id,product_id,quantity,factor,created_at,updated_at from stock_products where id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var pFound entity.StockProduct

	err = stmt.QueryRow(id).Scan(&pFound.Id, &pFound.Stock.Id, &pFound.Product.Id, &pFound.Quantity, &pFound.Factor, &pFound.CreatedAt, &pFound.UpdatedAt)

	return &pFound, err
}

func (dao *StockProductDAO) FindAll() ([]*entity.StockProduct, error) {

	rows, err := dao.Db.Query("select id,stock_id,product_id,quantity,factor,created_at,updated_at from stock_products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pList []*entity.StockProduct

	for rows.Next() {
		var pFound entity.StockProduct
		err = rows.Scan(&pFound.Id, &pFound.Stock.Id, &pFound.Product.Id, &pFound.Quantity, &pFound.Factor, &pFound.CreatedAt, &pFound.UpdatedAt)
		pList = append(pList, &pFound)
	}

	return pList, err
}
