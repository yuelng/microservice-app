package rpc

import (
	"fmt"
	"log"
	"reflect"

	"base/pool"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// 参考 worc
// StartServiceConns start grpc conns with balancer
// 可以结合服务发现 wonaming,但是这里直接用skydns做服务发现
// "greeter:50000"
func RegisterClient(service string, client interface{}) bool {
	// register factory, auto connect when conn is not set
	factory := func() (interface{}, error) { return grpc.Dial(service, grpc.WithInsecure()) }
	err := pool.RegisterConn(service, factory)
	if err != true {
		log.Printf(`connect to '%s' service failed: %v`, service, err)
	}
	// cache client for call
	return pool.CacheMap.Set(service, client)
}

// 这里可以嵌入 统计代码
// ctx: context
// client: grpc client  pb.NewGreeterClient(conn)
// serviceName: name of service
// metod: method name that you want to use
// req: grpc request
// c := pb.NewGreeterClient(conn)
// r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name,Num:"2"})
func Call(ctx context.Context, serviceName string, method string, req interface{}) (ret interface{}, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("call RPC '%s' error: %v", method, x)
		}
	}()
	conn_, err := pool.GetConn(serviceName)
	if conn_ == nil {
		return nil, fmt.Errorf("service conn '%s' not found", serviceName)
	}

	conn := conn_.(*grpc.ClientConn)
	client := pool.CacheMap.Get(serviceName)
	if client == nil {
		return nil, fmt.Errorf("you should register client for '%s' like pb.NewGreeterClient", serviceName)
	}

	// get NewServiceClient's reflect.Value
	vClient := reflect.ValueOf(client)
	var vParameter []reflect.Value
	vParameter = append(vParameter, reflect.ValueOf(conn))

	// c[0] is serviceServer reflect.Value
	c := vClient.Call(vParameter)

	// rpc param
	v := make([]reflect.Value, 2)
	v[0] = reflect.ValueOf(ctx)
	v[1] = reflect.ValueOf(req)

	// rpc method call
	f := c[0].MethodByName(method)
	resp := f.Call(v)
	if !resp[1].IsNil() {
		return nil, resp[1].Interface().(error)
	}
	return resp[0].Interface(), nil
}

// CloseServiceConns close all established conns
//func CloseServiceConns() {
//	for _, conn := range pool.ServiceConns.List() {
//		conn.(*grpc.ClientConn).Close()
//	}
//}

//func Server(port string)  {
//	lis, err := net.Listen("tcp", port)
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	s := grpc.NewServer()
//
//
//}
