package storage

import (
	"database/sql"
	"fmt"
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

func (s *SqlStore) RegisterUser(req *m.RegisterUserRequest) error {
	log.Println("registuser from storage running...")

	exists, err := s.checkUserName(req)
	if err != nil {
		return m.InternalErrResp{Orig: err, InternalCode: m.CodeInternalError}
	}

	if exists {
		log.Println("username exists, returning error")
		return m.InternalErrResp{Orig: fmt.Errorf("username already exists"), InternalCode: m.CodeInvalidArgument}
	}

	log.Println("username does not exist, creating user")

	query := "insert into users (user_name, pw) values (?, ?)"

	res, err := s.db.Exec(query, req.Username, req.Pw)
	if err != nil {
		return m.InternalErrResp{Orig: err, InternalCode: m.CodeInternalError}
	}
	log.Println(res)

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

func (s *SqlStore) checkUserName(req *m.RegisterUserRequest) (bool, error) {

	query := "select user_name from users where user_name=? and deleted_on is null"
	var username string

	err := s.db.QueryRow(query, req.Username).Scan(&username)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, m.InternalErrResp{Orig: err, InternalCode: m.CodeInternalError}
		}
		return false, nil
	}

	return true, nil
}
