package pool

import (
	"errors"
	"log"
	"time"
)

var CacheMap = newSafeMap()
var ServiceConns = newSafeMap()
var registerConns = newSafeMap()

// Factory is a function to create new connections.
// 可以使用map 同时也可以使用 channel做  p.store = make(chan interface{}, maxCap)
// 连接池并不需要反复连接,缓存连接即可
// http://www.ryanday.net/2012/09/12/golang-using-channels-for-a-connection-pool/
// https://www.vividcortex.com/blog/2015/01/19/gos-connection-pool-retries-and-timeouts/
type Factory func() (interface{}, error)

func RegisterConn(service string, factory Factory) bool {
	return registerConns.Set(service, factory)
}

func GetConn(service string) (interface{}, error) {
	var conn interface{}
	if conn = ServiceConns.Get(service); conn == nil {
		return SetConn(service)
	}
	return conn, nil
}

func SetConn(service string) (interface{}, error) {
	var conn interface{}
	factory := registerConns.Get(service)
	if factory == nil {
		return nil, errors.New("Must Register Conn!")
	}
	conn = redial(factory.(Factory))
	ServiceConns.Set(service, conn)
	return conn, nil
}

func redial(factory Factory) interface{} {
	for {
		conn, err := factory()
		if err == nil {
			return conn
		}

		log.Println(err)
		log.Println("Trying to reconnect to service failer")
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

// close service conn
