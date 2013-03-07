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

func (sqlClass SqlType)HaveIp(condStr string) int {
	rows, _ := sqlClass.Conn.Query("SELECT COUNT(*) AS count FROM ad_ip WHERE "+condStr)
	var count int
	for rows.Next(){
    		rows.Scan(&count);
	}
	return count
}
