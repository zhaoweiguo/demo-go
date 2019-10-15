package main

import (
	"github.com/hashicorp/go-retryablehttp"
	"time"
)

func main() {
	client := retryablehttp.NewClient()
	client.RetryMax = 30
	client.RetryWaitMax = time.Second * 10
	client.RetryWaitMin = time.Second * 1
	client.Logger = nil

}
