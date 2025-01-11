package db

import (
	"database/sql"
	"fmt"
)

type ConfigDB struct {
	IP       string
	Port     string
	DbName   string
	UserName string
	Password string
}

func ConnectMySql(c ConfigDB) (*sql.DB, error) {
	connctionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", c.UserName, c.Password, c.IP, c.DbName)
	con, err := sql.Open("mysql", connctionString)

	if err != nil {
		return nil, err
	}

	err = con.Ping()
	if err != nil {
		return nil, err
	}

	return con, nil

}
