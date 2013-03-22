package db

import(
	"fmt"
)

type Admin struct{
	Id int
	Username string
	Password string
	Permissions int
}

func (sqlClass SqlType)GetUser(condStr string, field string, limit string) []Admin {
	var user Admin
	var users []Admin
	rows, err := sqlClass.Conn.Query("SELECT "+ field +" FROM ad_admin "+condStr+" "+limit)
	if err != nil {
		fmt.Println(sqlClass)
		fmt.Println(err)
		return users
	}
    var id int
    var username string
    var password string
    var permissions int
    for rows.Next(){
    		rows.Scan(&id,&username,&password,&permissions);
    		user.Id = id
			user.Username = username
			user.Password = password
			user.Permissions = permissions
			users =append(users,user)
	}
	return users
}

func (sqlClass SqlType)Add() {
	
}

func (sqlClass SqlType)HaveUser(condStr string) int {
	rows, _ := sqlClass.Conn.Query("SELECT COUNT(*) AS count FROM ad_admin WHERE "+condStr)
	var count int
	for rows.Next(){
    		rows.Scan(&count);
	}
	return count
}