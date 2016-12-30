## safemap 可用于 grpc 连接池
## 使用读写锁
```go
type safeMap struct {
	lock *sync.RWMutex
	sm   map[string]*grpc.ClientConn
}
```
## 方法有 newMap Get Set Check Delete List
http://openmymind.net/Writing-An-LRU-Cache/
http://openmymind.net/High-Concurrency-LRU-Caching/
http://highscalability.com/blog/2016/1/25/design-of-a-modern-cache.html