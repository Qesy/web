package router

import(
	"strings"
	"controllers"
	"reflect"
	"net/http"
	"fmt"
)

const DEFAULT_CONTROLLER = "home"
const DEFAULT_ACTION	 = "index"

func Fetch_url(url string, w http.ResponseWriter, r *http.Request) {
	e := new(controllers.Entry)
	if(url == ""){
		e.Controller = DEFAULT_CONTROLLER
		e.Action = DEFAULT_ACTION
	}else{
		urlArr := strings.Split(url, "/")
		if len(urlArr) <= 1{
			e.Controller = urlArr[0]
			e.Action = DEFAULT_ACTION
		}else{
			e.Controller = urlArr[0]
			if urlArr[1] == ""{
				e.Action = DEFAULT_ACTION
			}else{
				e.Action = urlArr[1]
			}
			e.Params = urlArr[2:len(urlArr)]
		}
	}	
	fmt.Println(e)
	if e.Controller == "static"{	
		http.ServeFile(w, r , url)
		return
	}	
	fmt.Println(e)
	t := reflect.TypeOf(e)
	action := strings.Title(e.Controller)+"_"+e.Action
    m, ok := t.MethodByName(action)
    if !ok {
    	e.ErrorStatus(w, r, 404)
        return
    }
    preFunc()
    m.Func.Call([]reflect.Value{reflect.ValueOf(e), reflect.ValueOf(w), reflect.ValueOf(r)})
    nextFunc()
}

func preFunc() {
	//-- 连接数据库 --	
	controllers.DbConn()
}

func nextFunc() {
	//-- 关闭数据库 --
	controllers.DbDisconn()
}
