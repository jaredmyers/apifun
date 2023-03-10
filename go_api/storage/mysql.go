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

func NewMySqlStore() (UserServiceStorer, error) {

	/*
		cfg := mysql.Config{

		}
	*/

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	} else {
		log.Println("no ping error....")
	}

	return &MySqlStore{
		db: db,
	}, nil
}

func (ms *MySqlStore) CreateUser(*models.User) error {
	return nil
}
func (ms *MySqlStore) GetUser(*string) (*models.User, error) {
	return nil, nil
}
func (ms *MySqlStore) UpdateUser(*models.User) error {
	return nil
}
func (ms *MySqlStore) DeleteUser(*string) error {
	return nil
}
func (ms *MySqlStore) GetUsers() ([]*models.User, error) {
	log.Println("running GetUsers from MySqlStore through UserServiceStorer")

	var users []*models.User

	rows, err := ms.db.Query("select * from users")
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
