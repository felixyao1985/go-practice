package main

import (
	"fmt"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/quad"
)

/*
alice follows bob .
bob follows alice .
charlie follows bob .
dani follows charlie .
dani follows alice .
alice is cool .
bob is "not cool" .
charlie is cool .
dani is "not cool" .
*/

func main() {
	// Create a brand new graph
	store, err := cayley.NewMemoryGraph()
	if err != nil {
		fmt.Println(err)
	}

	store.AddQuad(quad.Make("alice", "follows", "bob", nil))
	store.AddQuad(quad.Make("bob", "follows", "alice", nil))
	store.AddQuad(quad.Make("alice", "is", "cool", nil))
	store.AddQuad(quad.Make("felix", "is", "cool", nil))
	store.AddQuad(quad.Make("tina", "is", "cool", nil))
	store.AddQuad(quad.Make("bob", "is", "hot", nil))
	store.AddQuad(quad.Make("tina", "is", "hot", nil))
	store.AddQuad(quad.Make("felix", "is", "gen", nil))

	p := cayley.StartPath(store, quad.String("alice")).Out("follows").Out("is")
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		fmt.Println("alice follows who and he is what:", value.String())
	})

	p2 := cayley.StartPath(store, quad.String("cool")).In("is")
	err = p2.Iterate(nil).EachValue(nil, func(value quad.Value) {
		fmt.Println("who is cool:", quad.NativeOf(value))
	})

	//startTine := cayley.StartPath(store, quad.String("tina"))
	p3 := cayley.StartPath(store, quad.String("felix"), quad.String("tina")).Out("is")
	err = p3.Iterate(nil).EachValue(nil, func(value quad.Value) {
		fmt.Println("felix and tina all is:", quad.NativeOf(value))
	})

	startTine := cayley.StartPath(store, quad.String("tina"))
	p4 := cayley.StartPath(store, quad.String("felix")).Out("is").Or(startTine).Out("is")
	err = p4.Iterate(nil).EachValue(nil, func(value quad.Value) {
		fmt.Println("felix and tina both is:", quad.NativeOf(value))
	})

	if err != nil {
		fmt.Println(err)
	}
}
