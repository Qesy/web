package db

import (
	_"code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
)

type SqlType struct{
	Conn *sql.DB
	Error string
}

func Connect() (*sql.DB, error){
	return sql.Open("mysql", "ss:ss@tcp(localhost:3306)/ad?charset=utf8")
}

func (sqlClass SqlType)Close() {
	sqlClass.Conn.Close()
}