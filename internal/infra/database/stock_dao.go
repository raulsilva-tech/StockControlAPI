package database

import (
	"database/sql"

	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
)

type StockDAO struct {
	Db *sql.DB
}

func NewStockDAO(db *sql.DB) *StockDAO {
	return &StockDAO{Db: db}
}

func (dao *StockDAO) Create(st *entity.Stock) error {

	stmt, err := dao.Db.Prepare("insert into stocks(id,description,updated_at,created_at) values($1,$2,$3,$4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(st.Id, st.Description, st.UpdatedAt, st.CreatedAt)

	return err
}

func (dao *StockDAO) Update(st *entity.Stock) error {

	stmt, err := dao.Db.Prepare("update stocks set description=$1, updated_at=$2 where id=$3")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(st.Description, st.UpdatedAt, st.Id)

	return err
}

func (dao *StockDAO) Delete(st *entity.Stock) error {

	stmt, err := dao.Db.Prepare("delete from stocks where id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(st.Id)

	return err
}

func (dao *StockDAO) FindById(id int) (*entity.Stock, error) {

	stmt, err := dao.Db.Prepare("select id,description,created_at,updated_at from stocks where id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var stFound entity.Stock

	err = stmt.QueryRow(id).Scan(&stFound.Id, &stFound.Description, &stFound.CreatedAt, &stFound.UpdatedAt)

	return &stFound, err
}

func (dao *StockDAO) FindAll() ([]*entity.Stock, error) {

	rows, err := dao.Db.Query("select id,description,created_at,updated_at from stocks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stList []*entity.Stock

	for rows.Next() {
		var st entity.Stock
		err = rows.Scan(&st.Id, &st.Description, &st.CreatedAt, &st.UpdatedAt)
		stList = append(stList, &st)
	}

	return stList, err
}
