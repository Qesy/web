package main

import (
	"net/http"
	"router"
	"common"
)

type MyMux struct {}

func main() {
	mux := &MyMux{}	
    http.ListenAndServe(":12345", mux)
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := common.SubStr(r.URL.Path, 1, len(r.URL.Path))
	router.Fetch_url(path, w, r)
}
