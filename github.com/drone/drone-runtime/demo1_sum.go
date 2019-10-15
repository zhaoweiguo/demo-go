package main

import (
	"context"
	"github.com/drone/drone-runtime/engine"
	"github.com/drone/drone-runtime/engine/docker"
	"github.com/drone/drone-runtime/runtime"
	"github.com/drone/drone-runtime/runtime/term"
	"github.com/drone/signal"
	"log"
	"os"
	"time"
)

func main() {
	source := "./github.com/drone/drone-runtime/usage/1_hello_world.json"
	// 解析json文件
	config, err := engine.ParseFile(source)
	if err != nil {
		log.Fatalln(err)
	}

	// 设定docker client
	var engine engine.Engine
	engine, err = docker.NewEnv()
	if err != nil {
		log.Fatalln(err)
	}

	hooks := &runtime.Hook{}
	hooks.GotLine = term.WriteLinePretty(os.Stdout)

	r := runtime.New(
		runtime.WithEngine(engine),
		runtime.WithConfig(config),
		runtime.WithHooks(hooks),
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	ctx = signal.WithContext(ctx)
	defer cancel()

	err = r.Run(ctx)
	if err != nil {
		log.Fatalln(err)
	}

}
