package repository

import (
	"cat-api/breed"
	"cat-api/models"
	"database/sql"
	"encoding/json"
	"fmt"

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
func NewRepository(db *sql.DB) (breed.Repository, error) {
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

func (r *repo) GetBreedByName(breedName string) (*models.Breed, error) {
	err := r.ping()

	if err != nil {
		return nil, err
	}
	
	stmt, err := r.db.Prepare("SELECT attributes FROM cats.breeds WHERE id = ?")

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(breedName)

	var breed models.Breed
	var att []byte

	err = row.Scan(&att)

	if err != nil {
		if err == sql.ErrNoRows {
			return &models.Breed{}, nil
		}

		return nil, err
	}

	err = json.Unmarshal(att, &breed)

	if err != nil {
		return nil, err
	}

	return &breed, nil
}

func (r *repo) InsertBreed(breed *models.Breed) error {
	err := r.ping()

	if err != nil {
		return err
	}

	stmt, err := r.db.Prepare("INSERT INTO cats.breeds (id, attributes) VALUES(?, ?)")

	if err != nil {
		return err
	}

	breedJson, err := json.Marshal(&breed)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(breed.ID, breedJson)

	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()

	if rows <= 0 {
		return fmt.Errorf("failed to insert breed to database")
	}

	return nil
}
