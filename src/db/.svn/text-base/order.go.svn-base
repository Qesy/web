package db

import(
	//"fmt"
)

type Order struct{
	Id		int
	Adid	int
	Uid		int
	order_id int
	money	int
	real_money	int
	Time	int
}

type OrderCount struct{
	Adid 		string
	OrderCount  string
	Money 		string
	TimeCount 	string
}

func (sqlClass SqlType)GetOrderCount(adidStr string, start string, end string) []OrderCount {
	var orderStruct OrderCount
	var orderStructs []OrderCount
	adrows, _ := sqlClass.Conn.Query("SELECT adid, COUNT(*) AS count, sum(money) as sum FROM ad_order WHERE adid in ("+adidStr+") AND time >= "+start+" AND time < "+end+" GROUP BY adid")
	var adid, count, sum string
	for adrows.Next(){
		adrows.Scan(&adid, &count, &sum)
		orderStruct.Adid = adid
		orderStruct.OrderCount = count
		orderStruct.Money = sum
		orderStructs =append(orderStructs,orderStruct)
	}
	return orderStructs
}

func (sqlClass SqlType)GetOrderByTime(condStr string, start string, end string) []OrderCount {
	var regStruct OrderCount
	var regStructs []OrderCount
	adrows, _ := sqlClass.Conn.Query("SELECT adid, COUNT(*) AS count, sum(money) as sum, time FROM ad_order WHERE "+condStr+" time >= "+start+" AND time < "+end+" GROUP BY time ORDER BY time ASC")
	var adid, count, sum, time string
	for adrows.Next(){
		adrows.Scan(&adid, &count, &sum, &time)
		regStruct.Adid = adid
		regStruct.OrderCount = count
		regStruct.Money = sum
		regStruct.TimeCount = time
		regStructs =append(regStructs,regStruct)
	}
	return regStructs
}