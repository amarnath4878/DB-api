package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Address  string
	Username string
	Password string
	Port     string
	Database string
}

func New(address, username, pwd, port, database string) *Database {
	return &Database{
		Address:  address,
		Username: username,
		Password: pwd,
		Port:     port,
		Database: database,
	}
}
func (db *Database) Connect() *sql.DB {
	conn, err := sql.Open("mysql", "root:Amarnath99@@tcp(localhost:3306)/table")
	if err != nil {
		log.Println(err)
	}
	return conn
}
