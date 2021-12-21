package main

import (
	"context"
	"fmt"
	"log"

	"github.com/coreos/etcd/clientv3/concurrency"
)

var count int

// 选主
func elect(e1 *concurrency.Election, electName string) {
	log.Println("acampaigning for ID:", *nodeID)
	// 调用Campaign方法选主,主的值为value-<主节点ID>-<count>
	if err := e1.Campaign(context.Background(), fmt.Sprintf("value-%d-%d", *nodeID, count)); err != nil {
		log.Println(err)
	}
	log.Println("campaigned for ID:", *nodeID)
	count++
}

// 为主设置新值
func proclaim(e1 *concurrency.Election, electName string) {
	log.Println("proclaiming for ID:", *nodeID)
	// 调用Proclaim方法设置新值,新值为value-<主节点ID>-<count>
	if err := e1.Proclaim(context.Background(), fmt.Sprintf("value-%d-%d", *nodeID, count)); err != nil {
		log.Println(err)
	}
	log.Println("proclaimed for ID:", *nodeID)
	count++
}

// 重新选主，有可能另外一个节点被选为了主
func resign(e1 *concurrency.Election, electName string) {
	log.Println("resigning for ID:", *nodeID)
	// 调用Resign重新选主
	if err := e1.Resign(context.TODO()); err != nil {
		log.Println(err)
	}
	log.Println("resigned for ID:", *nodeID)
}
