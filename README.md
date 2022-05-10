# Algorithm

## 图

### 矩阵结构

#### 图结构设计

```go
type Graph struct {
	Point []string //图顶点信息
	Edge  [][]int  //边关系表
	line  []Line   //连接信息
	Mark  []bool   //访问标识
}
```

#### 边信息设计

```go
// Line 连接信息
type Line struct {
	Star string
	End  string
}
```
