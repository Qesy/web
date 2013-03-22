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
	//"crypto/md5"
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
	t.Execute(w, temp)
}

func (p *Entry)Admin_list(w http.ResponseWriter, r *http.Request) {
	temp := tempInterface{}
	adminId,err := r.Cookie("adminId")
	adminName,_ := r.Cookie("adminName")
	if err != nil{
		fmt.Fprintf(w, Exec_script("alert('登录失败');history.back();"))
		return
	}
	temp["adminId"], _ = strconv.Atoi(adminId.Value)
	temp["adminName"] = adminName.Value

	temp["userRs"] = ConnDb.GetUser("", "*", "")
	t, _ := template.ParseFiles("templates/admin/list.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
	t.Execute(w, temp)
}

func (p *Entry)Admin_add(w http.ResponseWriter, r *http.Request) {
	temp := tempInterface{}
	adminId,err := r.Cookie("adminId")
	adminName,_ := r.Cookie("adminName")
	if err != nil{
		fmt.Fprintf(w, Exec_script("alert('登录失败');history.back();"))
		return
	}
	if r.Method == "POST"{
		/*h := md5.New()
		io.WriteString(h, r.FormValue("password"))
		password := h.Sum(nil)
		fmt.Printf("INSERT ad_admin SET username=%s, password=%s, permissions=1", r.FormValue("username"), password)
		return*/
		stmt, _ := ConnDb.Conn.Prepare("INSERT ad_admin SET username=?, password=?, permissions=?")
    	_, err = stmt.Exec(r.FormValue("username"), r.FormValue("password"), 1)
	    if err != nil {
	    	fmt.Fprintf(w, Exec_script("alert('添加失败');history.back();"))
	    	return
	    }else{
	    	fmt.Fprintf(w, Exec_script("alert('添加成功');window.location.href='/admin/add';"))
	    	return
	    }
	}
	temp["adminId"], _ = strconv.Atoi(adminId.Value)
	temp["adminName"] = adminName.Value
	t, _ := template.ParseFiles("templates/admin/add.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
	t.Execute(w, temp)
}

func (p *Entry)Admin_edit(w http.ResponseWriter, r *http.Request) {
	id := p.Params[0]
	temp := tempInterface{}
	adminId,err := r.Cookie("adminId")
	adminName,_ := r.Cookie("adminName")
	if err != nil{
		fmt.Fprintf(w, Exec_script("alert('登录失败');history.back();"))
		return
	}
	if r.Method == "POST"{
		stmt, _ := ConnDb.Conn.Prepare("UPDATE ad_admin SET username=?, password=? WHERE id=?")
    	_, err = stmt.Exec(r.FormValue("username"), r.FormValue("password"), id)
	    if err != nil {
	    	fmt.Fprintf(w, Exec_script("alert('修改失败');history.back();"))
	    	return
	    }else{
	    	fmt.Fprintf(w, Exec_script("alert('修改成功');window.location.href='/admin/list';"))
	    	return
	    }
	}
	temp["adminId"], _ = strconv.Atoi(adminId.Value)
	temp["adminName"] = adminName.Value
	temp["userRs"] = ConnDb.GetUser(" WHERE id = "+id+" ", "*", "")
	t, _ := template.ParseFiles("templates/admin/edit.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
	t.Execute(w, temp)
}

func (p *Entry)Admin_log(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	temp := tempInterface{}
	adminId,err := r.Cookie("adminId")
	adminName,_ := r.Cookie("adminName")
	if err != nil{
		fmt.Fprintf(w, Exec_script("alert('登录失败');history.back();"))
		return
	}
	temp["adminId"], _ = strconv.Atoi(adminId.Value)
	temp["adminName"] = adminName.Value
	temp["log"] = "你的访问路径： "+r.URL.Path+",查看日志请联系超级管理员。"
	t, _ := template.ParseFiles("templates/admin/log.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
	t.Execute(w, temp)
}

func (p *Entry)Admin_login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.Method == "POST"{
		if VERI != r.FormValue("veri"){
			fmt.Fprintf(w, Exec_script("alert('验证码错误');history.back();"))
			return
		}
		userInfo := ConnDb.GetUser(" WHERE username='"+r.FormValue("username")+"' && password='"+r.FormValue("password")+"'", "*", "limit 0,1")
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

