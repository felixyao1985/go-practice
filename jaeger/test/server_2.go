package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"practice/jaeger/jaegerx"
	"practice/proto"
)

type OrgService struct {
	*test.UnimplementedOrgServiceServer
}

func (s *OrgService) Get(ctx context.Context, in *test.Id) (*test.Org, error) {
	fmt.Println("Get Org:")
	return &test.Org{}, nil
}

func main() {
	servOpt, err := jaegerx.GenServerUnaryInterceptor("RSev2")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer([]grpc.ServerOption{
		servOpt,
	}...)

	test.RegisterOrgServiceServer(s, &OrgService{})
	reflection.Register(s)

	ln, err := net.Listen("tcp", "127.0.0.1:61002")
	if err != nil {
		panic(err)
	}
	err = s.Serve(ln)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
