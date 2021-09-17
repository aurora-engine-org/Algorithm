package tree

import (
	"testing"
)

func TestCreateBinaryTree(t *testing.T) {
	r:= &node{value: 10}
	c1:=&node{value: 5}
	c2:=&node{value: 15}
	c3:=&node{value: 3}
	c4:=&node{value: 7}
	c5:=&node{value: 12}
	c6:=&node{value: 6}
	c7:=&node{value: 0}
	c8:=&node{value: 1}
	BinarySearchTree(r,c1)
	BinarySearchTree(r,c2)
	BinarySearchTree(r,c3)
	BinarySearchTree(r,c4)
	BinarySearchTree(r,c5)
	BinarySearchTree(r,c6)
	BinarySearchTree(r,c7)
	BinarySearchTree(r,c8)
	//BinaryTreeInorderTraversal(r)
/*	h:=MaxDepth(r)

	b:=isAvl(r)
	fmt.Println(b)
	fmt.Println(h)*/
}

func TestAvl(t *testing.T) {
	r:= &avlNode{value: 30}
	c1:=&avlNode{value: 20}
	c2:=&avlNode{value: 50}
	c3:=&avlNode{value: 10}
	c4:=&avlNode{value: 7}
	//c5:=&avlNode{value: 12}
	//c6:=&avlNode{value: 6}
	//c7:=&avlNode{value: 0}
	//c8:=&avlNode{value: 1}
	AvlTree(r,c1)
	AvlTree(r,c2)
	AvlTree(r,c3)
	AvlTree(r,c4)
	if b,h:=isAvl(r);b{
		t.Log("is ",b," ",h)
	}

	if b,h:=isAvl(r);!b{
		t.Log("is ",b," ",h)
	}

}