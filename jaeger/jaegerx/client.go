package jaegerx

import (
	"fmt"
	"net/http"
	"os"

	"google.golang.org/grpc"

	"practice/errx"
)

func DialClientOption(name string) (grpc.DialOption, error) {
	if addr := os.Getenv("JAEGER_ADDR"); addr != "" {
		jaegerTracerServer = addr
	}

	tracer, _, err := NewJaegerTracer(name, jaegerTracerServer)
	if err != nil {
		return nil, err
	}

	if tracer != nil {
		return DialOption(tracer), err
	}
	return nil, errx.Errorx{Code: http.StatusBadRequest, Err: fmt.Errorf("parse parameters failed Validate %v", tracer)}
}
