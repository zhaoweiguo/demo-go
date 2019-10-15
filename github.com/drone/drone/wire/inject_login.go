package main

import (
	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/github"
	"github.com/drone/go-login/login/gitlab"
	"github.com/drone/go-scm/scm/transport/oauth2"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"github.com/zhaoweiguo/demo-go/github.com/drone/drone/wire/config"
)

// wire set for loading the authenticator.
var loginSet = wire.NewSet(
	provideLogin,
	provideRefresher,
)

// provideLogin is a Wire provider function that returns an
// authenticator based on the environment configuration.
func provideLogin(config config.Config) login.Middleware {
	switch {
	case config.Github.ClientID != "":
		return provideGithubLogin(config)
	case config.GitLab.ClientID != "":
		return provideGitlabLogin(config)
	}
	logrus.Fatalln("main: source code management system not configured")
	return nil
}

// provideRefresher is a Wire provider function that returns
// an oauth token refresher for Bitbucket and Gitea
func provideRefresher(config config.Config) *oauth2.Refresher {
	switch {
	case config.Bitbucket.ClientID != "":
		return &oauth2.Refresher{
			ClientID:     config.Bitbucket.ClientID,
			ClientSecret: config.Bitbucket.ClientSecret,
			Endpoint:     "https://bitbucket.org/site/oauth2/access_token",
			Source:       oauth2.ContextTokenSource(),
		}
	case config.Gitea.ClientID != "":
		return &oauth2.Refresher{
			ClientID:     config.Gitea.ClientID,
			ClientSecret: config.Gitea.ClientSecret,
			Endpoint:     config.Gitea.Server + "/login/oauth/access_token",
			Source:       oauth2.ContextTokenSource(),
		}
	}
	return nil
}

// provideGithubLogin is a Wire provider function that returns
// a GitHub authenticator based on the environment configuration.
func provideGithubLogin(config config.Config) login.Middleware {
	if config.Github.ClientID == "" {
		return nil
	}
	return &github.Config{
		ClientID:     config.Github.ClientID,
		ClientSecret: config.Github.ClientSecret,
		Scope:        config.Github.Scope,
		Server:       config.Github.Server,
		Client:       defaultClient(config.Github.SkipVerify),
		Logger:       logrus.StandardLogger(),
	}
}

// provideGitlabLogin is a Wire provider function that returns
// a GitLab authenticator based on the environment configuration.
func provideGitlabLogin(config config.Config) login.Middleware {
	if config.GitLab.ClientID == "" {
		return nil
	}
	return &gitlab.Config{
		ClientID:     config.GitLab.ClientID,
		ClientSecret: config.GitLab.ClientSecret,
		RedirectURL:  config.Server.Addr + "/login",
		Server:       config.GitLab.Server,
		Client:       defaultClient(config.GitLab.SkipVerify),
	}
}
