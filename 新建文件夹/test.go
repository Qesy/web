package main

import(
	//"router"
	"fmt"
	//"reflect"
	"strings"
)

type MyStruct struct{
	name string
}

func (this *MyStruct)GetName() {
	fmt.Println(this.name)
    //return this.name
}

func main() {
	a := "asdfasdfasf"
	fmt.Println(strings.Title(a))
	return
	//url := "Aaa/bbb/ccc"
	//url := ""
	//router.Fetch_url(url)

	//SubStr("abcdefg", 1, 2)
	//funcs := map[string]func() {"foobar":"foobar"}
	//funcs["foobar"]()
	//a := new(MyStruct)
	//var a MyStruct
	//a.name = "yejianfeng"
	//fmt.Println(a)
	//return
	//reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
    //fmt.Println(b)
	//reflect.ValueOf(a).MethodByName("foobar").Call([]reflect.Value{})
}



func SubStr(str string, start int, end int) string{
	var s string
	c := []byte(str)
	for k, v := range c{
		if(k >= start && k <= end){
			s += string(v)
		}
	}
	return s
}