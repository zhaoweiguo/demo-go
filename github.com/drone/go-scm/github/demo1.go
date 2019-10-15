package main

import (
	"crypto/tls"
	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/github"
	"github.com/drone/go-scm/scm/transport/oauth2"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	client := getClient()
	log.Println(client.Driver.String()) // github
}

func getClient() *scm.Client {
	client, err := github.New("https://github.com")
	if err != nil {
		log.Fatalln("err:", err)
	}
	client.DumpResponse = httputil.DumpResponse // optional
	client.Client = &http.Client{
		Transport: &oauth2.Transport{
			Source: oauth2.ContextTokenSource(),
			Base:   defaultTransport(false),
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
