package main

import (
	_"code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {
	_, err := sql.Open("mysql", "ss:ss@127.0.0.1/aa?charset=utf8")
	if err != nil {
        fmt.Println("Failed to connect to the database !")
    }
}