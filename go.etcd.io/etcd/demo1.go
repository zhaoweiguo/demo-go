package main

import (
	"go.etcd.io/etcd/clientv3"
	etcdnaming "go.etcd.io/etcd/clientv3/naming"
	"log"

	"google.golang.org/grpc"
)

/*
https://github.com/etcd-io/etcd/blob/master/Documentation/dev-guide/grpc_naming.md
*/
func main() {
	cli, cerr := clientv3.NewFromURL("http://localhost:2379")
	r := &etcdnaming.GRPCResolver{Client: cli}
	b := grpc.RoundRobin(r)
	conn, gerr := grpc.Dial("my-service", grpc.WithBalancer(b), grpc.WithBlock())

	log.Println(cerr, gerr, conn)
}
