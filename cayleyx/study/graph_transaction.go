package main

import (
	g "github.com/cayleygraph/cayley/graph"
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
	store := g.NewTransaction()

	store.AddQuad(quad.Make("alice", "follows", "bob", nil))
	store.AddQuad(quad.Make("bob", "follows", "alice", nil))
	store.AddQuad(quad.Make("alice", "is", "cool", nil))
	store.AddQuad(quad.Make("felix", "is", "cool", nil))
	store.AddQuad(quad.Make("tina", "is", "cool", nil))
	store.AddQuad(quad.Make("bob", "is", "hot", nil))
	store.AddQuad(quad.Make("tina", "is", "hot", nil))
	store.AddQuad(quad.Make("felix", "is", "gen", nil))
	store.AddQuad(quad.Make("felix", "follows", "alex", nil))

	// "alice" -- "follows" -> "bob" add

}
