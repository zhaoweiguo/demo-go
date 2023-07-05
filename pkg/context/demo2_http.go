package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 问题函数, 在请求完成后生成的进程不会停止
// 这儿:
// 假定请求需耗时3s，即请求在3s后返回，我们期望监控goroutine在打印3次“req is processing”后即停止。
// 但运行发现，监控goroutine打印3次后，其仍不会结束，而会一直打印下去
func main_with_bug() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// monitor
		go func() {
			for range time.Tick(time.Second) {
				fmt.Println("req is processing")
			}
		}()

		// assume req processing takes 3s
		time.Sleep(3 * time.Second)
		w.Write([]byte("hello"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// monitor
		go func() {
			for range time.Tick(time.Second) {
				select {
				case <-r.Context().Done():
					fmt.Println("req is outgoing")
					return
				default:
					fmt.Println("req is processing")
				}
			}
		}()

		// assume req processing takes 3s
		time.Sleep(3 * time.Second)
		w.Write([]byte("hello"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
