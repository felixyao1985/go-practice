package wrapper

import (
	"crypto/md5"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"practice/jaeger/jaegerx"
	pb "practice/proto"
)

type Server struct{}

func (s *Server) DoMD5(ctx context.Context, in *pb.Req) (*pb.Res, error) {
	fmt.Println("MD5方法请求JSON:" + in.Str)
	return &pb.Res{BackStr: "MD5 :" + fmt.Sprintf("%x", md5.Sum([]byte(in.Str)))}, nil
}

func Test_Tracing(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		os.Exit(-1)
	}

	var servOpts []grpc.ServerOption

	servOpt, err := jaegerx.GenServerUnaryInterceptor("nSrv")
	servOpts = append(servOpts, servOpt)
	svr := grpc.NewServer(servOpts...)
	pb.RegisterWaiterServer(svr, &Server{})

	go func() {
		time.Sleep(time.Second)

		dialOpts := []grpc.DialOption{grpc.WithInsecure()}

		servOpt, err := jaegerx.DialClientOption("felixCli")
		dialOpts = append(dialOpts, servOpt)
		conn, err := grpc.Dial("127.0.0.1:8001", dialOpts...)
		if err != nil {
			fmt.Printf("grpc connect failed, err:%+v\n", err)
			os.Exit(-1)
		}
		defer conn.Close()

		client := pb.NewWaiterClient(conn)
		resp, err := client.DoMD5(context.Background(), &pb.Req{Str: "felix-test"})
		if err != nil {
			fmt.Printf("call sayhello failed, err:%+v\n", err)
			os.Exit(-1)
		} else {
			fmt.Printf("call sayhello suc, res:%+v\n", resp)
		}
	}()

	go svr.Serve(ln)

	time.Sleep(time.Second * 3)
}
