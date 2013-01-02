package controllers

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func (p *Entry)Api_ip(w http.ResponseWriter, r *http.Request) {
	type temp map[string]string
	jSon := temp{}
	if r.FormValue("ip") == "" || r.FormValue("adid") == "" || r.FormValue("time") == "" {
		jSon["httpCode"] = "403"	//-- 缺少参数 --
		jSon["body"]	 = "Missing parameter !"
		b, _ := json.Marshal(jSon)
    	fmt.Fprintf(w, string(b))
    	return
	}	
	if r.Header["Token"][0] == token && r.Header["Secret"][0] == secret {
		stmt, _ := ConnDb.Conn.Prepare("INSERT ad_ip SET adid=?, ip=?, time=?, type=?, domain=?, refer=?")
	    _, err := stmt.Exec(r.FormValue("adid"), r.FormValue("ip"), r.FormValue("time"), r.FormValue("type"), r.FormValue("domain"), r.FormValue("refer"))
	    if err != nil {
	    	jSon["httpCode"] = "404"	//-- 保存失败 --
			jSon["body"]	 = "Api save failed !"
	    }else{
	    	jSon["httpCode"] = "200"	//-- 保存成功 --
			jSon["body"]	 = "Api success !"
	    }
	}else{
		jSon["httpCode"] = "400"	//-- 验证失败 --
		jSon["body"]	 = "Api verification failed !"
	}
	b, _ := json.Marshal(jSon)
    fmt.Fprintf(w, string(b))
}