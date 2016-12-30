## useage VERSION=v4 REGISTRY="docker.io/luebken" make build
## gateway 为web页面\移动app提供基于http协议的restful API接口
## 每个http业务请求都需要鉴定权限,使用jwt(json web token)实现客户端的请求状态信息传递(例如客户端token过期时间)
## 负责追踪 traceid wotraker

## build publish docker  
## VERSION=v1 REGISTRY="192.168.1.10:5000" make release
