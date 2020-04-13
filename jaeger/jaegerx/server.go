package jaegerx

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"practice/errx"
)

var (
	jaegerTracerServer = ":6831"
)

func NewTracer(name string) (tracer opentracing.Tracer, closer io.Closer, err error) {
	if addr := os.Getenv("JAEGER_ADDR"); addr != "" {
		jaegerTracerServer = addr
	}

	return NewJaegerTracer(name, jaegerTracerServer)
}

func GenServerUnaryInterceptor(name string) (grpc.ServerOption, error) {
	tracer, _, err := NewJaegerTracer(name, jaegerTracerServer)

	if err != nil {
		return nil, err
	}
	if tracer != nil {
		return ServerOption(tracer), err
	}
	return nil, errx.Errorx{Code: http.StatusBadRequest, Err: fmt.Errorf("FU")}
}
