package graph

import "fmt"

type Map struct {
	points []*Point
}

type Point struct {
	Name  string  // 标识符
	edges []*Edge // 边表
	mark  bool    //
}

type Edge struct {
	forward *Point //正向点
	reverse *Point //反向点
	weight  int    //权重
	mark    bool   //
}

func NewMap(p ...*Point) *Map {
	return &Map{
		points: p,
	}
}

func (m *Map) InsertEdge(p1, p2 *Point) {
	if p1 == p2 {
		return
	}
	e := &Edge{
		forward: p1,
		reverse: p2,
	}
	for i := 0; i < len(m.points); i++ {
		if m.points[i] == p1 || m.points[i] == p2 {
			if m.points[i].edges == nil {
				m.points[i].edges = make([]*Edge, 0)
				m.points[i].edges = append(m.points[i].edges, e)
				break
			}
			m.points[i].edges = append(m.points[i].edges, e)
		}
	}
}

func (m *Map) Dfs(p *Point) {
	if !p.mark {
		p.mark = true
	} else {
		return
	}
	fmt.Println(p.Name)
	for i := 0; i < len(p.edges); i++ {
		if !p.edges[i].mark {
			//p.edges[i].mark=true
			if p.edges[i].forward == p {
				m.Dfs(p.edges[i].reverse)
			}
		}
	}
}
