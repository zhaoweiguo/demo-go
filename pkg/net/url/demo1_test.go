package main

import (
	"fmt"
	"log"
	"net/url"
	"testing"

	"github.com/bmizerany/assert"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func TestA(t *testing.T) {
	u, err := url.Parse("http://www.baidu.com/search?q=dotnet")
	if err != nil {
		panic(err)
	}

	fmt.Printf("url.string(): %s\n", u.String())
}

func TestB(t *testing.T) {
	log.Println(url.ParseQuery("中国"))
	log.Println(url.PathEscape("aaa中国 北京bbb"))
}

func TestHttpReal(t *testing.T) {
	uri, err := url.Parse("http://zhaoweiguo.com/Downloads/tmp/ME2002_000044A.bin")
	assert.Equal(t, err, nil)
	assert.Equal(t, uri.Path, "/Downloads/tmp/ME2002_000044A.bin")
	assert.Equal(t, uri.Scheme, "http")
	assert.Equal(t, uri.Host, "zhaoweiguo.com")
	assert.Equal(t, uri.RawPath, "")
	assert.Equal(t, uri.RawQuery, "")
}

func TestHttpFake(t *testing.T) {
	uri, err := url.Parse("http://Downloads/tmp/ME2002_000044A.bin")
	assert.Equal(t, err, nil)
	assert.Equal(t, uri.Path, "/tmp/ME2002_000044A.bin")
	assert.Equal(t, uri.Scheme, "http")
	assert.Equal(t, uri.Host, "Downloads")
	assert.Equal(t, uri.RawPath, "")
	assert.Equal(t, uri.RawQuery, "")
}

func TestFilePath(t *testing.T) {
	uri, err := url.Parse("~/Downloads/tmp/ME2002_000044A.bin")
	assert.Equal(t, err, nil)
	assert.Equal(t, uri.Path, "~/Downloads/tmp/ME2002_000044A.bin")
	assert.Equal(t, uri.Scheme, "")
	assert.Equal(t, uri.Host, "")
	assert.Equal(t, uri.RawPath, "")
	assert.Equal(t, uri.RawQuery, "")
}
