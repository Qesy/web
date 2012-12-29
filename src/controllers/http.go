package controllers

import (
	"net/http"
	"fmt"
	_"code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
	"db"
)

type Entry struct{
	Controller string
	Action string
	Params []string
}

type tempInterface map[string]interface{}

var ConnDb db.SqlType

func Connect() (*sql.DB, error){
	return sql.Open("mysql", "ss:ss@tcp(localhost:3306)/aaa?charset=utf8")
}

func Close(db *sql.DB) {
	db.Close()
}

func (p *Entry)ErrorStatus(w http.ResponseWriter, r *http.Request, code int) {	
	w.WriteHeader(code)
	fmt.Fprintf(w, "Http page err code is %d", code)
}

func Exec_script(str string) string{
	return "<html><body><script>"+str+"</script></body></html>"
}

func DbConn() {
	conn, _ := db.Connect()
	ConnDb.Conn = conn
}

func DbDisconn() {
	ConnDb.Close()
}

func SetCookie() {
	
}

func GetCookie() {
	
}