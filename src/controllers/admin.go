package controllers

import (
	"net/http"
	"fmt"
	"html/template"
	"db"
	"strconv"
	"time"
	"os"
	"io"
	"path/filepath"
	"encoding/json"
)

const VERI = "qcms"

type AdminTemp struct{
	Admin db.Admin
}

func (p *Entry)Admin_index(w http.ResponseWriter, r *http.Request) {
	temp := tempInterface{}
	adminId,err := r.Cookie("adminId")
	adminName,_ := r.Cookie("adminName")
	if err != nil{
		fmt.Fprintf(w, Exec_script("alert('登录失败');history.back();"))
		return
	}
	temp["adminId"], _ = strconv.Atoi(adminId.Value)
	temp["adminName"] = adminName.Value
	t, _ := template.ParseFiles("templates/admin/index.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
	fmt.Println(temp)
	t.Execute(w, temp)
}

func (p *Entry)Admin_login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.Method == "POST"{
		if VERI != r.FormValue("veri"){
			fmt.Fprintf(w, Exec_script("alert('验证码错误');history.back();"))
			return
		}
		userInfo := ConnDb.GetUser("username='"+r.FormValue("username")+"' && password='"+r.FormValue("password")+"'", "*", "limit 0,1")
		if len(userInfo) >= 1 {
			timeStr := time.Now().AddDate(1,0,0)
			cookie := http.Cookie{Name: "adminId", Value: strconv.Itoa(userInfo[0].Id), Path: "/", Expires:timeStr}
			http.SetCookie(w, &cookie)
			cookie = http.Cookie{Name: "adminName", Value: userInfo[0].Username, Path: "/", Expires:timeStr}
    		http.SetCookie(w, &cookie)
			fmt.Fprintf(w, Exec_script("alert('登录成功');window.location.href='/admin'"))

		}else{
			fmt.Fprintf(w, Exec_script("alert('登录失败');history.back();"))
			return
		}
		return
	}	
	t, _ := template.ParseFiles("templates/admin/login.html")
	t.Execute(w, nil)
}

func (p *Entry)Admin_logout(w http.ResponseWriter, r *http.Request){
	timeStr := time.Now().AddDate(-1,0,0)
	cookie := http.Cookie{Name: "adminId", Value: "", Path: "/", Expires:timeStr}
	http.SetCookie(w, &cookie)
	cookie = http.Cookie{Name: "adminName", Value: "", Path: "/", Expires:timeStr}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, Exec_script("alert('安全退出');window.location.href='/admin/login'"))
}

func (p *Entry)Admin_upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	type temp map[string]string
	jSon := temp{}
	jSon["msg"] = ""
	jSon["err"] = "上传失败"
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("filedata")
        if err != nil {
        	jSon["msg"] = ""
        	jSon["err"] = "上传失败"
        	b, _ := json.Marshal(jSon)
        	fmt.Fprintf(w, string(b))
            return
        }
        defer file.Close()
        timeInt := strconv.Itoa(int(time.Now().Unix()))
        ext := filepath.Ext(handler.Filename)
        fileName := "/static/upload/"+timeInt+ext
        f, err := os.OpenFile("."+fileName, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
        	jSon["msg"] = ""
        	jSon["err"] = "上传失败"
            b, _ := json.Marshal(jSon)
            fmt.Fprintf(w, string(b))
            return
        }
        defer f.Close()
        io.Copy(f, file)
        jSon["msg"] = fileName
        jSon["err"] = ""
        b, _ := json.Marshal(jSon)
        fmt.Fprintf(w, string(b))
        return
	}else{
		b, _ := json.Marshal(jSon)
        fmt.Fprintf(w, string(b))
	}
}

