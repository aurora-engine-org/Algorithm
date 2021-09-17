package graph

import (
	"testing"
)

func TestDeepFirstSearch(t *testing.T) {
	g:=&Graph{}
	g.InitPoint("a","b","c","d","e","f","g")
	g.InitConnect(
		Connect{Star: "a",Ends: "b"},
		Connect{Star: "a",Ends: "c"},
		Connect{Star: "b",Ends: "d"},
		Connect{Star: "b",Ends: "e"},
		Connect{Star: "c",Ends: "f"},
		Connect{Star: "c",Ends: "g"},
		//Connect{Star: "e",Ends: "f"},
		)
	g.BuildGraph()
	g.ShowEdge()
	g.DeepFirstSearch(g.Edge,3)
}

func TestBreadthFirstSearch(t *testing.T) {
	g:=&Graph{}
	g.InitPoint("a","b","c","d","e","f","g")
	g.InitConnect(
		Connect{Star: "a",Ends: "b"},
		Connect{Star: "a",Ends: "c"},
		Connect{Star: "b",Ends: "d"},
		Connect{Star: "b",Ends: "e"},
		Connect{Star: "c",Ends: "f"},
		Connect{Star: "c",Ends: "g"},
		//Connect{Star: "e",Ends: "f"},
	)
	g.BuildGraph()
	g.ShowEdge()
	g.BreadthFirstSearch(3)
}

func TestQueue(t *testing.T) {
	q:=&Queue{}
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	e:=q.DeQueue()
	t.Log(e.value)
	e=q.DeQueue()
	t.Log(e.value)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)

}
