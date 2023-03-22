package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jaredmyers/apifun/go_api/models"
)

type SqlStore struct {
	db *sql.DB
}

func NewMySqlStore(dbName string) (UserServiceStorer, error) {

	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: dbName,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	} else {
		log.Println("no ping error....")
	}

	return &SqlStore{
		db: db,
	}, nil
}

func NewPostgresStore() (UserServiceStorer, error) {
	return nil, nil
}

func (s *SqlStore) CreateUser(*models.User) error {
	return nil
}
func (s *SqlStore) GetUser(id int) (*models.User, error) {

	var user models.User

	query := "select * from users where id=? and deleted_on is null"
	row := s.db.QueryRow(query, id)

	if err := row.Scan(&user.Id, &user.Username, &user.Pw, &user.CreatedOn, &user.DeletedOn); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user id %d does not exist", id)
		}
		return nil, err
	}

	return &user, nil
}
func (s *SqlStore) UpdateUser(*models.User) error {
	return nil
}
func (s *SqlStore) DeleteUser(id int) error {
	return nil
}
func (s *SqlStore) GetUsers() ([]*models.User, error) {
	log.Println("GetUsers from store...")

	var users []*models.User

	rows, err := s.db.Query("select * from users where deleted_on is null")
	if err != nil {
		log.Println("store error 1")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Username, &user.Pw, &user.CreatedOn, &user.DeletedOn)
		if err != nil {
			log.Println("store error 2")
			return nil, err
		}
		users = append(users, &user)
	}

	err = rows.Err()
	if err != nil {
		log.Println("store error 3")
		return nil, err
	}

	return users, nil
}
