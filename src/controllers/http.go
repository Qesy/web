package controllers

import (
	"net/http"
	"fmt"
	_"code.google.com/p/go-mysql-driver/mysql"
	"db"
)

const token = "young123"
const secret = "8aa2a6ae0b2a0e6c9390b8e3f6626b47"
const SITE_ROOT = "http://127.0.0.1:12345/"

type Entry struct{
	Controller string
	Action string
	Params []string
}

type tempInterface map[string]interface{}

var ConnDb db.SqlType

func (p *Entry)ErrorStatus(w http.ResponseWriter, r *http.Request, code int) {	
	w.WriteHeader(code)
	fmt.Fprintf(w, "Http page err code is %d", code)
}

func Exec_script(str string) string{
	return "<html><body><script>"+str+"</script></body></html>"
}

func DbConn() {
	cTemp := Siteinfo()
	ConnDb.Conn, ConnDb.Error = db.Connect(cTemp.DbInfo.DbHost, cTemp.DbInfo.DbName, cTemp.DbInfo.DbUser, cTemp.DbInfo.DbPass)
}

func DbDisconn() {
	ConnDb.Conn.Close()
}

func SetCookie() {
	
}

func GetCookie() {
	
}