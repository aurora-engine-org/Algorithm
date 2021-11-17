package main

import (
	"Algorithm/graph"
)

func main() {
	p1 := &graph.Point{Name: "A"}
	p2 := &graph.Point{Name: "B"}
	p3 := &graph.Point{Name: "C"}
	p4 := &graph.Point{Name: "D"}
	m := graph.NewMap(p1, p2, p3, p4)
	m.InsertEdge(p1, p3)
	m.InsertEdge(p1, p2)
	m.InsertEdge(p2, p4)

	m.InsertEdge(p3, p4)
	m.InsertEdge(p1, p2)
	m.InsertEdge(p1, p4)
	m.Dfs(p1)
}
