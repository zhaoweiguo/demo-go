package main

import (
	"net/url"
	"testing"

	"github.com/bmizerany/assert"
)

const (
	username   = "admin"
	password   = "admin123"
	userAndPwd = "admin:admin123"
	locAuth    = "http://admin:admin123@zhaoweiguo.com/hello/world"
	locNoAuth  = "http://zhaoweiguo.com/hello/world"
)

func TestBasicAuth(t *testing.T) {
	u, _ := url.Parse(locAuth)
	assert.Equal(t, u.User.Username(), username)
	pwd, _ := u.User.Password()
	assert.Equal(t, pwd, password)
	assert.Equal(t, u.User.String(), userAndPwd)
	assert.Equal(t, u.String(), locAuth)
}

func TestUser(t *testing.T) {
	ui := url.UserPassword(username, password)
	assert.Equal(t, ui.Username(), username)
	assert.Equal(t, ui.String(), userAndPwd)
	pwd, exist := ui.Password()
	assert.Equal(t, pwd, password)
	assert.Equal(t, exist, true)

	u, _ := url.Parse(locNoAuth)
	u.User = ui
	assert.Equal(t, u.String(), locAuth)
}
