package db

import(
	//"fmt"
)

type Ad struct{
	Id int
	Name string
	Img string
	Url string
	Width int
	Height int
	Status int
	Addtime int
	Count int
	IpCountArr IpCount
	RegCountArr RegCount
	OrderCountArr OrderCount
}

func (sqlClass SqlType)GetAd(condStr string, field string, limit string) []Ad {
	var user Ad
	var users []Ad
	adrows, _ := sqlClass.Conn.Query("SELECT "+ field +" FROM ad_list "+condStr+" "+limit)
    var id int
    var name string
    var img string
    var url string
    var width int
    var height int
    var status int
    var addtime int
    for adrows.Next(){
    		adrows.Scan(&id,&name,&img,&url,&width,&height,&status,&addtime);
    		user.Id = id
			user.Name = name
			user.Img = img
			user.Url = url
			user.Width = width
			user.Height = height
			user.Status = status
			user.Addtime = addtime
			users =append(users,user)
	}
	return users
}

func (sqlClass SqlType)GetCount(condStr string, limit string) int {
	var user Ad
	adrows, _ := sqlClass.Conn.Query("SELECT COUNT(*) AS count FROM ad_list "+condStr+" "+limit)
	var count int
	for adrows.Next(){
    		adrows.Scan(&count);
    		user.Count = count
	}
	return user.Count
}