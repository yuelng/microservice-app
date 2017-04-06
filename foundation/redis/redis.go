package redis

import (
	"base/pool"
	"gopkg.in/redis.v5"
	"log"
)

const serverName = "redis"

func RegisterClient(service string) bool {
	// register factory, auto connect when conn is not set
	factory := func() (interface{}, error) {
		conn := redis.NewClient(&redis.Options{
			Addr:     service,
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		return conn, nil
	}
	err := pool.RegisterConn(serverName, factory)
	if err != true {
		log.Printf(`connect to '%s' service failed: %v`, serverName, err)
	}
	// cache client for call
	return true
}

func GetConn() *redis.Client {
	conn, err := pool.GetConn(serverName)
	if err != nil {
		log.Printf(`connect to '%s' service failed: %v`, serverName, err)
	}
	return conn.(*redis.Client)
}
