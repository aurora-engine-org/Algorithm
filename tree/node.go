package tree

type node struct {
	 value int
	 left  *node
	 right *node
}

type avlNode struct {
	value int
	parent 	*avlNode
	brother *avlNode
	left  	*avlNode
	right 	*avlNode
}
