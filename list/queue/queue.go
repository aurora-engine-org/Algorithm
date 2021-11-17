package queue

/*
	实现通用队列
*/

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	Head *Node
	End  *Node
}

func (q *Queue) En(v int) {
	n := &Node{Value: v}
	if q.Head == nil {
		q.Head = n
		q.End = n
		return
	}
	q.End.Next = n
	q.End = n
}

func (q *Queue) De() *Node {
	if q.Head == nil {
		q.End = nil
		return nil
	}
	n := q.Head
	q.Head = q.Head.Next
	return n
}
