package main

import (
	"github.com/drone/drone/core"
	"github.com/drone/drone/service/license"
	"github.com/drone/go-scm/scm"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"github.com/zhaoweiguo/demo-go/github.com/drone/drone/wire/config"
)

// wire set for loading the license.
var licenseSet = wire.NewSet(
	provideLicense,
	license.NewService,
)

// provideLicense is a Wire provider function that returns a
// license loaded from a license file.
func provideLicense(client *scm.Client, config config.Config) *core.License {
	l, err := license.Load(config.License)
	if config.License == "" {
		l = license.Trial(client.Driver.String())
	} else if err != nil {
		logrus.WithError(err).
			Fatalln("main: invalid or expired license")
	}
	logrus.WithFields(
		logrus.Fields{
			"kind":        l.Kind,
			"expires":     l.Expires,
			"repo.limit":  l.Repos,
			"user.limit":  l.Users,
			"build.limit": l.Builds,
		},
	).Debugln("main: license loaded")
	return l
}
