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
		result := ip(r.FormValue("adid"), r.FormValue("ip"), r.FormValue("time"), r.FormValue("type"), r.FormValue("domain"), r.FormValue("refer"), r.FormValue("ua"), r.FormValue("goods_id"))
		if result == 0 {
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

func (p *Entry)Api_reg(w http.ResponseWriter, r *http.Request) {
	type temp map[string]string
	jSon := temp{}
	if r.FormValue("ip") == "" || r.FormValue("adid") == "" || r.FormValue("time") == "" || r.FormValue("uid") == "" {
		jSon["httpCode"] = "403"	//-- 缺少参数 --
		jSon["body"]	 = "Missing parameter !"
		b, _ := json.Marshal(jSon)
    	fmt.Fprintf(w, string(b))
    	return
	}	
	if r.Header["Token"][0] == token && r.Header["Secret"][0] == secret {
		result := reg(r.FormValue("adid"), r.FormValue("ip"), r.FormValue("uid"), r.FormValue("time"))
		if result == 0 {
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

func (p *Entry)Api_order(w http.ResponseWriter, r *http.Request) {
	type temp map[string]string
	jSon := temp{}
	if r.FormValue("adid") == "" || r.FormValue("time") == "" || r.FormValue("uid") == "" || r.FormValue("order_id") == "" || r.FormValue("money") == "" || r.FormValue("real_money") == "" {
		jSon["httpCode"] = "403"	//-- 缺少参数 --
		jSon["body"]	 = "Missing parameter !"
		b, _ := json.Marshal(jSon)
    	fmt.Fprintf(w, string(b))
    	return
	}	
	if r.Header["Token"][0] == token && r.Header["Secret"][0] == secret {
		result := order(r.FormValue("adid"), r.FormValue("uid"), r.FormValue("order_id"), r.FormValue("money"), r.FormValue("real_money"), r.FormValue("time"))
		if result == 0 {
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

func (p *Entry)Api_mail(w http.ResponseWriter, r *http.Request) {
	type temp map[string]string
	jSon := temp{}
	if r.FormValue("mail_id") == "" || r.FormValue("time") == "" || r.FormValue("uid") == "" || r.FormValue("ip") == "" || r.FormValue("goods_id") == "" {
		jSon["httpCode"] = "403"	//-- 缺少参数 --
		jSon["body"]	 = "Missing parameter !"
		b, _ := json.Marshal(jSon)
    	fmt.Fprintf(w, string(b))
    	return
	}	
	if r.Header["Token"][0] == token && r.Header["Secret"][0] == secret {
		result := mail(r.FormValue("uid"), r.FormValue("mail_id"), r.FormValue("goods_id"), r.FormValue("ip"), r.FormValue("time"))
		if result == 0 {
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

func (p *Entry)Api_activity(w http.ResponseWriter, r *http.Request) {
	type temp map[string]string
	jSon := temp{}
	if r.FormValue("activity_id") == "" || r.FormValue("time") == "" || r.FormValue("uid") == "" || r.FormValue("ip") == "" || r.FormValue("goods_id") == "" {
		jSon["httpCode"] = "403"	//-- 缺少参数 --
		jSon["body"]	 = "Missing parameter !"
		b, _ := json.Marshal(jSon)
    	fmt.Fprintf(w, string(b))
    	return
	}	
	if r.Header["Token"][0] == token && r.Header["Secret"][0] == secret {
		result := activity(r.FormValue("uid"), r.FormValue("activity_id"), r.FormValue("goods_id"), r.FormValue("ip"), r.FormValue("time"))
		if result == 0 {
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

func ip(adid string, ip string, time string, adType string, domian string, refer string, ua string, goods_id string)int {
	var result int
	var err error
	countIp := ConnDb.HaveIp("ip="+ip+" && time = '"+time+"' && adid="+adid+"")
	if countIp != 0 {
		stmt, _ := ConnDb.Conn.Prepare("UPDATE ad_ip SET pv=pv+1 where adid = ? && ip="+ip+" && time = '"+time+"'")
		_, err = stmt.Exec(adid)
	}else{
		stmt, _ := ConnDb.Conn.Prepare("INSERT ad_ip SET adid=?, ip=?, time=?, type=?, domain=?, refer=?, ua=?, goods_id=?")
    	_, err = stmt.Exec(adid, ip, time, adType, domian, refer, ua, goods_id)
	}
    if err != nil {
    	result = 0
    }else{
    	result = 1
    }
    return result
}

func reg(adid string, ip string, uid string, time string)int {
	var result int
	var err error
	stmt, _ := ConnDb.Conn.Prepare("INSERT ad_reg SET adid=?, ip=?, uid=?, time=?")
    _, err = stmt.Exec(adid, ip, uid, time)
	if err != nil {
    	result = 0
    }else{
    	result = 1
    }
    return result
}

func order(adid string, uid string, order_id string, money string, real_money string, time string) int {
	var result int
	var err error
	stmt, _ := ConnDb.Conn.Prepare("INSERT ad_order SET adid=?, uid=?, order_id=?, money=?, real_money=?, time=?")
    _, err = stmt.Exec(adid, uid, order_id, money, real_money, time)
	if err != nil {
    	result = 0
    }else{
    	result = 1
    }
    return result
}

func mail(uid string, mail_id string, goods_id string, ip string, time string) int {
	var result int
	var err error
	stmt, _ := ConnDb.Conn.Prepare("INSERT ad_email SET uid=?, mail_id=?, goods_id=?, ip=?, time=?")
    _, err = stmt.Exec(uid, mail_id, goods_id, ip, time)
	if err != nil {
    	result = 0
    }else{
    	result = 1
    }
    return result
}

func activity(uid string, activity_id string, goods_id string, ip string, time string) int {
	var result int
	var err error
	stmt, _ := ConnDb.Conn.Prepare("INSERT ad_activity SET uid=?, activity_id=?, goods_id=?, ip=?, time=?")
    _, err = stmt.Exec(uid, activity_id, goods_id, ip, time)
	if err != nil {
    	result = 0
    }else{
    	result = 1
    }
    return result
}