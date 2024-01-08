package database

import (
	"database/sql"

	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
)

/*
	id				BIGINT	NOT NULL,
  	user_id			BIGINT NOT NULL,
	operation_id 	BIGINT NOT NULL,
	performed_at	timestamp,
	stock_product_id	BIGINT,
	quantity		BIGINT,
	label_id		BIGINT,
*/

type TransactionDAO struct {
	Db *sql.DB
}

func NewTransactionDAO(db *sql.DB) *TransactionDAO {
	return &TransactionDAO{Db: db}
}

func (dao *TransactionDAO) Create(tr *entity.Transaction) error {

	stmt, err := dao.Db.Prepare("insert into transactions(id,user_id,operation_id, performed_at, stock_product_id, quantity, label_id) values($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tr.Id, tr.User.Id, tr.Operation.Id, tr.PerformedAt, tr.StockProduct.Id, tr.Quantity, tr.Label.Id)

	return err
}

func (dao *TransactionDAO) Update(tr *entity.Transaction) error {

	stmt, err := dao.Db.Prepare("update transactions set user_id=$1,operation_id=$2,performed_at=$3,stock_product_id=$4, quantity=$5,label_id=$6 where id=$7")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tr.User.Id, tr.Operation.Id, tr.PerformedAt, tr.StockProduct.Id, tr.Quantity, tr.Label.Id, tr.Id)

	return err
}

func (dao *TransactionDAO) Delete(tr *entity.Transaction) error {

	stmt, err := dao.Db.Prepare("delete from transactions where id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tr.Id)

	return err
}

func (dao *TransactionDAO) FindById(id int) (*entity.Transaction, error) {

	stmt, err := dao.Db.Prepare("select id,user_id,operation_id, performed_at, stock_product_id, quantity, label_id from transactions where id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var tr entity.Transaction

	err = stmt.QueryRow(id).Scan(&tr.Id, &tr.User.Id, &tr.Operation.Id, &tr.PerformedAt, &tr.StockProduct.Id, &tr.Quantity, &tr.Label.Id)

	return &tr, err
}

func (dao *TransactionDAO) FindAll() ([]*entity.Transaction, error) {

	rows, err := dao.Db.Query("select id,user_id,operation_id, performed_at, stock_product_id, quantity, label_id from transactions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*entity.Transaction

	for rows.Next() {
		var tr entity.Transaction
		err = rows.Scan(&tr.Id, &tr.User.Id, &tr.Operation.Id, &tr.PerformedAt, &tr.StockProduct.Id, &tr.Quantity, &tr.Label.Id)
		list = append(list, &tr)
	}

	return list, err
}
