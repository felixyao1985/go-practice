package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"practice/jaeger/jaegerx"
	"practice/proto"
)

type server struct{}

func DialD(ctx context.Context) {
	servOpt, _ := jaegerx.DialClientOption("RCli-GetOrg")
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	dialOpts = append(dialOpts, servOpt)
	conn, _ := grpc.Dial("127.0.0.1:61002", dialOpts...)
	//conn, _ := grpcx.Dial("127.0.0.1:61002", dialOpts...)
	defer conn.Close()
	clt := test.NewOrgServiceClient(conn)

	fmt.Println(clt.Get(ctx, &test.Id{Id: "103552335095005470"}))
	time.Sleep(time.Second * 3)
}

func (s *server) DoMD5(ctx context.Context, in *test.Req) (*test.Res, error) {
	fmt.Println("MD5方法请求JSON:" + in.Str)
	DialD(ctx)
	return &test.Res{BackStr: "MD5 :" + fmt.Sprintf("%x", md5.Sum([]byte(in.Str)))}, nil
}

func main() {
	servOpt, err := jaegerx.GenServerUnaryInterceptor("RSev1")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer([]grpc.ServerOption{
		servOpt,
	}...)

	test.RegisterWaiterServer(s, &server{})
	reflection.Register(s)

	ln, err := net.Listen("tcp", "127.0.0.1:61001")
	if err != nil {
		panic(err)
	}
	err = s.Serve(ln)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
