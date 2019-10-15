package main

import (
	"crypto/tls"
	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/github"
	"github.com/drone/go-scm/scm/driver/gitlab"
	"github.com/drone/go-scm/scm/transport/oauth2"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"github.com/zhaoweiguo/demo-go/github.com/drone/drone/wire/config"
	"net/http"
	"net/http/httputil"
)

// wire set for loading the scm client.
var clientSet = wire.NewSet(
	provideClient,
)

// provideBitbucketClient is a Wire provider function that
// returns a Source Control Management client based on the
// environment configuration.
func provideClient(config config.Config) *scm.Client {
	switch {
	case config.Github.ClientID != "":
		return provideGithubClient(config)
	case config.GitLab.ClientID != "":
		return provideGitlabClient(config)
	}
	logrus.Fatalln("main: source code management system not configured")
	return nil
}

// provideGithubClient is a Wire provider function that returns
// a GitHub client based on the environment configuration.
func provideGithubClient(config config.Config) *scm.Client {
	client, err := github.New(config.Github.APIServer)
	if err != nil {
		logrus.WithError(err).
			Fatalln("main: cannot create the GitHub client")
	}
	if config.Github.Debug {
		client.DumpResponse = httputil.DumpResponse
	}
	client.Client = &http.Client{
		Transport: &oauth2.Transport{
			Source: oauth2.ContextTokenSource(),
			Base:   defaultTransport(config.Github.SkipVerify),
		},
	}
	return client
}

// provideGitlabClient is a Wire provider function that returns
// a GitLab client based on the environment configuration.
func provideGitlabClient(config config.Config) *scm.Client {
	logrus.WithField("server", config.GitLab.Server).
		WithField("client", config.GitLab.ClientID).
		WithField("skip_verify", config.GitLab.SkipVerify).
		Debugln("main: creating the GitLab client")

	client, err := gitlab.New(config.GitLab.Server)
	if err != nil {
		logrus.WithError(err).
			Fatalln("main: cannot create the GitLab client")
	}
	if config.GitLab.Debug {
		client.DumpResponse = httputil.DumpResponse
	}
	client.Client = &http.Client{
		Transport: &oauth2.Transport{
			Source: oauth2.ContextTokenSource(),
			Base:   defaultTransport(config.GitLab.SkipVerify),
		},
	}
	return client
}

// defaultTransport provides a default http.Transport. If
// skipverify is true, the transport will skip ssl verification.
func defaultTransport(skipverify bool) http.RoundTripper {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipverify,
		},
	}
}

// defaultClient provides a default http.Client. If skipverify
// is true, the default transport will skip ssl verification.
func defaultClient(skipverify bool) *http.Client {
	client := &http.Client{}
	client.Transport = defaultTransport(skipverify)
	return client
}
