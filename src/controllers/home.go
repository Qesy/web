package controllers

import (
	"net/http"
	"fmt"
	"html/template"
	//_"code.google.com/p/go-mysql-driver/mysql"
	//"database/sql"
	//"db"
)

func (p *Entry)Home_index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	/*m := a{}
	m["user"] = []User{{Uid:1,UserName:"Young"},{Uid:2,UserName:"Cai"}}
	m["zhu"] = Zhu{"100g", "100cm"}
	fmt.Println(m)*/
	fmt.Fprintf(w, "<font color='#222222'><b>酒悦久广告系统! (Golang 版)</b></font>")
}

func (p *Entry)Home_ad(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.RemoteAddr)
	//fmt.Println(p.Params[0])
	w.Header().Set("Content-Type", "text/html")
	if(len(p.Params) == 0 || p.Params[0] == "" || p.Params[0] == "0"){
		fmt.Fprintf(w, "Get err !")
		return
	}	
	temp := tempInterface{}
	temp["adRs"] = ConnDb.GetAd(" WHERE id = "+p.Params[0]+"", "*", "")
	t, _ := template.ParseFiles("templates/home/ad.html")
	t.Execute(w, temp)
}

func (p *Entry)Aaa_bbb(w http.ResponseWriter, r *http.Request){
	/*w.Header().Set("Content-Type", "text/html")
	var users []User
	var user User
	temp := Temp{}
	//db, err := sql.Open("mysql", "ss:ss@tcp(localhost:3306)/aaa?charset=utf8")
	db, err := db.Connect()
	if err != nil{
		fmt.Fprintf(w, "%s", err)
		return
	}
	rows, err := db.Query("SELECT uid, username FROM qcms_user")
    var uid int;
    var username string;
    for rows.Next(){
    		rows.Scan(&uid,&username);
    		user.Uid = uid
			user.UserName = username
			users =append(users,user)
	}
	temp.User = users*/
    //-- 添加数据 --
    /*stmt, _ := db.Prepare("INSERT qcms_user SET username=?")
    res, _ := stmt.Exec("astaxie")
    id, _ := res.LastInsertId()
    fmt.Println(id)*/

    //-- 修改数据 --
    //stmt, err := db.Prepare("update qcms_user set username=? where uid=?")
    //stmt.Exec("astaxieupdate", 2)

	
	/*temp.User = []*User{
		{UserName : "Qesy", Uid : "31"}, 
		{UserName : "Cai", Uid : "32"}, 
		{UserName : "Xiaoqian", Uid : "6"}}

	temp.Zhu = []Zhu{{Weight: "150", Gender: "M"}}*/		
	//t, _ := template.ParseFiles("templates/hello.html")
	//t.Execute(w, temp)	
}
