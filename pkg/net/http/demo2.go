package main

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
)

// 前面的几个方法都是在http.DefaultClient的基础上进行的调用
type Client struct {
	Transport     RoundTripper                                       // Transport用于确定http的创建机制,默认为DefaultTransport
	CheckRedirect func(req *http.Request, via []*http.Request) error // 定义重定向策略
	Jar           http.CookieJar                                     // 若Jar为空,cookie不会在请求中发送,并在响应中忽略
}

// 自定义Transport
type Transport struct {
	Proxy               func(*http.Request) (*url.URL, error)          // 指定代理函数, 默认不使用代理
	Dial                func(net, addr string) (c net.Conn, err error) // 指定创建tcp连接的Dial()函数,默认为net.Dial()
	TLSClientConfig     *tls.Config                                    // 指定tls.Client的TLS配置(SSL连接专用)
	DisableKeepAlives   bool
	DisableCompression  bool
	MaxIdleConnsPerHost int // 控制每个host所需求保持的最大空闲连接数
}

type RoundTripper interface {
	RoundTrip(*http.Request) (*http.Response, error)
}
