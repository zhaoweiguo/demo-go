package main

import (
	"flag"
	"fmt"

	"github.com/zhaoweiguo/demo-go/github.com/tal-tech/go-zero/demo1-log/internal/config"
	"github.com/zhaoweiguo/demo-go/github.com/tal-tech/go-zero/demo1-log/internal/handler"
	"github.com/zhaoweiguo/demo-go/github.com/tal-tech/go-zero/demo1-log/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"

	_ "net/http/pprof"
)

var configFile = flag.String("f", "github.com/tal-tech/go-zero/demo1-log/etc/main-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
