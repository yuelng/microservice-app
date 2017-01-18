### 添加一个服务
文件结构基本如下
- 最外层的Dockerfile主要制作golang镜像,将
- makefile 执行docker操作,并将最终打包好的BusyBox上传到私人docker仓库 192.168.1.10
- 根目录下的Dockerfile在golang 1.7的container中打包二进制文件
- greeter 中的dockerfile 是将生成的 服务二进制文件打包到 busybox:Ubuntu 14.04


├── Dockerfile
├── Makefile
├── README.md
├── greeter
│   └── Dockerfile
├── handlers
│   └── hello.go
└── main.go


docker承载各个微服务,APIGateway,其他服务,docker编排工具使用 kubernetes,配置文件为根目录下config文件夹
通过更改image 版本达到多版本持续部署的效果.
build.sh 脚本的主要作用是 在golang1.7镜像的基础上,将源代码拷入镜像,在容器中编译文件,得到二进制文件
得到的二进制文件,再打包到busybox镜像中,此镜像就可以作为部署的image,将该镜像推送到私人image库

使用supervisor使得docker,kubelet,kube-proxy,fluentd自动化
触发服务器部署项目,使用kuberctl从私有仓库中拉取最新的image,实现部署
整个流程可分为 编写代码 编译 打包镜像 部署
编译 打包镜像 部署 都由jenkins代劳,当然也可以手动将代码拷到目标服务器,自己在编译,打包,部署

### use jwt stateless 
- 状态指的是请求在client/server交互过程中的上下文信息
- 有状态是指,server保存了client的请求状态,server会通过client传递的sessionId在其作用域内找到之前的交互信息
- 无状态是指,每个请求的是独立的,要求客户端保存所有需要认证的消息,每次认证带上自己的状态
- github.com/dgrijalva/jwt-go


### base
- 提供grpc client call 封装
- authorization
- connection pool safemap

### guestbook 
- restful api (gateway)
- authorization
- grpc client
- rabbitmq publisher

### greeter
- grpc server
- rabbitmq consumer

### api
- gin web
- restful API
- gorm postgresql

### rabbitmq
- 持久化,在第一次声明队列和发送消息时指定其持久化属性为true
- 接收应答,设置autoAck为false可以让客户端主动应答,当客户端拒绝此消息或者未应答便断开连接,就会使该消息重新入队
- 发送确认,设置channel为confirm模式,所有发送的消息都会被确认一次

restful gateway -- grpc call other service 
                -- publish message by rabbitmq

### TODO
- 监控组件: prometheus + grafana
- 跟踪组件: zipkin + elasticsearch
- https
postgis 地理坐标位置应用

- single responsibility principle
- bounded context
- well defined interfaces
- composability

sentry 捕捉错误信息
elasticsearch

self deploy,部署都在源代码中设置好了,

参考
- [jwt-go](https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac)
- [acme.sh](https://github.com/Neilpang/acme.sh)
