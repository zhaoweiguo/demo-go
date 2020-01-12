package main

import (
	"log"

	"github.com/zhaoweiguo/demo-go/github.com/bradrydzewski/togo/http"
	"github.com/zhaoweiguo/demo-go/github.com/bradrydzewski/togo/sql"
)

func main() {
	leftpadJs := http.MustLookup("/leftpad.js")
	userSelectSql := sql.Lookup("user-select-all")
	log.Println(string(leftpadJs))
	log.Println("======================================")
	log.Println(userSelectSql)
}
