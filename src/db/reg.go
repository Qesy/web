package db

import(
	//"fmt"
)

type Reg struct{
	Id		int
	Adid	int
	Ip 		int
	Uid		int
	Time	int
}

type RegCount struct{
	Adid 		string
	RegCount     string
	TimeCount 	string
}

func (sqlClass SqlType)GetRegCount(adidStr string, start string, end string) []RegCount {
	var regStruct RegCount
	var regStructs []RegCount
	adrows, _ := sqlClass.Conn.Query("SELECT adid, COUNT(*) AS count FROM ad_reg WHERE adid in ("+adidStr+") AND time >= "+start+" AND time < "+end+" GROUP BY adid")
	var adid, count string
	for adrows.Next(){
		adrows.Scan(&adid, &count)
		regStruct.Adid = adid
		regStruct.RegCount = count
		regStructs =append(regStructs,regStruct)
	}
	return regStructs
}

func (sqlClass SqlType)GetRegByTime(condStr string, start string, end string) []RegCount {
	var regStruct RegCount
	var regStructs []RegCount
	adrows, _ := sqlClass.Conn.Query("SELECT adid, COUNT(*) AS count, time FROM ad_reg WHERE "+condStr+" time >= "+start+" AND time < "+end+" GROUP BY time ORDER BY time ASC")
	var adid, count, time string
	for adrows.Next(){
		adrows.Scan(&adid, &count, &time)
		regStruct.Adid = adid
		regStruct.RegCount = count
		regStruct.TimeCount = time
		regStructs =append(regStructs,regStruct)
	}
	return regStructs
}