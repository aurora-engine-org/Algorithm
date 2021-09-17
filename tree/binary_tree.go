package tree

import (
	"fmt"
	"math"
)

// BinarySearchTree 创建排序二叉树
func BinarySearchTree(root *node,child *node) {
	if child.value<root.value{
		if root.left==nil{
			root.left=child
			return
		}
		BinarySearchTree(root.left,child)
		return
	}
	if child.value>root.value {
		if root.right==nil{
			root.right=child
			return
		}
		BinarySearchTree(root.right,child)
		return
	}

	if child.value==root.value{
		return
	}
}

// AvlTree 生产平衡二叉树
func AvlTree(root *avlNode,child *avlNode) *avlNode {
	if child.value<root.value{
		if root.left==nil{
			child.parent=root				//指向父节点
			root.left=child					//添加新左子节点
			return root
		}
		r:=AvlTree(root.left,child)            //添加完成
		//检测树是否平衡
		if b,h:=isAvl(r);!b {
			//如果此节点不平衡
			if h<0 {    //右子树不平衡
			}
		}
		return root
	}

	if child.value>root.value {
		if root.right==nil{
			child.parent=root 				//指向父节点
			root.right=child  				//添加新右子节点
			return root						//返回父节点
		}
		r:=AvlTree(root.right,child)
		//检测树是否平衡
		if b,h:=isAvl(r);!b {
			//如果此节点不平衡
			if h>0{
			}
			if h<0 {    //右子树不平衡
			}
		}
		return root
	}
	if child.value==root.value{
		//节点已存在
		return root
	}
	return nil
}

// 检测是否为avl树，并且返回高度差
func isAvl(root *avlNode) (bool,int) {
	h:=0				//当前节点层级深度
	ln:=0				//左子树深度
	rn:=0				//右子树深度
	h++				    //计算当前一步
	if root.left!=nil{
		ln=AvlDepth(root.left)
	}
	if root.right!=nil{
		rn=AvlDepth(root.right)
	}
	//检查当前节点的平衡性
	v:=math.Abs(float64(ln-rn))
	if v==0 || v==1{
		return true ,int(v)
	}
	return false,ln-rn  //以左子树 减 右子树为标准来区分  大于0 左子树过高，小于0 右子树过高
}


// BinaryTreeInorderTraversal 二叉树中序遍历
func BinaryTreeInorderTraversal(root *node) {
	 if root==nil{
		 return
	 }
	BinaryTreeInorderTraversal(root.left)
	 fmt.Println(root.value)
	BinaryTreeInorderTraversal(root.right)
}

// MaxDepth 计算二叉树深度
func MaxDepth(root *node) int  {
	 h:=0				//当前节点层级深度
	 ln:=0				//左子树深度
	 rn:=0				//右子树深度
	 h++				//计算当前一步
	if root.left!=nil{
	 	ln=MaxDepth(root.left)
	 }
	if root.right!=nil{
		rn=MaxDepth(root.right)
	}
	if ln>rn{ //对比那个子树深度最大就返回那个
		return ln+h
	}else {
		return rn+h
	}
}

// AvlDepth 平衡二叉树计算深度
func AvlDepth(root *avlNode) int  {
	h:=0				//当前节点层级深度
	ln:=0				//左子树深度
	rn:=0				//右子树深度
	h++				//计算当前一步
	if root.left!=nil{
		ln=AvlDepth(root.left)
	}
	if root.right!=nil{
		rn=AvlDepth(root.right)
	}
	if ln>rn{ //对比那个子树深度最大就返回那个
		return ln+h
	}else {
		return rn+h
	}
}
