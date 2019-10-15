package main

import (
	"github.com/drone/drone/core"
	"github.com/drone/drone/metric/sink"
	"github.com/drone/drone/operator/runner"
	"github.com/drone/drone/server"
	"github.com/drone/drone/trigger/cron"
)

func main() {

}

// application is the main struct for the Drone server.
type application struct {
	cron   *cron.Scheduler
	sink   *sink.Datadog
	runner *runner.Runner
	server *server.Server
	users  core.UserStore
}

// newApplication creates a new application struct.
func newApplication(
	cron *cron.Scheduler,
	sink *sink.Datadog,
	runner *runner.Runner,
	server *server.Server,
	users core.UserStore) application {
	return application{
		users:  users,
		cron:   cron,
		sink:   sink,
		server: server,
		runner: runner,
	}
}
