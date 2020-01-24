
## 前提

启动docker服务jaegertracing/all-in-one
```
$ docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.14
```

## 参考实例

* https://github.com/yurishkuro/opentracing-tutorial
* https://juejin.im/post/5d7ed711e51d4562165535ab

## golang语言版

https://github.com/yurishkuro/opentracing-tutorial/tree/master/go

## 其他实例

* https://github.com/PacktPublishing/Mastering-Distributed-Tracing
* https://www.shkuro.com/books/2019-mastering-distributed-tracing/

