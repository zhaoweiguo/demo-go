package main

import (
	"context"
	client_go "github.com/zhaoweiguo/demo-go/k8s.io/client-go"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	clientSet := client_go.GetClientset()
	verInfo, err := clientSet.ServerVersion()
	if err != nil {
		log.Panic(err)
	}
	log.Println(verInfo)
	ctx := context.Background()
	clientSet.CoreV1().Pods("default").Watch(ctx)
}
