package database

import (
	"database/sql"

	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
)

type OperationDAO struct {
	Db *sql.DB
}

func NewOperationDAO(db *sql.DB) *OperationDAO {
	return &OperationDAO{Db: db}
}

func (dao *OperationDAO) Create(st *entity.Operation) error {

	stmt, err := dao.Db.Prepare("insert into operations(id,name,updated_at,created_at) values($1,$2,$3,$4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(st.Id, st.Name, st.UpdatedAt, st.CreatedAt)

	return err
}

func (dao *OperationDAO) Update(st *entity.Operation) error {

	stmt, err := dao.Db.Prepare("update operations set name=$1, updated_at=$2 where id=$3")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(st.Name, st.UpdatedAt, st.Id)

	return err
}

func (dao *OperationDAO) Delete(st *entity.Operation) error {

	stmt, err := dao.Db.Prepare("delete from operations where id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(st.Id)

	return err
}

func (dao *OperationDAO) FindById(id int) (*entity.Operation, error) {

	stmt, err := dao.Db.Prepare("select id,name,created_at,updated_at from operations where id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var stFound entity.Operation

	err = stmt.QueryRow(id).Scan(&stFound.Id, &stFound.Name, &stFound.CreatedAt, &stFound.UpdatedAt)

	return &stFound, err
}

func (dao *OperationDAO) FindAll() ([]*entity.Operation, error) {

	rows, err := dao.Db.Query("select id,name,created_at,updated_at from operations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stList []*entity.Operation

	for rows.Next() {
		var st entity.Operation
		err = rows.Scan(&st.Id, &st.Name, &st.CreatedAt, &st.UpdatedAt)
		stList = append(stList, &st)
	}

	return stList, err
}
