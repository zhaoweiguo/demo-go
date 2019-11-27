package main

import (
	"github.com/golang/mock/gomock"
	"github.com/zhaoweiguo/demo-go/golang.org/mock/demo1/spider"
	"testing"
)

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}

func TestGetGoVersion(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockSpider := spider.NewMockSpider(mockCtl)
	mockSpider.EXPECT().GetBody().Return("go1.8.3")
	goVer := GetGoVersion(mockSpider)

	if goVer != "go1.8.3" {
		t.Error("Get wrong version %s", goVer)
	}

}
