package storage

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	m "github.com/jaredmyers/apifun/go_api/models"
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

func (s *SqlStore) CreateUser(*m.User) error {
	return nil
}
func (s *SqlStore) GetUser(id int) (*m.User, error) {

	var user m.User

	query := "select * from users where id=? and deleted_on is null"
	row := s.db.QueryRow(query, id)

	if err := row.Scan(&user.Id, &user.Username, &user.Pw, &user.CreatedOn, &user.DeletedOn); err != nil {
		if err == sql.ErrNoRows {
			return nil, m.InternalErrResp{Orig: err, InternalCode: m.CodeNotFound}
		}
		return nil, m.InternalErrResp{Orig: err, InternalCode: m.CodeInternalError}
	}

	return &user, nil
}
func (s *SqlStore) UpdateUser(*m.User) error {
	return nil
}
func (s *SqlStore) DeleteUser(id int) error {
	return nil
}
func (s *SqlStore) GetUsers() ([]*m.User, error) {
	log.Println("GetUsers from store...")

	var users []*m.User

	rows, err := s.db.Query("select * from users where deleted_on is null")
	if err != nil {
		log.Println("store error 1")
		return nil, m.InternalErrResp{Orig: err, InternalCode: m.CodeInternalError}
	}
	defer rows.Close()

	for rows.Next() {
		var user m.User
		err := rows.Scan(&user.Id, &user.Username, &user.Pw, &user.CreatedOn, &user.DeletedOn)
		if err != nil {
			log.Println("store error 2")
			return nil, m.InternalErrResp{Orig: err, InternalCode: m.CodeInternalError}
		}
		users = append(users, &user)
	}

	err = rows.Err()
	if err != nil {
		log.Println("store error 3")
		return nil, m.InternalErrResp{Orig: err, InternalCode: m.CodeInternalError}
	}

	return users, nil
}
