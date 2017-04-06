package services

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	//"gopkg.in/redis.v5"
)

const redisService = "localhost:6379"
const maxConnections = 30

//func InitRedis() *redis.Client {
//	var c *redis.Client
//	once.Do(func() {
//		c = redis.NewClient(&redis.Options{
//			Addr:     redisService,
//			Password: "", // no password set
//			DB:       0,  // use default DB
//		})
//	})
//	pong, err := c.Ping().Result()
//	fmt.Println(pong, err)
//	return c
//}

//func redigoClient() {
//	c, err := redigo.Dial("tcp", "localhost:6379")
//	if err != nil {
//		fmt.Println("Connect to redis error", err)
//		return
//	}
//	defer c.Close()
//}

func pool() {
	redisPool := redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", redisService)

		if err != nil {
			return nil, err
		}

		return c, err
	}, maxConnections)

	defer redisPool.Close()
	c := redisPool.Get()
	//status, err := c.Do("SET", key, value)
	defer c.Close()
}
func main() {
	//Connect
	c, err := redis.Dial("tcp", "6d0350b##########523956f768b.publb.rackspaceclouddb.com:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	//Authenticate
	c.Do("AUTH", "KtKzkTB##############4jExUADzRfKC")

	//Set two keys
	c.Do("SET", "best_car_ever", "Tesla Model S")
	c.Do("SET", "worst_car_ever", "Geo Metro")

	//Get a key
	best_car_ever, err := redis.String(c.Do("GET", "best_car_ever"))
	if err != nil {
		fmt.Println("best_car_ever not found")
	} else {
		//Print our key if it exists
		fmt.Println("best_car_ever exists: " + best_car_ever)
	}

	//Delete a key
	c.Do("DEL", "worst_car_ever")

	//Try to retrieve the key we just deleted
	worst_car_ever, err := redis.String(c.Do("GET", "worst_car_ever"))
	if err != nil {
		fmt.Println("worst_car_ever not found", err)
	} else {
		//Print our key if it exists
		fmt.Println(worst_car_ever)
	}
}
