package demo1

import (
	"github.com/zhaoweiguo/demo-go/golang.org/mock/demo1/spider"
)

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
