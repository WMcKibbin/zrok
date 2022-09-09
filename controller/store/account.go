package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Account struct {
	Model
	Email    string
	Password string
	Token    string
}

func (self *Store) CreateAccount(a *Account, tx *sqlx.Tx) (int, error) {
	stmt, err := tx.Prepare("insert into accounts (email, password, token) values (?, ?, ?)")
	if err != nil {
		return 0, errors.Wrap(err, "error preparing accounts insert statement")
	}
	res, err := stmt.Exec(a.Email, a.Password, a.Token)
	if err != nil {
		return 0, errors.Wrap(err, "error executing accounts insert statement")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "error retrieving last accounts insert id")
	}
	return int(id), nil
}

func (self *Store) GetAccount(id int, tx *sqlx.Tx) (*Account, error) {
	a := &Account{}
	if err := tx.QueryRowx("select * from accounts where id = ?", id).StructScan(a); err != nil {
		return nil, errors.Wrap(err, "error selecting account by id")
	}
	return a, nil
}

func (self *Store) FindAccountWithEmail(email string, tx *sqlx.Tx) (*Account, error) {
	a := &Account{}
	if err := tx.QueryRowx("select * from accounts where email = ?", email).StructScan(a); err != nil {
		return nil, errors.Wrap(err, "error selecting account by email")
	}
	return a, nil
}

func (self *Store) FindAccountWithToken(token string, tx *sqlx.Tx) (*Account, error) {
	a := &Account{}
	if err := tx.QueryRowx("select * from accounts where token = ?", token).StructScan(a); err != nil {
		return nil, errors.Wrap(err, "error selecting account by token")
	}
	return a, nil
}
