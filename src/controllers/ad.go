package controllers

import (
	"net/http"
	"text/template"
	"fmt"
	"time"
	"strconv"
)

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
	temp["adminId"], _ = strconv.Atoi(adminId.Value)
	temp["adminName"] = adminName.Value
	t.Execute(w, temp)
}

func (p *Entry)Ad_add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		timeInt := strconv.Itoa(int(time.Now().Unix()))
        stmt, _ := ConnDb.Conn.Prepare("INSERT ad_list SET name=?, img=?, url=?, width=?, height=?, addtime=?")
	    res, _ := stmt.Exec(r.FormValue("name"), r.FormValue("img"), r.FormValue("url"), r.FormValue("width"), r.FormValue("height"),timeInt)
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

func (p *Entry)Ad_edit(w http.ResponseWriter, r *http.Request) {
	id := p.Params[0]
	if r.Method == "POST"{
        stmt, _ := ConnDb.Conn.Prepare("UPDATE ad_list SET name=?, img=?, url=?, width=?, height=? where id = ?")
	    _, err := stmt.Exec(r.FormValue("name"), r.FormValue("img"), r.FormValue("url"), r.FormValue("width"), r.FormValue("height"), id)
	    if err != nil{
	    	fmt.Fprintf(w, Exec_script("alert('修改失败');window.location.href='/ad'"))
	    	return
	    }
	    fmt.Fprintf(w, Exec_script("alert('修改成功');window.location.href='/ad'"))
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
	t, _ := template.ParseFiles("templates/ad/edit.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
	adRs := ConnDb.GetAd(" WHERE id = "+id+"", "*", "")
	temp["AdOne"] = adRs[0]
	t.Execute(w, temp)	
}

func (p *Entry)Ad_code(w http.ResponseWriter, r *http.Request) {
	id := p.Params[0]
	temp := tempInterface{}
	temp["cTemp"] = Siteinfo()
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

//-- 以下为报表 --

func (p *Entry)Ad_all(w http.ResponseWriter, r *http.Request) {
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
	
	temp["adCount"] = ConnDb.GetCount("", "")
	t, _ := template.ParseFiles("templates/ad/all.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
	t.Execute(w, temp)
}

func (p *Entry)Ad_ipreport(w http.ResponseWriter, r *http.Request) {
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
	t, _ := template.ParseFiles("templates/ad/ipreport.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
	t.Execute(w, temp)
}

func (p *Entry)Ad_pvreport(w http.ResponseWriter, r *http.Request) {
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
	t, _ := template.ParseFiles("templates/ad/pvreport.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
	t.Execute(w, temp)
}

func (p *Entry)Ad_regreport(w http.ResponseWriter, r *http.Request) {
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
	t, _ := template.ParseFiles("templates/ad/regreport.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
	t.Execute(w, temp)
}

func (p *Entry)Ad_loginreport(w http.ResponseWriter, r *http.Request) {
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
	t, _ := template.ParseFiles("templates/ad/loginreport.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
	t.Execute(w, temp)
}

func (p *Entry)Ad_buyreport(w http.ResponseWriter, r *http.Request) {
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
	t, _ := template.ParseFiles("templates/ad/buyreport.html", "templates/header.html", "templates/lefter.html", "templates/footer.html")
	t.Execute(w, temp)
}
