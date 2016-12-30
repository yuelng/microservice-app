package main

import (
	"fmt"
	"sync"
	""
	"github.com/streadway/amqp"
	"golang.org/x/net/context"
	"log"
	"time"
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

// 参考streadway amqp pubsub
// ctx, done := context.WithCancel(context.Background())
// <-ctx.Done()
func dial(ctx context.Context)  {
	sessions := make(chan chan interface{})
	//ctx, done := context.WithCancel(context.Background())
	go func() {
		sess := make(chan interface{})
		for {
			select {
			case sessions <- sess:
			case <- ctx.Done():
				log.Println("finished")
				return
			}

			conn, err := amqp.Dial("")
			if err != nil {
				log.Println("dial error")
			}
			select {
			case sess <- conn:
			case <- ctx.Done():
				log.Println("finished")
				return
			}
		}
	}()
	return sessions
}


func redial(uri string) interface{} {
	for {
		conn, err := amqp.Dial(uri)

		if err == nil {
			return conn
		}

		log.Println(err)
		log.Printf("Trying to reconnect to RabbitMQ at %s\n", uri)
		time.Sleep(500 * time.Millisecond)
	}
}

var Mqonce sync.Once
var conn interface{}
func reredial(uri string) interface{} {
	Mqonce.Do(func(){
		conn, err := amqp.Dial("")

		if err == nil {
			return conn
		}
	})
	return conn
}

var safeMap safeMap
func rereredial(uri string) interface{} {
	if conn := safeMap.Get("rabbitmq");conn == nil {
		conn, err := amqp.Dial("")

		if err == nil {
			return conn
		}
		safeMap.Set("rabbitmq", conn)
	}
	return conn
}
// 使用channel构建pool池子 make(chan conn, maxCap)
// connection pool from golang sql
// Connections are created when needed and there isn’t a free connection in the pool
// retries timeout,
// https://github.com/fatih/pool