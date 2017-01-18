package reflect

import (
	"context"
	"fmt"
	"golibs/safemap"
	"google.golang.org/grpc"
	"log"
	"reflect"
)

var serviceConns = cache.NewSafeMap()

// StartServiceConns start grpc conns with balancer
// address is etcd address, servicelist is servicename list
// grpc roundRobin grpc etcd 服务发现
func StartServiceConns(address string, serviceList []string) {
	for _, serviceName := range serviceList {
		go func(name string) {
			// 这里应该是根据服务名称获取服务地址??
			// r := wonaming.NewResolver(name)
			// b := grpc.RoundRobin(r)
			var b int = nil
			conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBalancer(b))
			if err != nil {
				log.Printf(`connect to '%s' service failed: %v`, name, err)
			}
			serviceConns.Set(name, conn)
		}(serviceName)
	}
}

// CloseServiceConns close all established conns
func CloseServiceConns() {
	for _, conn := range serviceConns.List() {
		conn.Close()
	}
}

// params ctx func serviceName method req
// reflect.CallRPC(ctx, pb.NewGreeterClient, hello, hello, &pb.HelloRequest{Name: name,Num:"2"})
func CallRPC(ctx context.Context, client interface{}, serviceName string, method string, req interface{}) (ret interface{}, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("call RPC '%s' error: %v", method, x)
		}
	}()

	conn := serviceConns.Get(serviceName)
	if conn == nil {
		return nil, fmt.Errorf("service conn '%s' not found", serviceName)
	}

	// get NewServiceClient's reflect.Value
	// client 即pb.NewGreeterClient
	// 使用reflection实现,NewGreeterClient函数调用 c := pb.NewGreeterClient(conn)
	vClient := reflect.ValueOf(client)
	var vParameter []reflect.Value
	vParameter = append(vParameter, reflect.ValueOf(conn))

	// c[0] is serviceServer reflect.Value
	c := vClient.Call(vParameter) // ==> c := pb.NewGreeterClient(conn) 使用反射的函数调用返回值为reflect.Value类型

	// rpc param
	v := make([]reflect.Value, 2)
	v[0] = reflect.ValueOf(ctx)
	v[1] = reflect.ValueOf(req)

	// rpc method call
	// 使用reflection从实体中获取函数,methodByName, 反射函数调用返回值也是reflect.Value数组
	f := c[0].MethodByName(method)
	resp := f.Call(v)
	if !resp[1].IsNil() {
		return nil, resp[1].Interface().(error)
	}
	return resp[0].Interface(), nil
}
