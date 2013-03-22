package controllers

import (
	"net/http"
	"fmt"
	"strconv"
	"encoding/json"
	"common"
	"db"
)

type itemps map[string]db.IpCount
type rtemps map[string]db.RegCount
type otemps map[string]db.OrderCount

func (p *Entry)Ajax_adlist(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")
	adid := r.FormValue("adid")
	temp := tempInterface{}
	//-- 验证登录 --
	_,err := r.Cookie("adminId")
	if err != nil{
		fmt.Fprintf(w, Exec_script("alert('登录失败');history.back();"))
		return
	}
	temp["count"] = ConnDb.GetCount("", "")
	offset := 20
	page := 0
	if len(p.Params) > 0 {
		page, _ = strconv.Atoi(p.Params[0])
	}
	temp["page"] = page
	var condStr string
	if adid == ""{
		condStr = ""
	}else{
		condStr = " WHERE id = "+adid+""
	}
	adRs := ConnDb.GetAd(condStr, "*", "LIMIT "+strconv.Itoa(page * offset)+", "+strconv.Itoa(offset)+"")
	if len(adRs) == 0 {
		temp["adlist"] = ""
		b, _ := json.Marshal(temp)
    	fmt.Fprintf(w, string(b))
    	return
	}
	var idStr string
	for _, v := range adRs{
		idStr += strconv.Itoa(v.Id)+","
	}
	idStrN := common.SubStr(idStr, 0, len(idStr)-2)

	//-- ip & pv --
	ipTemp := itemps{}
	ipRs := ConnDb.GetIpCount(idStrN, start, end)
	for _, v := range ipRs{
		ipTemp[v.Adid] = v
	}

	//-- reg --
	regTemp := rtemps{}
	regRs := ConnDb.GetRegCount(idStrN, start, end)
	for _, v := range regRs{
		regTemp[v.Adid] = v
	}

	//-- order --
	orderTemp := otemps{}
	orderRs := ConnDb.GetOrderCount(idStrN, start, end)
	for _, v := range orderRs{
		orderTemp[v.Adid] = v
	}

	for k, v := range adRs{
		adRs[k].IpCountArr = ipTemp[strconv.Itoa(v.Id)]
		adRs[k].RegCountArr = regTemp[strconv.Itoa(v.Id)]
		adRs[k].OrderCountArr = orderTemp[strconv.Itoa(v.Id)]
	}
	temp["adlist"] = adRs
	b, _ := json.Marshal(temp)
    fmt.Fprintf(w, string(b))
}

func (p *Entry)Ajax_ipreport(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")
	adid := r.FormValue("adid")
	//-- 验证登录 --
	_,err := r.Cookie("adminId")
	if err != nil{
		fmt.Fprintf(w, Exec_script("alert('登录失败');history.back();"))
		return
	}
	var condStr string
	if adid == ""{
		condStr = ""
	}else{
		condStr = " adid = "+adid+" AND "
	}
	adRs := ConnDb.GetIPAll(condStr, start, end)
	if len(adRs) == 0 {
		b, _ := json.Marshal("")
    	fmt.Fprintf(w, string(b))
		return
	}
	b, _ := json.Marshal(adRs)
    fmt.Fprintf(w, string(b))
}

func (p *Entry)Ajax_regreport(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")
	adid := r.FormValue("adid")
	//-- 验证登录 --
	_,err := r.Cookie("adminId")
	if err != nil{
		fmt.Fprintf(w, Exec_script("alert('登录失败');history.back();"))
		return
	}
	var condStr string
	if adid == ""{
		condStr = ""
	}else{
		condStr = " adid = "+adid+" AND "
	}
	adRs := ConnDb.GetRegByTime(condStr, start, end)
	if len(adRs) == 0 {
		b, _ := json.Marshal("")
    	fmt.Fprintf(w, string(b))
		return
	}
	b, _ := json.Marshal(adRs)
    fmt.Fprintf(w, string(b))
}

func (p *Entry)Ajax_loginreport(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")
	adid := r.FormValue("adid")
	//-- 验证登录 --
	_,err := r.Cookie("adminId")
	if err != nil{
		fmt.Fprintf(w, Exec_script("alert('登录失败');history.back();"))
		return
	}
	var condStr string
	if adid == ""{
		condStr = ""
	}else{
		condStr = " adid = "+adid+" AND "
	}
	adRs := ConnDb.GetLoginByTime(condStr, start, end)
	if len(adRs) == 0 {
		b, _ := json.Marshal("")
    	fmt.Fprintf(w, string(b))
		return
	}
	b, _ := json.Marshal(adRs)
    fmt.Fprintf(w, string(b))
}

func (p *Entry)Ajax_buyreport(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")
	adid := r.FormValue("adid")
	//-- 验证登录 --
	_,err := r.Cookie("adminId")
	if err != nil{
		fmt.Fprintf(w, Exec_script("alert('登录失败');history.back();"))
		return
	}
	var condStr string
	if adid == ""{
		condStr = ""
	}else{
		condStr = " adid = "+adid+" AND "
	}
	adRs := ConnDb.GetOrderByTime(condStr, start, end)
	if len(adRs) == 0 {
		b, _ := json.Marshal("")
    	fmt.Fprintf(w, string(b))
		return
	}
	b, _ := json.Marshal(adRs)
    fmt.Fprintf(w, string(b))
}