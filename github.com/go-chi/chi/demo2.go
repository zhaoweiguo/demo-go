package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	/*
		1. 不加此条语句，日志打印结果为:
		2019/09/23 16:45:41 "GET http://127.0.0.1:3333/ HTTP/1.1" from 127.0.0.1:61625 - 200 11B in 16.685µs

		2. 加了此条语句，日志打印结果为:
		2019/09/23 16:46:54 [zhaowgMac/jYcmMU4Ucc-000001] "GET http://127.0.0.1:3333/ HTTP/1.1" from 127.0.0.1:61638 - 200 11B in 16.999µs
	*/
	r.Use(middleware.RequestID) // 记录请求id

	r.Use(middleware.Logger) // 打开日志
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":3333", r)
}
