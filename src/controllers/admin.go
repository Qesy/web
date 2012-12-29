package controllers

import (
	"net/http"
	"fmt"
	"html/template"
	"db"
	"strconv"
	"time"
)

const VERI = "qcms"

type AdminTemp struct{
	Admin db.Admin
}

func (p *Entry)Admin_index(w http.ResponseWriter, r *http.Request) {
	temp := AdminTemp{}
	adminIdCookie,err := r.Cookie("adminId")
	if err != nil{
		fmt.Fprintf(w, Exec_script("alert('登录失败');history.back();"))
		return
	}
	temp.Admin.Id, _ = strconv.Atoi(adminIdCookie.Value)
	adminUsernameCookie,_ := r.Cookie("adminName")
	temp.Admin.Username = adminUsernameCookie.Value
	t, _ := template.ParseFiles("templates/admin/index.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
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

