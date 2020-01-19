
## 使用

### formatter
$ go run ./lesson03/exercise/formatter/formatter.go
2017/09/24 16:35:56 Initializing logging reporter
2017/09/24 16:36:06 Reporting span 731020308bd6d05d:48394b5372417ee4:3535cabe610946bb:1

### publisher
$ go run ./lesson03/exercise/publisher/publisher.go
2017/09/24 16:35:59 Initializing logging reporter
Hello, Bryan!
2017/09/24 16:36:06 Reporting span 731020308bd6d05d:37908db2de452ea2:4ef2c9b5523bca3b:1


### client
$ go run ./lesson03/exercise/client/hello.go Bryan
2017/09/24 16:36:06 Initializing logging reporter
2017/09/24 16:36:06 Reporting span 731020308bd6d05d:3535cabe610946bb:731020308bd6d05d:1
2017/09/24 16:36:06 Reporting span 731020308bd6d05d:4ef2c9b5523bca3b:731020308bd6d05d:1
2017/09/24 16:36:06 Reporting span 731020308bd6d05d:731020308bd6d05d:0:1



