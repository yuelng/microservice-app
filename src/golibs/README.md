## 惯用方法
## 微服务经验积累

## 代码结构
https://www.innoq.com/en/blog/the-perils-of-shared-code/

## "一起美"项目关于grpc的一些实践经验

- worc 使用safemap缓存连接,避免频繁的连接
- wonaming 使用etcd做grpc round robin服务发现,负载均衡
- worpc grpc中间件,
- wotracher 在http server grpc 进行跟踪,是否可以将该组件化成其他组件一部分,对业务代码无侵入

https://1024coder.com/topic/43

etcd 使用grpc 进行通信

## practice with code
定义全局错误,例如ErrTokenInvalid = errors.New("JWT Token was invalid"),在代码中使用它
handler + logic + storage

protobuf definitions

protobuf + rabbitmq 统一文件格式??
protobuf + grpc 

self-configuring external service adapters

service.Register{
    &Endpoint{
        Name:"foo",
        Handler
        Request
        Response
    }
}

package handler

func Foo(ctx context.Context, req server.Request)(protobuf.Message, error){

}

deliveries, err := server.Listen()
for delivery := range deliveries {
    go handleDelivery(delivery)
}

var tokens = make(chan struct{}, 20) // 如何为处理链接限速,防止数据库压力过大
                                     // case <- time.After(100 * time.Millisecond)
func handleDelivery(delivery amqp.Delivery){
    // do something for this request
    go c.traceReq(req)
    response, err := executeRequest(delivery)
    go c.traceRsp(req, response, err)
    // send.response
    server.Respond(response)
}

package main
bedrock.Init(config{Name:"",Description:""})
bedrock.RegisterEndpoint(handler.CreateUser)
bedrock.Run()

provisioning
service discovery
configuration
monitoring
Authentication
AB testing
self configuring connectivity to third-party services


gin 框架可以采用嵌入中间件的方式,注入 services client,client连接可以用safemap进行缓存