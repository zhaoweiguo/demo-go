## 简介
```
groupcache is a distributed caching and cache-filling library, 
intended as a replacement for a pool of memcached nodes in many cases.
```

## Comparison to memcached

#### Like memcached, groupcache:

```
shards by key to select which peer is responsible for that key
```

#### Unlike memcached, groupcache:

```
1. does not require running a separate set of servers, 
thus massively reducing deployment/configuration pain. 
groupcache is a client library as well as a server. 
It connects to its own peers, forming a distributed cache.

2. comes with a cache filling mechanism. 
Whereas memcached just says "Sorry, cache miss", 
often resulting in a thundering herd of database (or whatever) 
loads from an unbounded number of clients (which has resulted in several fun outages), 
groupcache coordinates cache fills such that only one load in one process of 
an entire replicated set of processes populates the cache, 
then multiplexes the loaded value to all callers.

3. does not support versioned values. 
If key "foo" is value "bar", key "foo" must always be "bar". 
There are neither cache expiration times, nor explicit cache evictions. 
Thus there is also no CAS, nor Increment/Decrement. 
This also means that groupcache....

4. ... supports automatic mirroring of super-hot items to multiple processes. 
This prevents memcached hot spotting where 
a machine's CPU and/or NIC are overloaded by very popular keys/values.

5. is currently only available for Go. 
It's very unlikely that I (bradfitz@) will port the code to any other language.
```

## singleFlight
```
里面有一个很有意思的库，叫singleFlight, 
因为groupCache从节点上获取缓存如果未命中，则会去其他节点寻找，
其他节点还没有的话再从数据源获取，所以这个步骤对于防击穿非常有必要。
singleFlight使得groupCache在多个并发请求对一个失效的key进行源数据获取时，
只让其中一个得到执行，其余阻塞等待到执行的那个请求完成后，
将结果传递给阻塞的其他请求达到防止击穿的效果。
```

