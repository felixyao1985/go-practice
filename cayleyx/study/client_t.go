package main

import (
	"fmt"
	v2 "practice/cayleyx/v2"
)

func main() {
	clt := v2.NewAPIClient("http://localhost:64210")

	d, err := clt.Gremlin("graph.Vertex().all()")
	fmt.Println(string(d), err)
	t := v2.Triads{}

	t = append(t, v2.Quad{
		Subject:   "felix",
		Predicate: "time",
		Object:    "2000-12-02",
		Label:     "",
	}, v2.Quad{
		Subject:   "felix",
		Predicate: "time",
		Object:    "2000-12-01",
		Label:     "",
	})
	fmt.Println("test~~~~~~~")
	for _, item := range t {
		i := fmt.Sprintln(item.Subject, item.Predicate, item.Object, item.Label)
		fmt.Println(i)
	}
	//
	err = clt.Write(t)
	fmt.Println(err)
	//err = clt.Delete(t)
	//fmt.Println(err)
}
