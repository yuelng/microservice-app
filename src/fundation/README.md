## base code include
- jwt
- pool safemap use map or channel
- grpc protos
- rpc encapsulation
- authentication(使用cookie和session jwt OAuth)


authentication 
session and cookie
set cookie from server
```go
    expiration := time.Now()
    expiration = expiration.AddDate(1, 0, 0)
    cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
    http.SetCookie(w, &cookie)
```
read cookie
```go
    cookie, _ := r.Cookie("username")
    fmt.Fprint(w, cookie)
    // or
    for _, cookie := range r.Cookies() {
        fmt.Fprint(w, cookie.Name)
    }
```


'Authorization':'Bearer ' + token
token header.payload.signature

jwt process
post/users/login with username and password
create a jwt with a secret
return the jwt to the browser
sends the jwt on the Authorization Header
check jwt signature get user information from the jwt
sends response to client

model connection
- Global variables(不利于单元测试)
- Dependency injection(encapsulation db,and pass db from handler to model,injected in main)
- Using an interface
- Request-scoped context
- use closure

use cache
- 使用redis做缓存
- 使用memcache做缓存 caching relatively small and static data
- eviction policies 驱逐策略 lru lfu 
- concurrent 并发情况
- max size limit 最大容量
- persistent store 持久化
- weak references keys 弱引用key
- statistics 分析

use serialization 
- 使用Protobuf做序列化(定义消息问价proto)主要可以将 结构化的数据序列化,可用于数据存储,通信协议
- 编写协议文件,即.proto 文件
- 利用protoc 命令将协议文件转换为我们需要开发语言接口
- 主要接口 Unmarshal Size

在后台系统中总会有一些异步的任务不需要快速执行,所以可以将所有的异步任务放在一个 异步服务中执行,例如 发送邮件,发送短信,推送
一个思路 将各个任务划分为不同的消息队列,在rabbitmq中具体表现是 不同的queue,使用不同的goroutine 将这些work进行异步

use grpc
- 使用gin的中间件功能,将实例化的client注入到context中,然后在handlers中调用
- 或者在main函数中初始化到rpc的包,在handler中直接引入rpc包

email
- use mailgun 每个月10000 封免费邮件

github pages
- 免费托管,轻松发布的web页面

监控
- prometheus

分布式追踪系统

推送服务? gopush-cluster  Apple PushNotification Service APNS

视频处理 ffmpeg
图片处理 graphicsmagick


参考
- [Practical Persistence in Go: Organising Database Access](http://www.alexedwards.net/blog/organising-database-access)
- [gogo protobuf](https://github.com/gogo/protobuf)

