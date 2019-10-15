//+build wireinject

package main

import (
	"github.com/drone/drone/cmd/drone-server/config"
	"github.com/google/wire"
)

func InitializeApplication(config config.Config) (application, error) {
	wire.Build(
		clientSet,
		licenseSet,
		loginSet,
		pluginSet,
		runnerSet,
		schedulerSet,
		serverSet,
		serviceSet,
		storeSet,
		newApplication,
	)
	return application{}, nil
}
