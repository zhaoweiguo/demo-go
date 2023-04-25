package main

import (
	"net/url"
	"testing"

	"github.com/bmizerany/assert"
)

func TestBasicAuth(t *testing.T) {
	u, _ := url.Parse("http://admin:admin123@zhaoweiguo.com/hello/world")
	assert.Equal(t, u.User.Username(), "admin")
	pwd, _ := u.User.Password()
	assert.Equal(t, pwd, "admin123")
	assert.Equal(t, u.User.String(), "admin:admin123")
}
