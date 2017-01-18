/*
Copyright 2014 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	pb "base/protos/helloworld"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	//appContext "base/util"
	"base/rpc"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/pborman/uuid"
	"github.com/xyproto/simpleredis"
	//"github.com/uber-go/zap"
	"fmt"
	"golang.org/x/net/context"
)

var (
	masterPool *simpleredis.ConnectionPool
	slavePool  *simpleredis.ConnectionPool
)

//var httpContext = context.Background()

func ListRangeHandler(rw http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["key"]
	list := simpleredis.NewList(slavePool, key)
	members := HandleError(list.GetAll()).([]string)
	membersJSON := HandleError(json.MarshalIndent(members, "", "  ")).([]byte)
	rw.Write(membersJSON)
}

func ListPushHandler(rw http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["key"]
	value := mux.Vars(req)["value"]
	list := simpleredis.NewList(masterPool, key)
	HandleError(nil, list.Add(value))
	ListRangeHandler(rw, req)
}

func InfoHandler(rw http.ResponseWriter, req *http.Request) {
	rqId := uuid.NewRandom()
	fmt.Println(rqId)
	go greeterClient()
	//rqCtx := appContext.WithRqId(httpContext, string(rqId))
	//logger := appContext.Logger(rqCtx)

	info := HandleError(masterPool.Get(0).Do("INFO")).([]byte)
	//logger.Info("handling /doit request", zap.String("for",string(info)))

	rw.Write(info)
}

func EnvHandler(rw http.ResponseWriter, req *http.Request) {
	environment := make(map[string]string)
	for _, item := range os.Environ() {
		splits := strings.Split(item, "=")
		key := splits[0]
		val := strings.Join(splits[1:], "=")
		environment[key] = val
	}

	envJSON := HandleError(json.MarshalIndent(environment, "", "  ")).([]byte)
	rw.Write(envJSON)
}

func HandleError(result interface{}, err error) (r interface{}) {
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	masterPool = simpleredis.NewConnectionPoolHost("redis-master:6379")
	defer masterPool.Close()
	slavePool = simpleredis.NewConnectionPoolHost("redis-slave:6379")
	defer slavePool.Close()

	r := mux.NewRouter()
	r.Path("/lrange/{key}").Methods("GET").HandlerFunc(ListRangeHandler)
	r.Path("/rpush/{key}/{value}").Methods("GET").HandlerFunc(ListPushHandler)
	r.Path("/info").Methods("GET").HandlerFunc(InfoHandler)
	r.Path("/env").Methods("GET").HandlerFunc(EnvHandler)
	go greeterClient()
	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":3000")

}

const (
	//address     = "greeter:50000"
	greeterService = "localhost:50000"
	defaultName    = "world"
)

func greeterClient() {
	ctx := context.Background()
	rpc.RegisterClient(greeterService, pb.NewGreeterClient)

	name := defaultName
	r, err := rpc.Call(ctx, greeterService, "SayHello", &pb.HelloRequest{Name: name, Num: "2"})
	r1, err := rpc.Call(ctx, greeterService, "SayHelloAgain", &pb.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("hello from client" + r.(*pb.HelloReply).Message)
	log.Printf("Greeting: %s", r1.(*pb.HelloReply).Message)

}
