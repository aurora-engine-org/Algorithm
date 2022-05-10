package graph

import (
	"testing"
)

func TestDeepFirstSearch(t *testing.T) {
	g := &Graph{}
	g.InitPoint("a", "b", "c", "d", "e", "f", "g")
	g.InitConnect(
		Line{Star: "a", End: "b"},
		Line{Star: "a", End: "c"},
		Line{Star: "b", End: "d"},
		Line{Star: "b", End: "e"},
		Line{Star: "c", End: "f"},
		Line{Star: "c", End: "g"},
		Line{Star: "a", End: "f"},
	)
	g.BuildGraph()
	g.ShowEdge()
	g.DeepFirstSearch(0)
}

func TestBreadthFirstSearch(t *testing.T) {
	g := &Graph{}
	g.InitPoint("a", "b", "c", "d", "e", "f", "g")
	g.InitConnect(
		Line{Star: "a", End: "b"},
		Line{Star: "a", End: "c"},
		Line{Star: "b", End: "d"},
		Line{Star: "b", End: "e"},
		Line{Star: "c", End: "f"},
		Line{Star: "c", End: "g"},
		Line{Star: "a", End: "f"},
	)
	g.BuildGraph()
	g.ShowEdge()
	g.BreadthFirstSearch(0)
}

func TestQueue(t *testing.T) {
	q := &Queue{}
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	e := q.DeQueue()
	t.Log(e.value)
	e = q.DeQueue()
	t.Log(e.value)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)

}
