package main

import (
	"fmt"
	"sync"
)
// 创建的时候可以指定一个New函数，
// 获取对象的时候如何在池里面找不到缓存的对象将会使用指定的new函数创建一个返回，
// 如果没有new函数则返回nil
// sync.Pool的定位不是做类似连接池的东西，它的用途仅仅是增加对象重用的几率，减少gc的负担
// relieving pressure on the garbage collector.
func main() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	a := p.Get().(int)
	p.Put(1)

	b := p.Get().(int)
	fmt.Println(a, b)
}

// goroutine pool

//sync.Pool
//bytes.Buffer
//<- chan []byte (LeakyBuffer)