package handlers

import (
	"base/protos/message"
	"fmt"
	me "github.com/bradfitz/gomemcache/memcache"
)

func main() {
	fmt.Println("hello")

	mc := me.New("192.168.1.10:11211")
	mc.Set(&me.Item{Key: "foo", Value: []byte("my value")})

	it, err := mc.Get("foo")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(it.Value))

	testSerialzation()
}

func testSerialzation() {
	mc := me.New("192.168.1.10:11211")

	ms := message.HelloRequest{
		Name: "hello",
		Num:  "123",
	}
	byteValue, err := ms.Marshal()
	if err != nil {
		fmt.Println(err)
	}
	mc.Set(&me.Item{Key: "foo", Value: byteValue})

	it, err := mc.Get("foo")
	if err != nil {
		fmt.Println(err)
	}

	newMessage := message.HelloRequest{}
	err = newMessage.Unmarshal(it.Value)
	dealError(err)
	fmt.Println(newMessage.Name)
	fmt.Println(newMessage.Num)
}

func dealError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
