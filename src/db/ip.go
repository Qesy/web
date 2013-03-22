package db

import(
	//"fmt"
)


type Ip struct{
	Id		int
	Adid	int
	Ip 		int
	Time	int
	Type 	int
	Domain	string
	Refer	string
}

type IpCount struct{
	Adid 		string
	IpCount     string
	PvCount		string
	TimeCount	string
}

func (sqlClass SqlType)HaveIp(condStr string) int {
	rows, _ := sqlClass.Conn.Query("SELECT COUNT(*) AS count FROM ad_ip WHERE "+condStr)
	var count int
	for rows.Next(){
    		rows.Scan(&count);
	}
	return count
}

func (sqlClass SqlType)GetIp(condStr string, field string, limit string) []Ip {
	var adIp Ip
	var adIps []Ip
	adrows, _ := sqlClass.Conn.Query("SELECT "+ field +" FROM ad_ip "+condStr+" "+limit)
    var id int
    var adid int
    var ip int
    var time int
    var adtype int
    var domain string
    var refer string
    for adrows.Next(){
    		adrows.Scan(&id,&adid,&ip,&time,&adtype,&domain,&refer)
    		adIp.Id = id
			adIp.Adid = adid
			adIp.Ip = ip
			adIp.Time = time
			adIp.Type = adtype
			adIp.Domain = domain
			adIp.Refer = refer
			adIps =append(adIps,adIp)
	}
	return adIps
}

func (sqlClass SqlType)GetIpCount(adidStr string, start string, end string) []IpCount {
	var ipStruct IpCount
	var ipStructs []IpCount
	adrows, _ := sqlClass.Conn.Query("SELECT adid, COUNT(*) AS count, SUM(pv) AS sum FROM ad_ip WHERE adid in ("+adidStr+") AND time >= "+start+" AND time < "+end+" GROUP BY adid")
	var adid, count, sum string
	for adrows.Next(){
		adrows.Scan(&adid, &count,&sum, )
		ipStruct.Adid = adid
		ipStruct.IpCount = count
		ipStruct.PvCount = sum
		ipStructs =append(ipStructs,ipStruct)
	}
	return ipStructs
}

func (sqlClass SqlType)GetIPAll(condStr string, start string, end string) []IpCount {
	var ipStruct IpCount
	var ipStructs []IpCount
	adrows, _ := sqlClass.Conn.Query("SELECT adid, COUNT(*) AS count, SUM(pv) AS sum, time FROM ad_ip WHERE "+condStr+" time >= "+start+" AND time < "+end+" GROUP BY time ORDER BY time ASC")
	var adid, count, sum, time string
	for adrows.Next(){
		adrows.Scan(&adid, &count,&sum, &time)
		ipStruct.Adid = adid
		ipStruct.IpCount = count
		ipStruct.PvCount = sum
		ipStruct.TimeCount = time
		ipStructs =append(ipStructs,ipStruct)
	}
	return ipStructs
}