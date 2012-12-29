package controllers

import (
	"net/http"
	"text/template"
	"fmt"
	"os"
	"io"
	"path/filepath"
	"time"
	"strconv"
	"db"
)

const SITE_ROOT = "http://127.0.0.1:12345/"

func (p *Entry)Ad_index(w http.ResponseWriter, r *http.Request) {
	temp := tempInterface{}
	temp["ad"] = ConnDb.GetAd("", "*", "")
	t, _ := template.ParseFiles("templates/ad/index.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
	//-- 验证登录 --
	adminId,err := r.Cookie("adminId")
	adminName,_ := r.Cookie("adminName")
	if err != nil{
		fmt.Fprintf(w, Exec_script("alert('登录失败');history.back();"))
		return
	}
	adminRs := db.Admin{}
	adminRs.Id, _ = strconv.Atoi(adminId.Value)
	adminRs.Username = adminName.Value
	temp["adminId"], _ = strconv.Atoi(adminId.Value)
	temp["adminName"] = adminName.Value
	temp["adminRs"] = adminRs
	fmt.Println(temp["adminRs"])
	//a := t.Templates()
	//fmt.Println(a)
	t.Execute(w, temp)
	//t.ExecuteTemplate(w, "html", temp)
}

func (p *Entry)Ad_add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("img")
        if err != nil {
        	fmt.Println(err)
            fmt.Fprintf(w, Exec_script("alert('上传失败');history.back();"))
            return
        }
        defer file.Close()
        timeInt := strconv.Itoa(int(time.Now().Unix()))
        ext := filepath.Ext(handler.Filename)
        fileName := "/static/upload/"+timeInt+ext
        f, err := os.OpenFile("."+fileName, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Fprintf(w, Exec_script("alert('上传失败');history.back();"))
            return
        }
        defer f.Close()
        io.Copy(f, file)
        stmt, _ := ConnDb.Conn.Prepare("INSERT ad_list SET name=?, img=?, url=?, width=?, height=?, addtime=?")
	    res, _ := stmt.Exec(r.FormValue("name"), fileName, r.FormValue("url"), r.FormValue("width"), r.FormValue("height"),timeInt)
	    res.LastInsertId()
	    fmt.Fprintf(w, Exec_script("alert('添加成功');window.location.href='/ad'"))
        return
	}
	temp := tempInterface{}
	//-- 验证登录 --
	adminId,err := r.Cookie("adminId")
	if err != nil{
		fmt.Fprintf(w, Exec_script("alert('登录失败');history.back();"))
		return
	}
	adminName,_ := r.Cookie("adminName")	
	temp["adminId"], _ = strconv.Atoi(adminId.Value)
	temp["adminName"] = adminName.Value
	t, _ := template.ParseFiles("templates/ad/add.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
	t.Execute(w, temp)
}

func (p *Entry)Ad_code(w http.ResponseWriter, r *http.Request) {
	id := p.Params[0]
	temp := tempInterface{}
	adRs := ConnDb.GetAd(" WHERE id = "+id+"", "*", "")
	temp["AdOne"] = adRs[0]
	//-- 验证登录 --
	adminId,err := r.Cookie("adminId")
	if err != nil{
		fmt.Fprintf(w, Exec_script("alert('登录失败');history.back();"))
		return
	}
	adminName,_ := r.Cookie("adminName")
	temp["adminId"], _ = strconv.Atoi(adminId.Value)
	temp["adminName"] = adminName.Value
	t, _ := template.ParseFiles("templates/ad/code.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
	t.Execute(w, temp)
}