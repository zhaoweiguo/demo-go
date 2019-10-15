package main

import (
	"github.com/drone/drone-runtime/engine/docker"
	"github.com/drone/drone/core"
	"github.com/drone/drone/operator/manager"
	"github.com/drone/drone/operator/runner"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"github.com/zhaoweiguo/demo-go/github.com/drone/drone/wire/config"
)

// wire set for loading the server.
var runnerSet = wire.NewSet(
	provideRunner,
)

// provideRunner is a Wire provider function that returns a
// local build runner configured from the environment.
func provideRunner(
	manager manager.BuildManager,
	secrets core.SecretService,
	registry core.RegistryService,
	config config.Config,
) *runner.Runner {
	// the local runner is only created when the nomad scheduler,
	// kubernetes scheduler, and remote agents are disabled
	if config.Nomad.Enabled || config.Kube.Enabled || (config.Agent.Disabled == false) {
		return nil
	}
	engine, err := docker.NewEnv()
	if err != nil {
		logrus.WithError(err).
			Fatalln("cannot load the docker engine")
		return nil
	}
	return &runner.Runner{
		Platform:   config.Runner.Platform,
		OS:         config.Runner.OS,
		Arch:       config.Runner.Arch,
		Kernel:     config.Runner.Kernel,
		Variant:    config.Runner.Variant,
		Engine:     engine,
		Manager:    manager,
		Secrets:    secrets,
		Registry:   registry,
		Volumes:    config.Runner.Volumes,
		Networks:   config.Runner.Networks,
		Devices:    config.Runner.Devices,
		Privileged: config.Runner.Privileged,
		Machine:    config.Runner.Machine,
		Labels:     config.Runner.Labels,
		Environ:    config.Runner.Environ,
		Limits: runner.Limits{
			MemSwapLimit: int64(config.Runner.Limits.MemSwapLimit),
			MemLimit:     int64(config.Runner.Limits.MemLimit),
			ShmSize:      int64(config.Runner.Limits.ShmSize),
			CPUQuota:     config.Runner.Limits.CPUQuota,
			CPUShares:    config.Runner.Limits.CPUShares,
			CPUSet:       config.Runner.Limits.CPUSet,
		},
	}
}
