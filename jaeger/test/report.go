package main

import (
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"

	"practice/jaeger/jaegerx"
)

func main() {
	jt, _, _ := jaegerx.NewJaegerTracer("RCli-DoMd5", "139.196.120.170:6831")
	t, err := jaeger.NewUDPTransport("139.196.120.170:6831", 1024)
	if err != nil {
		panic(err)
	}

}
