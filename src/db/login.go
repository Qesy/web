package db

import(
	//"fmt"
)


type Login struct{
	Id		int
	Adid	int
	Ip 		int
	Uid 	int
	Time	int
}

type LoginCount struct{
	Adid 		string
	LoginCount     string
	TimeCount	string
}

func (sqlClass SqlType)GetLoginByTime(condStr string, start string, end string) []LoginCount {
	var regStruct LoginCount
	var regStructs []LoginCount
	adrows, _ := sqlClass.Conn.Query("SELECT adid, COUNT(*) AS count, time FROM ad_login WHERE "+condStr+" time >= "+start+" AND time < "+end+" GROUP BY time ORDER BY time ASC")
	var adid, count, time string
	for adrows.Next(){
		adrows.Scan(&adid, &count, &time)
		regStruct.Adid = adid
		regStruct.LoginCount = count
		regStruct.TimeCount = time
		regStructs =append(regStructs,regStruct)
	}
	return regStructs
}