package database

import (
	"database/sql"

	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
)

type ProductDAO struct {
	Db *sql.DB
}

func NewProductDAO(db *sql.DB) *ProductDAO {
	return &ProductDAO{Db: db}
}

func (dao *ProductDAO) Create(p *entity.Product) error {

	stmt, err := dao.Db.Prepare("insert into products(id,description,updated_at,created_at,type_id) values($1,$2,$3,$4,$5)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Id, p.Description, p.UpdatedAt, p.CreatedAt, p.ProductType.Id)

	return err
}

func (dao *ProductDAO) Update(p *entity.Product) error {

	stmt, err := dao.Db.Prepare("update products set description=$1, updated_at=$2, type_id=$3 where id=$4")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Description, p.UpdatedAt, p.ProductType.Id, p.Id)

	return err
}

func (dao *ProductDAO) Delete(p *entity.Product) error {

	stmt, err := dao.Db.Prepare("delete from products where id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Id)

	return err
}

func (dao *ProductDAO) FindById(id int) (*entity.Product, error) {

	stmt, err := dao.Db.Prepare("select id,description,created_at,updated_at,type_id from products where id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var pFound entity.Product

	err = stmt.QueryRow(id).Scan(&pFound.Id, &pFound.Description, &pFound.CreatedAt, &pFound.UpdatedAt, &pFound.ProductType.Id)

	return &pFound, err
}

func (dao *ProductDAO) FindAll() ([]*entity.Product, error) {

	rows, err := dao.Db.Query("select id,description,created_at,updated_at,type_id from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pList []*entity.Product

	for rows.Next() {
		var p entity.Product
		err = rows.Scan(&p.Id, &p.Description, &p.CreatedAt, &p.UpdatedAt, &p.ProductType.Id)
		pList = append(pList, &p)
	}

	return pList, err
}
