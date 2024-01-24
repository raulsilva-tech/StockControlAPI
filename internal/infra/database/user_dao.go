<<<<<<< HEAD
package database

import (
	"database/sql"

	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
)

type UserDAO struct {
	Db *sql.DB
}

func NewUserDAO(db *sql.DB) *UserDAO {
	return &UserDAO{Db: db}
}

func (dao *UserDAO) Create(user *entity.User) error {

	stmt, err := dao.Db.Prepare("insert into users(id,name,email,password,updated_at,created_at) values($1,$2,$3,$4,$5,$6)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Name, user.Email, user.Password, user.UpdatedAt, user.CreatedAt)

	return err
}

func (dao *UserDAO) Update(user *entity.User) error {

	stmt, err := dao.Db.Prepare("update users set name=$1, email=$2,password=$3,updated_at=$4 where id=$5")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Password, user.UpdatedAt, user.Id)

	return err
}

func (dao *UserDAO) Delete(user *entity.User) error {

	stmt, err := dao.Db.Prepare("delete from users where id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)

	return err
}

func (dao *UserDAO) FindById(id int) (*entity.User, error) {

	stmt, err := dao.Db.Prepare("select id,name,email,password,created_at,updated_at from users where id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user entity.User

	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	return &user, err
}

func (dao *UserDAO) FindAll() ([]*entity.User, error) {

	rows, err := dao.Db.Query("select id,name,email,password,created_at,updated_at from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*entity.User

	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		list = append(list, &user)
	}

	return list, err
}

func (dao *UserDAO) FindByEmailAndPassword(email, password string) (*entity.User, error) {

	stmt, err := dao.Db.Prepare("select id, name, email, password, created_at, updated_at from users where email=$1 and password=$2")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user entity.User

	err = stmt.QueryRow(email, password).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	return &user, err

}
=======
package database

import (
	"database/sql"

	"github.com/raulsilva-tech/StockControlAPI/internal/entity"
)

type UserDAO struct {
	Db *sql.DB
}

func NewUserDAO(db *sql.DB) *UserDAO {
	return &UserDAO{Db: db}
}

func (dao *UserDAO) Create(user *entity.User) error {

	stmt, err := dao.Db.Prepare("insert into users(id,name,email,password,updated_at,created_at) values($1,$2,$3,$4,$5,$6)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Name, user.Email, user.Password, user.UpdatedAt, user.CreatedAt)

	return err
}

func (dao *UserDAO) Update(user *entity.User) error {

	stmt, err := dao.Db.Prepare("update users set name=$1, email=$2,password=$3,updated_at=$4 where id=$5")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Password, user.UpdatedAt, user.Id)

	return err
}

func (dao *UserDAO) Delete(user *entity.User) error {

	stmt, err := dao.Db.Prepare("delete from users where id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)

	return err
}

func (dao *UserDAO) FindById(id int) (*entity.User, error) {

	stmt, err := dao.Db.Prepare("select id,name,email,password,created_at,updated_at from users where id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user entity.User

	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	return &user, err
}

func (dao *UserDAO) FindAll() ([]*entity.User, error) {

	rows, err := dao.Db.Query("select id,name,email,password,created_at,updated_at from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*entity.User

	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		list = append(list, &user)
	}

	return list, err
}

func (dao *UserDAO) FindByEmailAndPassword(email, password string) (*entity.User, error) {

	stmt, err := dao.Db.Prepare("select id, name, email, password, created_at, updated_at from users where email=$1 and password=$2")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user entity.User

	err = stmt.QueryRow(email, password).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	return &user, err

}
>>>>>>> d4eba3be9444a00975090f26358cb6323f2e2548
