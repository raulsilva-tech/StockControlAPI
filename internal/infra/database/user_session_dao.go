package database

import (
	"database/sql"
	"errors"
	"time"

	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
)

var ErrLastSessionAlreadyFinished = errors.New("last session from this user already finished")

type UserSessionDAO struct {
	Db *sql.DB
}

func NewUserSessionDAO(db *sql.DB) *UserSessionDAO {
	return &UserSessionDAO{Db: db}
}

func (dao *UserSessionDAO) Create(us *entity.UserSession) error {

	//automatically finishing last session from the same user
	dao.CheckAndLogoutLastUserSession(us.User.Id)

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

func (dao *UserSessionDAO) GetNextId() (int, error) {

	var nextId int

	stmt, err := dao.Db.Prepare("select id from user_sessions order by id desc limit 1")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRow().Scan(&nextId)
	if err != nil {
		if err == sql.ErrNoRows {
			nextId = 1
			return nextId, nil
		} else {
			return 0, err
		}
	}

	nextId += 1

	return nextId, nil

}
func (dao *UserSessionDAO) CheckAndLogoutLastUserSession(userId int) error {

	stmt, err := dao.Db.Prepare("select id,user_id,started_at,finished_at from user_sessions where user_id=$1 order by id desc limit 1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	var us entity.UserSession

	err = stmt.QueryRow(userId).Scan(&us.Id, &us.User.Id, &us.StartedAt, &us.FinishedAt)

	if err != nil {
		return err
	}

	if us.Id != 0 {
		if us.FinishedAt.IsZero() {
			us.FinishedAt = time.Now()
			dao.Update(&us)

		} else {
			return ErrLastSessionAlreadyFinished
		}
	}

	return nil
}
