package graph

import "fmt"
/*
	此图算法基于 无向图，遍历顺序取决于 Point第一个点的位置
*/
type Graph struct {
	Point  []string				//图顶点信息
	Edge   [][]int				//边关系表
	connect   []Connect			//连接信息
	Mark   []bool				//访问标识
}

// Connect 连接信息
type Connect struct {
	 Star string
	 Ends string
}

// InitPoint 初始化点信息
func (g *Graph) InitPoint(points ...string)  {
		g.Point=points
		l:= len(points)
		g.Edge=make([][]int,l)
		for i,_:=range g.Edge{
			g.Edge[i]=make([]int,l)
		}
		g.Mark=make([]bool,l)
}

// InitConnect 初始化连接信息
func (g *Graph) InitConnect(connects... Connect) {
	g.connect=connects
}

// BuildGraph 构建图
func (g *Graph) BuildGraph()  {
	for i:=0;i< len(g.connect);i++{			//外层循环读取连接信息
		for j:=0;j< len(g.Point);j++{		//找到第一个点
			if g.Point[j]==g.connect[i].Star{
				for k:=0;k< len(g.Point);k++{	//找到第二个点
					if g.Point[k]==g.connect[i].Ends{
						//找到点所在的坐标
						s:=j
						e:=k
						//初始化Edge的边表连接信息，此处为无向图
						g.Edge[s][e]=1
						g.Edge[e][s]=1
						break
					}
				}
				break
			}

		}
	}
}

// DeepFirstSearch 深度优先查询算法
func (g *Graph) DeepFirstSearch(edge [][]int, star int)  {
	  pl:= len(g.Point)
	  for i:=star;i<pl;i++{
		  for j:=0;j<pl;j++{

			  if !g.Mark[i] {
				  g.Mark[i]=true
				  fmt.Println(g.Point[i])
			  }
			  if edge[i][j]!=0 && !g.Mark[j]{
				  g.DeepFirstSearch(edge,j)
			  }
		  }
		  return
	  }
}

// BreadthFirstSearch 广度优先查询算法
func (g *Graph) BreadthFirstSearch(start int)  {
		queue:=&Queue{}						//声明队列，此处的bfs 第一个元素不用入队 ，直接采取从第一个元素开始
		for i:=start;i< len(g.Point);{
			if !g.Mark[i]{
				g.Mark[i]=true
				fmt.Println(g.Point[i])
			}

			for j:=0;j< len(g.Point);j++{   //把相关元素依次入队等候遍历
				if g.Edge[i][j]==1 && !g.Mark[j] {
					g.Mark[j]=true			//标记访问
					queue.EnQueue(j)
				}
			}

			if v:=queue.DeQueue();v!=nil{
				fmt.Println(g.Point[v.value])
				i=v.value		//更新 遍历入口，吧当前遍历到的节点 所连接的节点添加到 队列
			}else {
				g.Mark=make([]bool, len(g.Point))
				return			//当没有新的元素入队 读取为nil 则结束遍历
			}
		}
}

func (g *Graph) ShowEdge()  {
	for _,v:=range g.Edge {
		fmt.Println(v)
	}
}



