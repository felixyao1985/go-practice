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
	store.AddQuad(quad.Make("felix", "follows", "alex", nil))
	store.AddQuad(quad.Make("felix", "follows", "bob", nil))

	store.AddQuad(quad.Make("bob", "have", "cat", nil))

	p := cayley.StartPath(store).Out("follows").Is(quad.String("bob"))

	err = p.Reverse().Iterate(nil).EachValue(nil, func(value quad.Value) {
		fmt.Println("需要通过Reverse()来输出谁关注了bob", value.String())
	})

	err = cayley.StartPath(store, quad.String("bob")).OutPredicates().Iterate(nil).EachValue(store, func(value quad.Value) {
		fmt.Println("所有bob的入口有那些:", value.String())
	})

	p2 := cayley.StartPath(store).Out("follows").Out("is")

	err = p2.Iterate(nil).EachValue(nil, func(value quad.Value) {
		fmt.Println("有人被folows 并且 is 为：", value.Native())
	})

	if err != nil {
		fmt.Println(err)
	}
}
