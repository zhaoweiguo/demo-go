package main

import (
	"context"
	"log"

	"github.com/coreos/etcd/clientv3/concurrency"
)

// 查询主的信息
func query(e1 *concurrency.Election, electName string) {
	// 调用Leader返回主的信息，包括key和value等信息
	resp, err := e1.Leader(context.Background())
	if err != nil {
		log.Printf("failed to get the current leader: %v", err)
	}
	log.Println("current leader:", string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
}

// 可以直接查询主的rev信息
func rev(e1 *concurrency.Election, electName string) {
	rev := e1.Rev()
	log.Println("current rev:", rev)
}

func watch(e1 *concurrency.Election, electName string) {
	ch := e1.Observe(context.TODO())
	log.Println("start to watch for ID:", *nodeID)
	for i := 0; i < 10; i++ {
		resp := <-ch
		log.Println("leader changed to", string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
	}
}
