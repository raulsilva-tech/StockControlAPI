package database

import (
	"database/sql"

	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
)

type UserSessionDAO struct {
	Db *sql.DB
}

func NewUserSessionDAO(db *sql.DB) *UserSessionDAO {
	return &UserSessionDAO{Db: db}
}

func (dao *UserSessionDAO) Create(us *entity.UserSession) error {

	stmt, err := dao.Db.Prepare("insert into user_sessions(id,user_id,started_at,finished_at) values($1,$2,$3,$4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(us.Id, us.User.Id, us.StartedAt, us.FinishedAt)

	return err
}

func (dao *UserSessionDAO) Update(us *entity.UserSession) error {

	stmt, err := dao.Db.Prepare("update user_sessions set user_id=$1,started_at=$2,finished_at=$3 where id=$4")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(us.User.Id, us.StartedAt, us.FinishedAt, us.Id)

	return err
}

func (dao *UserSessionDAO) Delete(us *entity.UserSession) error {

	stmt, err := dao.Db.Prepare("delete from user_sessions where id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(us.Id)

	return err
}

func (dao *UserSessionDAO) FindById(id int) (*entity.UserSession, error) {

	stmt, err := dao.Db.Prepare("select id,user_id,started_at,finished_at from user_sessions where id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var us entity.UserSession

	err = stmt.QueryRow(id).Scan(&us.Id, &us.User.Id, &us.StartedAt, &us.FinishedAt)

	return &us, err
}

func (dao *UserSessionDAO) FindAll() ([]*entity.UserSession, error) {

	rows, err := dao.Db.Query("select id,user_id,started_at,finished_at from user_sessions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*entity.UserSession

	for rows.Next() {
		var us entity.UserSession
		err = rows.Scan(&us.Id, &us.User.Id, &us.StartedAt, &us.FinishedAt)
		list = append(list, &us)
	}

	return list, err
}
