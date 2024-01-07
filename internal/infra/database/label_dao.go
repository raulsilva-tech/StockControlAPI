package database

import (
	"database/sql"

	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
)

type LabelDAO struct {
	Db *sql.DB
}

func NewLabelDAO(db *sql.DB) *LabelDAO {
	return &LabelDAO{Db: db}
}

func (dao LabelDAO) Create(l *entity.Label) error {

	stmt, err := dao.Db.Prepare("insert into labels(id,code,product_id,valid_date,updated_at,created_at) values($1,$2,$3,$4,$5,$6)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(l.Id, l.Code, l.Product.Id, l.ValidDate, l.CreatedAt, l.UpdatedAt)
	return err
}

func (dao LabelDAO) Update(l *entity.Label) error {

	stmt, err := dao.Db.Prepare("update labels set code=$1,product_id=$2,valid_date=$3,updated_at=$4,created_at=$5 where id=$6")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(l.Code, l.Product.Id, l.ValidDate, l.CreatedAt, l.UpdatedAt, l.Id)
	return err
}

func (dao LabelDAO) Delete(l *entity.Label) error {

	stmt, err := dao.Db.Prepare("delete from labels where id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(l.Id)
	return err
}

func (dao LabelDAO) FindById(id int) (*entity.Label, error) {

	stmt, err := dao.Db.Prepare("select id,code,product_id,valid_date,created_at,updated_at from labels where id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var l entity.Label

	err = stmt.QueryRow(id).Scan(&l.Id, &l.Code, &l.Product.Id, &l.ValidDate, &l.CreatedAt, &l.UpdatedAt)

	return &l, err
}

func (dao LabelDAO) FindAll() ([]*entity.Label, error) {

	rows, err := dao.Db.Query("select id,code,product_id,valid_date,created_at,updated_at from labels")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lList []*entity.Label

	for rows.Next() {
		var l entity.Label
		err = rows.Scan(&l.Id, &l.Code, &l.Product.Id, &l.ValidDate, &l.CreatedAt, &l.UpdatedAt)
		if err == nil {
			lList = append(lList, &l)
		}
	}

	return lList, err
}
