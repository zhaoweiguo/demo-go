package httpexpect

import (
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"testing"
	"time"
)

func init() {
	go start_a_server()
	time.Sleep(time.Second)
}

func TestSimple(t *testing.T) {
	e := httpexpect.New(t, "http://127.0.0.1:8080")
	e.GET("/hello").
		Expect().
		Status(http.StatusOK).Body().Contains("hello")
}

func TestBaidu(t *testing.T) {
	e := httpexpect.New(t, "https://baidu.com")
	e.GET("/search/error.html").Expect().Status(http.StatusOK)
}

func start_a_server() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{hello: world}"))
}
