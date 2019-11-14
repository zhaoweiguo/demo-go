package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io"
	"log"
	"net/http"
)

/*
启动后, 浏览器访问:
1. http://127.0.0.1:4000/hello
hello
2. http://127.0.0.1:4000/healthz/hello
another hello
*/
func main() {
	router := get_router()
	addr := "127.0.0.1:4000"
	fmt.Printf("server is starting at %s", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Println(err)
	}
}

func get_router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/hello", Hello())
	healthz := anotherHandler()
	r.Mount("/healthz", healthz) // 合并其他http.Handler or chi Router到本路由路径(方便在大项目拆分api到多个子项目中用)

	return r
}

// New returns a new health check router.
func anotherHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)  // 使用中间件
	r.Handle("/hello", Hello2()) // 匹配所有的http method
	return r
}

func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func Hello2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "another hello")
	}
}

func Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "hello")
	}
}
