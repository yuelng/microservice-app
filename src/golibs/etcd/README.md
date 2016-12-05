## zookeeper,etcd和 consul
## etcd 是基于raft的键值存储系统,可用于服务发现
## 在kubernetes已经存在skydns同样使用etcd作为服务发现

定位服务:

- 服务注册——该步骤存储的信息至少包括正在运行的服务的主机和端口信息
- 服务发现——该步骤允许其他用户可以发现在服务注册阶段存储的信息。

// register
// resolver
// watcher

参考:
- [服务发现：Zookeeper vs etcd vs Consul](http://studygolang.com/articles/4837)
- [github etcd golang client](https://github.com/coreos/etcd/tree/master/client)
- [wothing wonaming](https://github.com/wothing/wonaming)  grpc balancer.