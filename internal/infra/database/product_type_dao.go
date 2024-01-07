package database

import (
	"database/sql"

	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
)

type ProductTypeDAO struct {
	Db *sql.DB
}

func NewProductTypeDAO(db *sql.DB) *ProductTypeDAO {
	return &ProductTypeDAO{Db: db}
}

func (dao *ProductTypeDAO) Create(pt *entity.ProductType) error {

	stmt, err := dao.Db.Prepare("insert into product_types(id,description,updated_at,created_at) values($1,$2,$3,$4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pt.Id, pt.Description, pt.UpdatedAt, pt.CreatedAt)

	return err
}

func (dao *ProductTypeDAO) Update(pt *entity.ProductType) error {

	stmt, err := dao.Db.Prepare("update product_types set description=$1, updated_at=$2 where id=$3")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pt.Description, pt.UpdatedAt, pt.Id)

	return err
}

func (dao *ProductTypeDAO) Delete(pt *entity.ProductType) error {

	stmt, err := dao.Db.Prepare("delete from product_types where id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pt.Id)

	return err
}

func (dao *ProductTypeDAO) FindById(id int) (*entity.ProductType, error) {

	stmt, err := dao.Db.Prepare("select id,description,created_at,updated_at from product_types where id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var ptFound entity.ProductType

	err = stmt.QueryRow(id).Scan(&ptFound.Id, &ptFound.Description, &ptFound.CreatedAt, &ptFound.UpdatedAt)

	return &ptFound, err
}

func (dao *ProductTypeDAO) FindAll() ([]*entity.ProductType, error) {

	rows, err := dao.Db.Query("select id,description,created_at,updated_at from product_types")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ptList []*entity.ProductType

	for rows.Next() {
		var pt entity.ProductType
		err = rows.Scan(&pt.Id, &pt.Description, &pt.CreatedAt, &pt.UpdatedAt)
		ptList = append(ptList, &pt)
	}

	return ptList, err
}
