# 添加一个服务
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

# use jwt stateless 
状态指的是请求在client/server交互过程中的上下文信息
有状态是指,server保存了client的请求状态,server会通过client传递的sessionId在其作用域内找到之前的交互信息
无状态是指,每个请求的是独立的,要求客户端保存所有需要认证的消息,每次认证带上自己的状态
github.com/dgrijalva/jwt-go

# 为http加上ssl 


参考
[jwt-go](https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac)
[acme.sh](https://github.com/Neilpang/acme.sh)