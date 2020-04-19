package goodhosts

import (
	"github.com/lextoumbourou/goodhosts"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func TestListEntries(t *testing.T) {
	hosts, err := goodhosts.NewHosts() // 注: 首先要获得hosts对象
	if err != nil {
		panic(err.Error())
	}

	for _, line := range hosts.Lines {
		log.Println(line.Raw)
	}

}

func TestCheckEntry(t *testing.T) {
	hosts, err := goodhosts.NewHosts()
	if err != nil {
		panic(err.Error())
	}

	cases := []struct {
		ip   string
		host string
	}{
		{"127.0.0.1", "facebook.com"},
		{"192.168.99.205", "default.backend.com"},
		{"10.128.132.38", "consul.siot.go"},
	}

	for _, c := range cases {
		if hosts.Has(c.ip, c.host) {
			log.Printf("Exist: %s, %s\n", c.ip, c.host)
		} else {
			log.Printf("Not Exist: %s, %s\n", c.ip, c.host)
		}
	}

}

// 注: 需要sudo权限
func TestAddEntry(t *testing.T) {
	hosts, err := goodhosts.NewHosts()
	if err != nil {
		panic(err.Error())
	}

	cases := []struct {
		ip   string
		host string
	}{
		{"127.0.0.1", "goodhosts.qq.com"},
		{"192.168.99.205", "goodhosts.baidu.com"},
		{"10.128.132.38", "goodhosts.google.com"},
	}

	for _, c := range cases {
		if err = hosts.Add(c.ip, c.host); err != nil {
			log.Printf("Success: %s, %s\n", c.ip, c.host)
		} else {
			log.Printf("Not Success: %s, %s\n", c.ip, c.host)
		}
	}
}

// 注: 需要sudo权限
func TestRemoveEntry(t *testing.T) {
	hosts, err := goodhosts.NewHosts()
	if err != nil {
		panic(err.Error())
	}

	cases := []struct {
		ip   string
		host string
	}{
		{"127.0.0.1", "goodhosts.qq.com"},
		{"192.168.99.205", "goodhosts.baidu.com"},
		{"10.128.132.38", "goodhosts.google.com"},
	}

	for _, c := range cases {
		if err = hosts.Remove(c.ip, c.host); err != nil {
			log.Printf("Success: %s, %s\n", c.ip, c.host)
		} else {
			log.Printf("Not Success: %s, %s\n", c.ip, c.host)
		}
	}
	if err := hosts.Flush(); err != nil {
		panic(err)
	}
}
