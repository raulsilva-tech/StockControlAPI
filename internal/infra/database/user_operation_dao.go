package database

import (
	"database/sql"

	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
)

type UserOperationDAO struct {
	Db *sql.DB
}

func NewUserOperationDAO(db *sql.DB) *UserOperationDAO {
	return &UserOperationDAO{Db: db}
}

func (dao *UserOperationDAO) Create(uo *entity.UserOperation) error {

	stmt, err := dao.Db.Prepare("insert into user_operations(id,user_id,operation_id,updated_at,created_at) values($1,$2,$3,$4,$5)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(uo.Id, uo.User.Id, uo.Operation.Id, uo.UpdatedAt, uo.CreatedAt)

	return err
}

func (dao *UserOperationDAO) Update(uo *entity.UserOperation) error {

	stmt, err := dao.Db.Prepare("update user_operations set user_id=$1,operation_id=$2, updated_at=$3 where id=$4")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(uo.User.Id, uo.Operation.Id, uo.UpdatedAt, uo.Id)

	return err
}

func (dao *UserOperationDAO) Delete(uo *entity.UserOperation) error {

	stmt, err := dao.Db.Prepare("delete from user_operations where id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(uo.Id)

	return err
}

func (dao *UserOperationDAO) FindById(id int) (*entity.UserOperation, error) {

	stmt, err := dao.Db.Prepare("select id,user_id,operation_id,created_at,updated_at from user_operations where id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var found entity.UserOperation

	err = stmt.QueryRow(id).Scan(&found.Id, &found.User.Id, &found.Operation.Id, &found.CreatedAt, &found.UpdatedAt)

	return &found, err
}

func (dao *UserOperationDAO) FindAll() ([]*entity.UserOperation, error) {

	rows, err := dao.Db.Query("select id,user_id,operation_id,created_at,updated_at from user_operations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*entity.UserOperation

	for rows.Next() {
		var found entity.UserOperation
		err = rows.Scan(&found.Id, &found.User.Id, &found.Operation.Id, &found.CreatedAt, &found.UpdatedAt)
		list = append(list, &found)
	}

	return list, err
}

func (dao *UserOperationDAO) FindByUserAndOperation(userId, operationId int) (*entity.UserOperation, error) {

	stmt, err := dao.Db.Prepare("select id,user_id,operation_id,created_at,updated_at from user_operations where user_id=$1, operation_id=$2")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var found entity.UserOperation

	err = stmt.QueryRow(userId, operationId).Scan(&found.Id, &found.User.Id, &found.Operation.Id, &found.CreatedAt, &found.UpdatedAt)

	return &found, err
}
