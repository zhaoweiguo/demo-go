package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


/*
//测试:
sh>curl http://127.0.0.1:7800/abc/123/def
//结果:(后台打印)
start...
map[id:123]
123
*/
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/abc/{id}/def", Handler)    // http://127.0.0.1:7800/abc/123/def
	r.HandleFunc("/test/{category}/{id:[0-9]+}", Handler)  // :7800/test/<category>/<id>
	http.Handle("/", r)
	addr := ":7800"
	server := &http.Server{
		Addr:           addr,
		Handler:        r,
	}
	server.ListenAndServe()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start...")
	vars := mux.Vars(r)
	fmt.Println(vars)
	id := vars["id"]
	fmt.Println(id)
}


