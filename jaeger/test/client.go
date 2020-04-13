package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"practice/jaeger/jaegerx"
	"practice/proto"
)

func main() {
	servOpt, err := jaegerx.DialClientOption("RCli-DoMd5")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	dialOpts = append(dialOpts, servOpt)
	conn, _ := grpc.Dial("127.0.0.1:61001", dialOpts...)
	defer conn.Close()
	clt := test.NewWaiterClient(conn)

	res := "test123"

	tr, err := clt.DoMD5(context.Background(), &test.Req{Str: res})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("服务端响应: %s", tr.BackStr)

	time.Sleep(time.Second * 3)
}
