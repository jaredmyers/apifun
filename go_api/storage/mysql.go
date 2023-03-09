package storage

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jaredmyers/apifun/go_api/models"
)

type MySqlStore struct {
	db *sql.DB
}

func NewMySqlStore() (*MySqlStore, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MySqlStore{
		db: db,
	}, nil
}

func (db *MySqlStore) CreateUser(*models.User) error {
	return nil
}
func (db *MySqlStore) GetUser(*string) (*models.User, error) {
	return nil, nil
}
func (db *MySqlStore) UpdateUser(*models.User) error {
	return nil
}
func (db *MySqlStore) DeleteUser(*string) error {
	return nil
}
func (db *MySqlStore) GetUsers() error {
	log.Println("running GetUsers from MySqlStore")
	return nil
}
