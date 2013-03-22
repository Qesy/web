package controllers

import (
	"net/http"
	"fmt"
	"html/template"
	"time"
)

func (p *Entry)Home_index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<font color='#222222'><b>酒悦久广告系统! (Golang 版)</b></font>")
}

func (p *Entry)Home_ad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if(len(p.Params) == 0 || p.Params[0] == "" || p.Params[0] == "0"){
		fmt.Fprintf(w, "Get err !")
		return
	}	
	temp := tempInterface{}
	temp["cTemp"] = Siteinfo()
	adRs := ConnDb.GetAd(" WHERE id = "+p.Params[0]+"", "*", "")
	temp["adRs"] = adRs[0]
	t, _ := template.ParseFiles("templates/home/ad.html")
	t.Execute(w, temp)
}

func (p *Entry)Aaa_bbb(w http.ResponseWriter, r *http.Request){
	//timeNow := time.Now()
	//timeNow := time.Format(time.Now().Unix())
	//fmt.Println(timeNow)
	//fmt.Println(month)
	//fmt.Println(day)
	//t := time.Now().Unix()
	//fmt.Println(t)
	//t, err := time.Parse("2006 01month 02day", "1362984425")
	//t := time.Now(1362984425)
	t := time.Unix(1362984425, 0000000000)
	//fmt.Println(err)
	fmt.Println(t)
	//fmt.Println(time.Unix(t, 0).String())
	//fmt.Println(time.Now().Format("2006year 01month 02day"))
}
