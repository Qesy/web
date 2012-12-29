package main

import (
	"net/http"
	"fmt"
	"html/template"
	_"code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
	"encoding/json"
	"common"
	"router"
)

type MyMux struct {
}

type Temp struct{
	User []*User
	Zhu	[]*Zhu
}

type User struct {
    UserName string
    Age string
}

type Zhu struct{
	Weight string
	Gender string
}

func HelloServer(w http.ResponseWriter, req *http.Request) { //-- 模版 --
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("charset", "utf-8")
	temp := Temp{
		User: []*User{
		{UserName : "Qesy", Age : "31"}, 
		{UserName : "Cai", Age : "32"}, 
		{UserName : "Xiaoqian", Age : "6"}}}
	temp.Zhu = []*Zhu{{Weight: "150", Gender: "M"}}		
	t, _ := template.ParseFiles("templates/hello.html")
	t.Execute(w, temp)
}

func CookieServer(w http.ResponseWriter, req *http.Request) {
	//y, _, d := time.Now().Date()
	cookie := http.Cookie{Name: "username", Value: "Qesy", Path: "/", Domain: "127.0.0.1"}
	http.SetCookie(w, &cookie)
}

func MysqlServer(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", "test:test@/test?charset=utf8")
	if err != nil {
        fmt.Println("数据库连接失败 !")
    }
    stmt, _ := db.Prepare("INSERT userinfo SET username=?")
    res, _ := stmt.Exec("astaxie")
    id, err := res.LastInsertId()
    fmt.Println(id)
}

func JsonServer(w http.ResponseWriter, req *http.Request) {
	user := User{UserName:"Qesy", Age:"31"}
	b, _ := json.Marshal(user)
	w.Write([]byte(b))
	//fmt.Fprintf(w, byte[]b)
	//fmt.Println(b)
}

func FanseServer(w http.ResponseWriter, req *http.Request) {
	path := common.SubStr(req.URL.Path, 1, len(req.URL.Path))
	router.Fetch_url(path, w, req)
}

func main() {
	/*http.Handle("/styles/", http.FileServer(http.Dir("static")))
	http.Handle("/scripts/", http.FileServer(http.Dir("static")))
	http.Handle("/images/", http.FileServer(http.Dir("static")))*/
	/*http.HandleFunc("/hello", HelloServer)	
	http.HandleFunc("/cookie", CookieServer)	
	http.HandleFunc("/mysql", MysqlServer)	
	http.HandleFunc("/json", JsonServer)
	http.HandleFunc("/test", controllers.TestServer)*/	
	/*http.HandleFunc("/", FanseServer)	
	err := http.ListenAndServe(":12345", nil)
	fmt.Println("Server start ok !")
	if err != nil {
		fmt.Println("Server start err !")
	}*/
	mux := &MyMux{}
    http.ListenAndServe(":12345", mux)
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := common.SubStr(r.URL.Path, 1, len(r.URL.Path))
	router.Fetch_url(path, w, r)
    /*if r.URL.Path == "/" {
        //sayhelloName(w, r)
        fmt.Println("aa")
        return
    }
    http.NotFound(w, r)*/
    return
}
