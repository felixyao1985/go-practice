package main

import (
	"fmt"
	"net"

	"github.com/cayleygraph/cayley/client"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/graph/memstore"
	cayleyhttp "github.com/cayleygraph/cayley/server/http"
	"github.com/cayleygraph/cayley/writer"
	"github.com/cayleygraph/quad"
)

func Handle(quads ...quad.Quad) *graph.Handle {
	qs := memstore.New(quads...)
	wr, _ := writer.NewSingleReplication(qs, nil)

	return &graph.Handle{qs, wr}
}

func main() {
	h := Handle()
	api := cayleyhttp.NewAPIv2(h)
	fmt.Println(api)

	l, err := net.Listen("tcp", "localhost:64210")
	fmt.Println(err)
	fmt.Println(l.Addr().String())
	cli := client.New(l.Addr().String())

	qr, err := cli.QuadReader()
	fmt.Println(qr)

}
