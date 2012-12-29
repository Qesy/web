package controllers

import (
	"net/http"
	"html/template"
	"fmt"
	"encoding/json"
	"os"
	"bufio"
	"io/ioutil"
)

type Config struct{
	DbInfo Db
	SiteInfo Site
}

type Db struct{
	DbHost	string
	DbName string
	DbUser string
	DbPass string
}

type Site struct{
	Title string
	Email string
	Veri  string
}

func (p *Entry)Config_index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")	
	var cTemp Config
	config, _ := ioutil.ReadFile("config/site")
	jsondecode := json.Unmarshal(config, &cTemp)
	if jsondecode != nil {
		fmt.Fprintf(w, "Config jsondecode err !")
		return
	}
	if r.Method == "POST"{	
		var c Config
		if cTemp.SiteInfo.Veri != "" && cTemp.SiteInfo.Veri != r.FormValue("Veri") {
			fmt.Fprintf(w, "Veri code err !")
			return
		}			
		c.SiteInfo = Site{Title : r.FormValue("Title"), Email : r.FormValue("Email"), Veri : r.FormValue("Veri")}
		c.DbInfo = Db{DbHost : r.FormValue("DbHost"), DbName : r.FormValue("DbName"), DbUser : r.FormValue("DbUser"), DbPass : r.FormValue("DbPass")}
		b, _ := json.Marshal(c)
		file, err := os.Create("config/site")
		if err != nil {
			fmt.Fprintf(w, "writer %s",err)
			return
		}
		defer file.Close() 
		writer := bufio.NewWriter(file)
		writer.Write(b)
		writer.Flush()
		script := Exec_script("alert('Configuration Settings complete !');history.back();")
		fmt.Fprintf(w, script)
		fmt.Println("Configuration Settings complete !")
		return
	}	
	t, err := template.ParseFiles("templates/config/index.html")
	if err != nil{
		fmt.Fprintf(w, "template not find !")
		return
	}
	t.Execute(w, nil)

}