package db

import (
	_"code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
)

type SqlType struct{
	Conn *sql.DB
	Error error
}

func Connect(host string, name string, user string, password string) (*sql.DB, error){
	return sql.Open("mysql", user+":"+password+"@tcp("+host+":3306)/"+name+"?charset=utf8")
}

func (sqlClass SqlType)Close() {
	sqlClass.Conn.Close()
}