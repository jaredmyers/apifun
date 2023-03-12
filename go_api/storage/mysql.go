package storage

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jaredmyers/apifun/go_api/models"
)

type SqlStore struct {
	db *sql.DB
}

func NewMySqlStore() (UserServiceStorer, error) {

	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "testdb",
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

func (sql *SqlStore) CreateUser(*models.User) error {
	return nil
}
func (sql *SqlStore) GetUser(*string) (*models.User, error) {
	return nil, nil
}
func (sql *SqlStore) UpdateUser(*models.User) error {
	return nil
}
func (sql *SqlStore) DeleteUser(*string) error {
	return nil
}
func (sql *SqlStore) GetUsers() ([]*models.User, error) {
	log.Println("running GetUsers from MySqlStore through UserServiceStorer")

	var users []*models.User

	rows, err := sql.db.Query("select * from users")
	log.Println(rows)
	if err != nil {
		log.Println("returning error")
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		log.Println("inside the loop")
		err := rows.Scan(&user.Id, &user.Username, &user.Pw, &user.CreatedOn, &user.DeletedOn)
		if err != nil {
			log.Println(err)
			log.Println("returning error")
			return nil, err
		}
		users = append(users, &user)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		log.Println("returning error")
		return nil, err
	}

	return users, nil
}
