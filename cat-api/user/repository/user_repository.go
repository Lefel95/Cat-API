package repository

import (
	"cat-api/models"
	"cat-api/user"
	"database/sql"
	//mysql is a package to use mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type repo struct {
	db *sql.DB
}

const (
	driver         = "mysql"
	dataSourceName = "root:root@tcp(172.28.1.1)/conversion?parseTime=true"
)

//NewRepository instatiates a new breed.Repository using a mysql connection
func NewRepository(db *sql.DB) (user.Repository, error) {
	var r repo
	r.db = db

	err := r.ping()

	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (r *repo) ping() error {
	return r.db.Ping()
}

func (r *repo) FindUserByCredentials(login models.UserLogin) (bool, error) {
	err := r.ping()

	if err != nil {
		return false, err
	}

	stmt, err := r.db.Prepare("SELECT COUNT(*) FROM cats.users WHERE username = ? AND pass = SHA1(?)")

	if err != nil {
		return false, err
	}

	row := stmt.QueryRow(login.UserName, login.Password)

	var count int

	err = row.Scan(&count)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return count > 0, nil
}
