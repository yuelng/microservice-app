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